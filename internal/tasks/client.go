/**
 * @Author: ekin
 * @Description: 任务客户端
 * @File: client
 * @Version: 1.0.0
 * @Date: 2022/6/7 11:58
 */

package tasks

import (
	"github.com/china-xs/gin-tpl/pkg/redis"
	"github.com/hibiken/asynq"
)

func NewTaskClient(options *redis.Options) *asynq.Client {
	conf := asynq.RedisClientOpt{Addr: options.Addr, Password: options.Password, DB: options.DB}
	return asynq.NewClient(conf)
}
