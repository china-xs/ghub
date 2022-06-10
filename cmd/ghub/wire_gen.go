// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ghub/internal/cache"
	role2 "ghub/internal/cache/role"
	"ghub/internal/data"
	"ghub/internal/data/account"
	"ghub/internal/data/role"
	"ghub/internal/routes"
	service4 "ghub/internal/service"
	service3 "ghub/internal/service/v1/apidemo"
	service2 "ghub/internal/service/v1/auth"
	"ghub/internal/service/v1/helloword"
	"ghub/internal/tasks"
	"github.com/china-xs/gin-tpl"
	"github.com/china-xs/gin-tpl/pkg/api_sign"
	"github.com/china-xs/gin-tpl/pkg/config"
	"github.com/china-xs/gin-tpl/pkg/db"
	"github.com/china-xs/gin-tpl/pkg/jwt_auth"
	"github.com/china-xs/gin-tpl/pkg/log"
	"github.com/china-xs/gin-tpl/pkg/redis"
	"github.com/google/wire"
)

// Injectors from wire.go:

// cf config path
func initApp(path string) (*gin_tpl.Server, func(), error) {
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
	redisOptions, err := redis.NewOps(viper)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	client, err := redis.New(redisOptions)
	if err != nil {
		cleanup()
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
	transaction := data.NewTx(dbData)
	greeterService := service.NewGreeterService(logger, client, gormDB, transaction)
	repo := account.NewRepo(dbData, logger)
	roleRepo := role.NewRepo(dbData, logger)
	cache := role2.NewCache(logger, client, roleRepo)
	signupService := service2.NewSignupService(logger, transaction, repo, roleRepo, cache)
	jwt_authOptions, err := jwt_auth.NewOps(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	jwtAuth := jwt_auth.NewJwtAuth(jwt_authOptions)
	api_signOptions, err := api_sign.NewOps(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	apidemoService := service3.NewApidemoService(jwt_authOptions, jwtAuth, api_signOptions)
	routesRoutes := routes.Routes{
		HelloSrv:  greeterService,
		V1Signup:  signupService,
		V1Apidemo: apidemoService,
	}
	server := newApp(routesRoutes, logger, viper)
	return server, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var ProviderSet = wire.NewSet(config.ProviderSet, log.ProviderSet, redis.ProviderSet, service4.ProviderSet, routes.ProviderSet, data.ProviderSet, cache.ProviderSet, tasks.ProviderSet, jwt_auth.ProviderSet, api_sign.ProviderSet)
