//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-blog/internal/biz"
	"kratos-blog/internal/conf"
	"kratos-blog/internal/data"
	"kratos-blog/internal/server"
	"kratos-blog/internal/service"
	"kratos-blog/pkg/zap"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

func newLogger(cfg *conf.Bootstrap) log.Logger {
	c := cfg.Log
	logger := log.With(zap.NewLogger(c),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	return logger
}

func newConfig(path string) *conf.Bootstrap {
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	bc := new(conf.Bootstrap)
	if err := c.Scan(bc); err != nil {
		panic(err)
	}

	return bc
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

var providerSet = wire.NewSet(
	newConfig,
	newLogger,
	server.ProviderSet,
	data.ProviderSet,
	biz.ProviderSet,
	service.ProviderSet,
)

// createApp init kratos application.
func createApp(path string) (*kratos.App, func(), error) {
	panic(wire.Build(providerSet, newApp))
}
