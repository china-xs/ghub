// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ghub/internal/cache"
	"ghub/internal/data"
	"ghub/internal/data/account"
	"ghub/internal/service"
	"ghub/internal/tasks"
	"ghub/internal/tasks/routes"
	"ghub/internal/tasks/sendmsg"
	"github.com/china-xs/gin-tpl/pkg/config"
	"github.com/china-xs/gin-tpl/pkg/db"
	"github.com/china-xs/gin-tpl/pkg/log"
	"github.com/china-xs/gin-tpl/pkg/redis"
	"github.com/google/wire"
)

// Injectors from wire.go:

// cf config path
func initApp(path string) (*tasks.Worker, func(), error) {
	viper, err := config.New(path)
	if err != nil {
		return nil, nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup, err := log.New(options)
	if err != nil {
		return nil, nil, err
	}
	dbOptions, err := db.New(viper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	gormDB, cleanup2, err := db.NewDb(dbOptions, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dbData, err := data.StepUp(gormDB, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repo := account.NewRepo(dbData, logger)
	handler := sendmsg.NewSendMsgHandler(logger, repo)
	handlers := &routes.Handlers{
		SendmsgHandler: handler,
	}
	redisOptions, err := redis.NewOps(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	worker := newApp(handlers, redisOptions, logger)
	return worker, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var ProviderSet = wire.NewSet(config.ProviderSet, log.ProviderSet, redis.ProviderSet, service.ProviderSet, data.ProviderSet, cache.ProviderSet, tasks.ProviderSet)
