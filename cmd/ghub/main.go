// Package ghub
// @author: xs
// @date: 2022/5/17
// @Description: ghub
package main

import (
	"flag"
	"fmt"
	"ghub/internal/routes"
	tpl "github.com/china-xs/gin-tpl"
	"github.com/china-xs/gin-tpl/middleware"
	"github.com/china-xs/gin-tpl/middleware/logger"
	"github.com/china-xs/gin-tpl/middleware/validate"
	"github.com/kataras/i18n"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 定义项目配置文件
var configFile = flag.String("f", "../../configs/app.yaml", "set config file which viper will loading.")

// 定义项目i18n 文件配置路径
var i18nFile = flag.String("if", "../../configs/locales", "set i18n config file")

func main() {
	flag.Parse()
	app, fc, err := initApp(*configFile)
	if err != nil {
		panic(err)
	}
	defer fc()
	if err := app.Run(); err != nil {
		panic(err)
	}

}

func newApp(routes routes.Routes, log *zap.Logger, v *viper.Viper) *tpl.Server {
	var ops []tpl.ServerOption
	I18n, err := i18n.New(i18n.Glob(fmt.Sprintf("%s/*/*", *i18nFile)), "en-US", "zh-CN")
	if err != nil {
		panic(err)
	}
	opts, err := tpl.NewSerOpts(v)
	if err != nil {
		panic(err)
	}
	ms := make([]middleware.Middleware, 2)
	ms[0] = validate.Validator2I18n(I18n)
	ms[1] = logger.Logger(log) // 记录系统级别日志 ps 请求出入request|reply 请求耗时
	ops = append(ops, opts, tpl.Middleware(ms...))
	app := tpl.NewServer(ops...)
	routes.InitRoutes(app)
	return app
}
