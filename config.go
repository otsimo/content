package content

import "fmt"

const (
	DefaultGrpcPort = 18859
	DefaultHttpPort = 18871
)

type Config struct {
	Debug            bool
	GrpcPort         int
	InsecureGrpcPort int
	HttpPort         int
	TlsCertFile      string
	TlsKeyFile       string
	NoRedis          bool
	GitUrl           string
	GitBranch        string
	GitFolder        string
	RedisAddr        string
	RedisPassword    string
	RedisDB          int64
	PublicDir        string
	Host             string
	Secret           string
	DefaultLanguage  string
}

func (c *Config) GetGrpcPortString() string {
	return fmt.Sprintf(":%d", c.GrpcPort)
}

func (c *Config) GetHttpPortString() string {
	return fmt.Sprintf(":%d", c.HttpPort)
}

func NewConfig() *Config {
	return &Config{GrpcPort: DefaultGrpcPort, HttpPort: DefaultHttpPort, InsecureGrpcPort: 0}
}
