/**
 * @Author: ekin
 * @Description:异步任务worker
 * @File: asynqTask
 * @Version: 1.0.0
 * @Date: 2022/6/6 17:24
 */

package tasks

import (
	"errors"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"runtime"
)

/**
监听异步/延迟任务
*/

var defaultConfig asynq.Config = asynq.Config{
	Concurrency: runtime.NumCPU(), //并发
	Queues: map[string]int{ //优先级
		"critical": 6,
		"default":  3,
		"low":      1,
	},
}

type (
	Worker struct {
		log     *zap.Logger
		configs asynq.Config
		server  *asynq.Server
		mux     *asynq.ServeMux
	}
)

func NewAsynqWorker(log *zap.Logger,
) *Worker {
	return &Worker{
		log:     log,
		configs: defaultConfig,
	}
}

//handlers
func (l *Worker) Mux(mux *asynq.ServeMux) *Worker {
	l.mux = mux
	return l
}

//server
func (l *Worker) Server(s *asynq.Server) *Worker {
	l.server = s
	return l
}

//server
func (l *Worker) Run() error {
	if l.mux == nil || l.server == nil {
		return errors.New("handlers or server is nil")
	}
	return l.server.Run(l.mux)
}
