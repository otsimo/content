package content

import "fmt"

const (
	DefaultPort = 18859
)

type Config struct {
	Debug         bool
	Port          int
	TlsCertFile   string
	TlsKeyFile    string
	NoRedis       bool
	GitUrl        string
	GitFolder     string
	RedisAddr     string
	RedisPassword string
	RedisDB       int64
	PublicDir     string
	Host          string
}

func (c *Config) GetPortString() string {
	return fmt.Sprintf(":%d", c.Port)
}

func NewConfig() *Config {
	return &Config{Port: DefaultPort}
}
