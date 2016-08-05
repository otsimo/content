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

func RunAction(c *cli.Context) error {
	config.Debug = c.Bool("debug")
	config.NoRedis = c.Bool("no-redis")
	config.GitUrl = c.String("git-url")
	config.GitFolder = c.String("git-path")
	config.PublicDir = c.String("public-dir")
	config.Host = c.String("host")
	config.Secret = c.String("webhook-secret")
	config.RedisAddr = c.String("redis-addr")
	config.RedisDB = int64(c.Int("redis-db"))
	config.RedisPassword = c.String("redis-password")
	config.HttpPort = c.Int("http-port")
	config.GrpcPort = c.Int("grpc-port")
	config.TlsCertFile = c.String("tls-cert-file")
	config.TlsKeyFile = c.String("tls-key-file")
	config.DefaultLanguage = c.String("default-lang")

	if config.Debug {
		log.SetLevel(log.DebugLevel)
	}

	server := content.NewServer(config)
	return server.Start()
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
		cli.IntFlag{Name: "grpc-port", Value: config.GrpcPort, Usage: "server port"},
		cli.IntFlag{Name: "http-port", Value: config.HttpPort, Usage: "server port"},
		cli.StringFlag{Name: "git-url", Value: "https://github.com/otsimo/wiki.git", Usage: "the content wiki git project url"},
		cli.StringFlag{Name: "host", Value: "https://content.otsimo.com", Usage: "services' host url"},
		cli.StringFlag{Name: "git-path", Value: "/opt/otsimo/project", Usage: "where to put git project"},
		cli.StringFlag{Name: "tls-cert-file", Value: "", Usage: "the server's certificate file for TLS connection"},
		cli.StringFlag{Name: "tls-key-file", Value: "", Usage: "the server's private key file for TLS connection"},
		cli.StringFlag{Name: "redis-addr", Value: "localhost:6379", Usage: "redis address"},
		cli.StringFlag{Name: "redis-password", Value: "", Usage: "redis password"},
		cli.StringFlag{Name: "webhook-secret", Value: "", Usage: "the webhook secret"},
		cli.StringFlag{Name: "default-lang", Value: "tr", Usage: "the default language"},

		cli.IntFlag{Name: "redis-db", Value: 0, Usage: "redis db"},
		cli.BoolFlag{Name: "no-redis", Usage: "don't use redis"},
		cli.BoolFlag{Name: "debug, d", Usage: "enable verbose log"},
		cli.StringFlag{Name: "public-dir", Value: "public", Usage: "public directory where parsed markdown files will put"},
	}
	flags = withEnvs("OTSIMO_CONTENT", flags)

	app.Flags = flags
	app.Action = RunAction

	log.Infoln("running", app.Name, "version:", app.Version)

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
