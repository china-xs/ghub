//go:build wireinject
// +build wireinject

// Package ghub
// @author: xs
// @date: 2022/5/17
// @Description: ghub
package main

import (
	"ghub/internal/cache"
	"ghub/internal/data"
	"ghub/internal/routes"
	"ghub/internal/service"
	"ghub/internal/tasks"
	tpl "github.com/china-xs/gin-tpl"
	"github.com/china-xs/gin-tpl/pkg/api_sign"
	"github.com/china-xs/gin-tpl/pkg/config"
	"github.com/china-xs/gin-tpl/pkg/jwt_auth"
	"github.com/china-xs/gin-tpl/pkg/log"
	"github.com/china-xs/gin-tpl/pkg/redis"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	config.ProviderSet, //读取yaml 配置
	log.ProviderSet,    // 系统日志
	redis.ProviderSet,
	//db.ProviderSet,		// db 依赖会导致无法识别
	service.ProviderSet, // 控制器
	routes.ProviderSet,  // 路由注册
	data.ProviderSet,
	cache.ProviderSet,    // 缓存
	tasks.ProviderSet,    //任务
	jwt_auth.ProviderSet, //jwt
	api_sign.ProviderSet, //api_sign
)

// cf config path
func initApp(path string) (*tpl.Server, func(), error) {
	panic(wire.Build(
		ProviderSet,
		newApp))
}
