package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/otsimo/content"
)

var Version string
var config = content.NewConfig()

func RunAction(c *cli.Context) error {
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
			flgs = append(flgs, cli.IntFlag{Name: v.Name, Value: v.Value, Destination: v.Destination, Usage: v.Usage, EnvVar: env})
		case cli.Int64Flag:
			flgs = append(flgs, cli.Int64Flag{Name: v.Name, Value: v.Value, Destination: v.Destination, Usage: v.Usage, EnvVar: env})
		case cli.Float64Flag:
			flgs = append(flgs, cli.Float64Flag{Name: v.Name, Value: v.Value, Destination: v.Destination, Usage: v.Usage, EnvVar: env})
		case cli.Uint64Flag:
			flgs = append(flgs, cli.Uint64Flag{Name: v.Name, Value: v.Value, Destination: v.Destination, Usage: v.Usage, EnvVar: env})
		case cli.StringFlag:
			flgs = append(flgs, cli.StringFlag{Name: v.Name, Value: v.Value, Destination: v.Destination, Usage: v.Usage, EnvVar: env})
		case cli.BoolFlag:
			flgs = append(flgs, cli.BoolFlag{Name: v.Name, Usage: v.Usage, Destination: v.Destination, EnvVar: env})
		case cli.DurationFlag:
			flgs = append(flgs, cli.DurationFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env, Destination: v.Destination})
		case cli.StringSliceFlag:
			flgs = append(flgs, cli.StringSliceFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env})
		case cli.IntSliceFlag:
			flgs = append(flgs, cli.IntSliceFlag{Name: v.Name, Value: v.Value, Usage: v.Usage, EnvVar: env})
		default:
			fmt.Println("unknown flag %#v", f)
			panic(fmt.Sprintf("%#v", f))
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
		cli.IntFlag{Name: "grpc-port", Value: config.GrpcPort, Destination: &config.GrpcPort, Usage: "server port"},
		cli.IntFlag{Name: "insecure-port", Value: config.InsecureGrpcPort, Destination: &config.InsecureGrpcPort, Usage: "server port"},
		cli.IntFlag{Name: "http-port", Value: config.HttpPort, Destination: &config.HttpPort, Usage: "server port"},
		cli.StringFlag{Name: "git-url", Value: "https://github.com/otsimo/wiki.git", Destination: &config.GitUrl, Usage: "the content wiki git project url"},
		cli.StringFlag{Name: "host", Value: "https://content.otsimo.com", Destination: &config.Host, Usage: "services' host url"},
		cli.StringFlag{Name: "git-path", Value: "/opt/otsimo/project", Destination: &config.GitFolder, Usage: "where to put git project"},
		cli.StringFlag{Name: "git-branch", Value: "master", Destination: &config.GitBranch, Usage: "git branch to clone"},
		cli.StringFlag{Name: "tls-cert-file", Value: "", Destination: &config.TlsCertFile, Usage: "the server's certificate file for TLS connection"},
		cli.StringFlag{Name: "tls-key-file", Value: "", Destination: &config.TlsKeyFile, Usage: "the server's private key file for TLS connection"},
		cli.StringFlag{Name: "redis-addr", Destination: &config.RedisAddr, Value: "localhost:6379", Usage: "redis address"},
		cli.StringFlag{Name: "redis-password", Value: "", Destination: &config.RedisPassword, Usage: "redis password"},
		cli.StringFlag{Name: "webhook-secret", Value: "", Destination: &config.Secret, Usage: "the webhook secret"},
		cli.StringFlag{Name: "default-lang", Value: "tr", Destination: &config.DefaultLanguage, Usage: "the default language"},
		cli.Int64Flag{Name: "redis-db", Value: 0, Destination: &config.RedisDB, Usage: "redis db"},
		cli.BoolFlag{Name: "no-redis", Destination: &config.NoRedis, Usage: "don't use redis"},
		cli.BoolFlag{Name: "debug, d", Destination: &config.Debug, Usage: "enable verbose log"},
		cli.StringFlag{Name: "public-dir", Value: "public", Destination: &config.PublicDir, Usage: "public directory where parsed markdown files will put"},
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
