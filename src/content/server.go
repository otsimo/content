package content

import (
	"net"
	"os"
	"strings"

	"net/http"

	log "github.com/Sirupsen/logrus"
	pb "github.com/otsimo/api/apipb"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type Server struct {
	Config  *Config
	Content *ContentManager
	Redis   *RedisClient

	Secret     string     // Option secret key for authenticating via HMAC
	IgnoreTags bool       // If set to false, also execute command if tag is pushed
	Events     chan Event // Channel of events. Read from this channel to get push events as they happen.
}

func (s *Server) GRPCServer() *grpc.Server {
	var l = &log.Logger{
		Out:       os.Stdout,
		Formatter: &log.TextFormatter{FullTimestamp: true},
		Hooks:     make(log.LevelHooks),
		Level:     log.GetLevel(),
	}
	grpclog.SetLogger(l)

	var opts []grpc.ServerOption
	if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(s.Config.TlsCertFile, s.Config.TlsKeyFile)
		if err != nil {
			log.Fatalf("server.go: Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	contentGrpc := &contentGrpcServer{
		server: s,
	}

	pb.RegisterContentServiceServer(grpcServer, contentGrpc)
	log.Infof("server.go: Binding %s for grpc", s.Config.GetPortString())
	return grpcServer
}

func NewServer(config *Config) *Server {
	server := &Server{
		Config:     config,
		Content:    NewContentManager(config),
		IgnoreTags: true,
		Events:     make(chan Event, 10), // buffered to 10 items
		Secret:     config.Secret,
	}
	return server
}

func (s *Server) Start() {
	err := s.Content.Init()

	if err != nil {
		panic(err)
	}

	if !s.Config.NoRedis {
		s.Redis = NewRedisClient(s.Config, s.Content)
	}

	go s.TrackEvent()
	s.Listen()
}

func (s *Server) TrackEvent() {
	for {
		select {
		case e, ok := <-s.Events:
			log.Debugf("server.go: event: %v", e)
			if !ok {
				return
			}
			if e.Type != "push" {
				continue
			}
			//todo(sercan) check repo
			log.Infof("updating repo by event %+v", e)

			err := s.Content.Update(e.Commit)
			if err != nil {
				log.Errorf("failed to update repository")
				continue
			}
			if !s.Config.NoRedis {
				//todo(sercan) publish to redis
			}
		}
	}
}

func (s *Server) grpcHandlerFunc(rpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			rpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func (s *Server) Listen() {
	//Non-TLS
	if s.Config.TlsCertFile == "" || s.Config.TlsKeyFile == "" {
		// Create the main listener.
		l, err := net.Listen("tcp", s.Config.GetPortString())
		if err != nil {
			log.Fatalf("server.go: failed to listen %v", err)
		}

		// Create a cmux.
		m := cmux.New(l)

		// Match connections in order:
		grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		httpL := m.Match(cmux.HTTP1Fast())

		// Create your protocol servers.
		grpcS := s.GRPCServer()
		echo := s.HttpServer()
		httpS := echo.Server(s.Config.GetPortString())

		go grpcS.Serve(grpcL)
		go httpS.Serve(httpL)

		if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
			panic(err)
		}
	} else {
		//TLS
		gserver := s.GRPCServer()
		echo := s.HttpServer()
		srv := &http.Server{
			Addr:    s.Config.GetPortString(),
			Handler: s.grpcHandlerFunc(gserver, echo),
		}
		if err := srv.ListenAndServeTLS(s.Config.TlsCertFile, s.Config.TlsKeyFile); err != nil {
			panic(err)
		}
		log.Infoln("closing server")
	}
}
