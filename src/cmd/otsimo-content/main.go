package main

import (
	"content"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var Version string
var config *content.Config = content.NewConfig()

func RunAction(c *cli.Context) {
	config.Debug = c.Bool("debug")
	config.NoRedis = c.Bool("no-redis")
	config.GitUrl = c.String("git-url")
	config.GitFolder = c.String("git-path")
	config.PublicDir = c.String("public-dir")
	config.Host = c.String("host")

	config.RedisAddr = c.String("redis-addr")
	config.RedisDB = int64(c.Int("redis-db"))
	config.RedisPassword = c.String("redis-password")

	config.Port = c.Int("port")

	config.TlsCertFile = c.String("tls-cert-file")
	config.TlsKeyFile = c.String("tls-key-file")

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	server := content.NewServer(config)
	server.Start()
}

func withEnvs(prefix string, flags []cli.Flag) []cli.Flag {
	var flgs []cli.Flag
	for _, f := range flags {
		env := ""
		spr := strings.Split(f.GetName(), ",")
		env = prefix + "_" + strings.ToUpper(strings.Replace(spr[0], "-", "_", -1))
		switch v := f.(type) {
		case cli.IntFlag:
			flgs = append(flgs, cli.IntFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env})
		case cli.StringFlag:
			flgs = append(flgs, cli.StringFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env})
		case cli.BoolFlag:
			flgs = append(flgs, cli.BoolFlag{Name: v.Name, Usage: v.Usage, EnvVar: env})
		default:
			fmt.Println("unknown")
		}
	}
	return flgs
}

func main() {
	app := cli.NewApp()
	app.Name = "otsimo-content"
	app.Version = Version
	app.Usage = "Otsimo Content Service"
	app.Author = "Sercan Degirmenci <sercan@otsimo.com>"
	var flags []cli.Flag

	flags = []cli.Flag{
		cli.IntFlag{Name: "port", Value: content.DefaultPort, Usage: "server port"},
		cli.StringFlag{Name: "git-url", Value: "https://github.com/otsimo/wiki.git", Usage: "the content wiki git project url"},
		cli.StringFlag{Name: "host", Value: "https://content.otsimo.com", Usage: "services' host url"},
		cli.StringFlag{Name: "git-path", Value: "/opt/otsimo/project", Usage: "where to put git project"},
		cli.StringFlag{Name: "tls-cert-file", Value: "", Usage: "the server's certificate file for TLS connection"},
		cli.StringFlag{Name: "tls-key-file", Value: "", Usage: "the server's private key file for TLS connection"},
		cli.StringFlag{Name: "redis-addr", Value: "localhost:6379", Usage: "redis address"},
		cli.StringFlag{Name: "redis-password", Value: "", Usage: "redis password"},
		cli.IntFlag{Name: "redis-db", Value: 0, Usage: "redis db"},
		cli.BoolFlag{Name: "no-redis", Usage: "don't use redis"},
		cli.BoolFlag{Name: "debug, d", Usage: "enable verbose log"},
		cli.StringFlag{Name: "public-dir", Value: "public", Usage: "public directory where parsed markdown files will put"},
	}
	flags = withEnvs("OTSIMO_CONTENT", flags)

	app.Flags = flags
	app.Action = RunAction

	log.Infoln("running", app.Name, "version:", app.Version)

	app.Run(os.Args)
}

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
