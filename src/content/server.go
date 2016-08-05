package content

import (
	"net"
	"os"

	log "github.com/Sirupsen/logrus"
	pb "github.com/otsimo/otsimopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"github.com/labstack/echo/engine/standard"
)

type Server struct {
	Config     *Config
	Content    *ContentManager
	Redis      *RedisClient

	Secret     string     // Option secret key for authenticating via HMAC
	IgnoreTags bool       // If set to false, also execute command if tag is pushed
	Events     chan Event // Channel of events. Read from this channel to get push events as they happen.
}

func init() {
	var l = &log.Logger{
		Out:       os.Stdout,
		Formatter: &log.TextFormatter{FullTimestamp: true},
		Hooks:     make(log.LevelHooks),
		Level:     log.GetLevel(),
	}
	grpclog.SetLogger(l)
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

func (s *Server) Start() error {
	err := s.Content.Init()

	if err != nil {
		panic(err)
	}

	if !s.Config.NoRedis {
		s.Redis = NewRedisClient(s.Config, s.Content)
	}

	go s.TrackEvent()
	return s.listenGRPC()
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

func (s *Server) listenGRPC() error {
	grpcPort := s.Config.GetGrpcPortString()
	//Listen
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Errorf("server.go: failed to listen %v for grpc", err)
		return err
	}
	var opts []grpc.ServerOption
	if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
		creds, err := credentials.NewServerTLSFromFile(s.Config.TlsCertFile, s.Config.TlsKeyFile)
		if err != nil {
			log.Fatalf("server.go: Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	//register services
	contentGrpc := &contentGrpcServer{
		server: s,
	}
	pb.RegisterContentServiceServer(grpcServer, contentGrpc)

	log.Infof("server.go: Binding %s for grpc", grpcPort)
	//Serve
	return grpcServer.Serve(lis)
}

func (s *Server) listenHTTP() {
	e := s.HttpServer()
	if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
		e.Run(standard.WithTLS(s.Config.GetHttpPortString(), s.Config.TlsCertFile, s.Config.TlsKeyFile))
	} else {
		e.Run(standard.New(s.Config.GetHttpPortString()))
	}
}