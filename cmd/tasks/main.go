// Package tasks
// @author: ekin
// @date: 2022/5/17
// @Description: 任务服务端
package main

import (
	"flag"
	"ghub/internal/tasks"
	"ghub/internal/tasks/middlewares"
	"ghub/internal/tasks/routes"
	"github.com/china-xs/gin-tpl/pkg/redis"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"runtime"
)

// 定义项目配置文件
var configFile = flag.String("f", "../../configs/app.yaml", "set config file which viper will loading.")

func main() {
	flag.Parse()
	worker, fc, err := initApp(*configFile)
	if err != nil {
		panic(err)
	}
	defer fc()

	if err := worker.Run(); err != nil {
		panic(err)
	}

}

func newApp(handlers *routes.Handlers, options *redis.Options, log *zap.Logger) *tasks.Worker {
	worker := tasks.NewAsynqWorker(log)
	mux := asynq.NewServeMux()
	mux.Use(middlewares.LoggingHandler(log))
	handlers.InitHandlers(mux)
	worker.Mux(mux)

	c := asynq.Config{
		Concurrency: runtime.NumCPU(), //并发
		Queues: map[string]int{ //优先级
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	}
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: options.Addr, Password: options.Password, DB: options.DB},
		c,
	)
	worker.Server(server)
	return worker
}
