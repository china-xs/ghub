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
	WorkerOption func(*Worker)
	Worker       struct {
		log      *zap.Logger
		configs  asynq.Config
		server   *asynq.Server
		mux      *asynq.ServeMux
		redisOpt *asynq.RedisClientOpt
	}
)

func NewAsynqWorker(opts ...WorkerOption) (*Worker, error) {
	worker := &Worker{
		configs: defaultConfig,
	}
	for _, o := range opts {
		o(worker)
	}

	if worker.redisOpt == nil || worker.log == nil {
		return nil, errors.New("redisOpt and logger needed")
	}

	server := asynq.NewServer(
		worker.redisOpt,
		worker.configs,
	)
	worker.server = server
	return worker, nil
}

//server
func (l *Worker) Run() error {
	if l.mux == nil || l.server == nil {
		return errors.New("handlers or server is nil")
	}
	return l.server.Run(l.mux)
}

//setting logger
func WithLogger(l *zap.Logger) WorkerOption {
	return func(o *Worker) {
		o.log = l
	}
}

//worker setting
func WithConfig(c asynq.Config) WorkerOption {
	return func(o *Worker) {
		o.configs = c
	}
}

//worker setting
func WithHandlers(m *asynq.ServeMux) WorkerOption {
	return func(o *Worker) {
		o.mux = m
	}
}

//redis opt
func WithRedisOpt(r *asynq.RedisClientOpt) WorkerOption {
	return func(o *Worker) {
		o.redisOpt = r
	}
}
