// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"fxkt.tech/bj21/internal/biz"
	"fxkt.tech/bj21/internal/conf"
	"fxkt.tech/bj21/internal/data"
	"fxkt.tech/bj21/internal/server"
	"fxkt.tech/bj21/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	blackJackRepo := data.Newbj21Repo(dataData, logger)
	blackJackUsecase := biz.NewBlackJackUsecase(blackJackRepo, logger)
	blackJackService := service.NewBlackJackService(blackJackUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, blackJackService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
