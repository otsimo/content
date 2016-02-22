package content

import (
	"net"
	"os"
	"strings"

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
		Config:  config,
		Content: NewContentManager(config),
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
	s.Listen()
}

func (s *Server) Listen() {
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
}
