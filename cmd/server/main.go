package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"fxkt.tech/bj21/internal/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "bj21"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// flagcmd is the command flag.
	flagcmd string

	id, _ = os.Hostname()
)

func init() {
	rand.Seed(time.Now().Unix())
	json.MarshalOptions.UseProtoNames = true
	json.MarshalOptions.EmitUnpopulated = false
	json.UnmarshalOptions.DiscardUnknown = false
	flag.StringVar(&flagcmd, "cmd", "serve", "command. eg: -cmd serve")
	flag.StringVar(&flagconf, "conf", "conf/config.yaml", "config path, eg: -conf conf/config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(gs),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.Timestamp("2006-01-02T15:04:05 "),
		// "caller", log.DefaultCaller,
		// "service.id", id,
		// "service.name", Name,
		// "service.version", Version,
	)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	switch flagcmd {
	case "serve":
		app, cleanup, err := initApp(bc.Server, bc.Data, logger)
		if err != nil {
			panic(err)
		}
		defer cleanup()

		// start and wait for stop signal
		if err := app.Run(); err != nil {
			panic(err)
		}
	}
}
