package content

import (
	"net"
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/labstack/echo/engine/standard"
	"github.com/otsimo/health"
	pb "github.com/otsimo/otsimopb"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"fmt"
)

type Server struct {
	Config     *Config
	Content    *ContentManager
	Redis      *RedisClient
	Secret     string     // Option secret key for authenticating via HMAC
	IgnoreTags bool       // If set to false, also execute command if tag is pushed
	Events     chan Event // Channel of events. Read from this channel to get push events as they happen.
	checks     []health.Checkable
	errChan    chan error
}

func NewServer(config *Config) *Server {
	server := &Server{
		Config:     config,
		Content:    NewContentManager(config),
		IgnoreTags: true,
		Events:     make(chan Event, 10), // buffered to 10 items
		Secret:     config.Secret,
		checks:     []health.Checkable{},
		errChan:    make(chan error, 1),
	}
	return server
}

func (s *Server) run(run func() error) {
	e := run()
	select {
	case s.errChan <- e:
	default:
	}
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
	go s.run(s.listenHTTP)
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
func (s *Server) Healthy() error {
	return nil
}

func (s *Server) listenGRPC() error {
	h := health.New(s)
	var svcs []*grpc.Server

	m := http.NewServeMux()
	m.Handle("/metrics", prometheus.Handler())
	m.Handle("/health", h)
	contentGrpc := &contentGrpcServer{server: s}

	serveGrpc := func(port int, opt ...grpc.ServerOption) error {
		grpcServer := grpc.NewServer(opt...)
		svcs = append(svcs, grpcServer)

		//register services
		pb.RegisterContentServiceServer(grpcServer, contentGrpc)
		grpc_prometheus.Register(grpcServer)
		reflection.Register(grpcServer)
		grpc_health_v1.RegisterHealthServer(grpcServer, h)
		log.Infof("binding :%d for grpc", port)
		//Listen
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Errorf("server.go: failed to listen %v for grpc", err)
			return err
		}
		err = grpcServer.Serve(lis)
		if err != nil {
			log.WithError(err).Errorf("grpc server Serve failed")
		}
		return err
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	}

	go s.run(func() error {
		log.Infof("binding :8080 for health check")
		return http.ListenAndServe(":8080", m)
	})
	//Serve
	if s.Config.InsecureGrpcPort != 0 {
		go s.run(func() error {
			return serveGrpc(s.Config.InsecureGrpcPort, opts...)
		})
	}
	go s.run(func() error {
		o2 := opts
		if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
			creds, err := GrpcCredentials(s.Config.TlsCertFile, s.Config.TlsKeyFile)
			if err != nil {
				log.WithError(err).Errorf("failed to generate credentials")
				return err
			}
			o2 = append(o2, grpc.Creds(creds))
		}
		return serveGrpc(s.Config.GrpcPort, o2...)
	})
	e := <-s.errChan
	for _, g := range svcs {
		g.GracefulStop()
	}
	return e
}

func (s *Server) listenHTTP() error {
	e := s.HttpServer()
	log.Infof("server.go: Binding %s for http", s.Config.GetHttpPortString())
	if s.Config.TlsCertFile != "" && s.Config.TlsKeyFile != "" {
		return e.Run(standard.WithTLS(s.Config.GetHttpPortString(), s.Config.TlsCertFile, s.Config.TlsKeyFile))
	} else {
		return e.Run(standard.New(s.Config.GetHttpPortString()))
	}
}
