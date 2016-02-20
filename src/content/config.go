package content

import "fmt"

const (
	DefaultGrpcPort = 18859
	DefaultHttpPort = 18860
)

type Config struct {
	Debug         bool
	GrpcPort      int
	TlsCertFile   string
	TlsKeyFile    string
	NoRedis       bool
	GitUrl        string
	GitFolder     string
	RedisAddr     string
	RedisPassword string
	RedisDB       int64
}

func (c *Config) GetGrpcPortString() string {
	return fmt.Sprintf(":%d", c.GrpcPort)
}

func NewConfig() *Config {
	return &Config{GrpcPort: DefaultGrpcPort}
}
