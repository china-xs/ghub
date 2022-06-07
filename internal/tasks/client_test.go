/**
 * @Author: ekin
 * @Description:异步任务测试
 * @File: client_test.go
 * @Version: 1.0.0
 * @Date: 2022/6/7 12:27
 */

package tasks

import (
	"ghub/internal/tasks/sendmsg"
	"github.com/china-xs/gin-tpl/pkg/config"
	"github.com/china-xs/gin-tpl/pkg/redis"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//任务投递
func TestNewTaskClient(t *testing.T) {
	viper, err := config.New("../../configs/app.yaml")
	assert.Nil(t, err)

	options, err := redis.NewOps(viper)
	assert.Nil(t, err)

	client := NewTaskClient(options)

	//立即执行
	task, _ := sendmsg.NewMemberRegisterSendMsgTask("test")
	info, err := client.Enqueue(task)
	t.Log(info)
	assert.Nil(t, err)

	//延迟任务
	task1, _ := sendmsg.NewMemberRegisterSendMsgTask("test1")
	info1, err := client.Enqueue(task1, asynq.ProcessIn(2*time.Minute))
	t.Log(info1)
	assert.Nil(t, err)

	//紧急
	task2, _ := sendmsg.NewMemberRegisterSendMsgTask("test2")
	info2, err := client.Enqueue(task2, asynq.Queue("critical"))
	t.Log(info2)
	assert.Nil(t, err)
}
