/**
 * @Author: EDZ
 * @Description:
 * @File: routes
 * @Version: 1.0.0
 * @Date: 2022/6/7 10:19
 */

package routes

import (
	"ghub/internal/tasks/sendmsg"
	"github.com/google/wire"
	"github.com/hibiken/asynq"
)

var HandlersSet = wire.NewSet(wire.Struct(new(Handlers), "*"))

type Handlers struct {
	SendmsgHandler *sendmsg.Handler
}

//任务路由
func (l *Handlers) InitHandlers(mux *asynq.ServeMux) *asynq.ServeMux {

	//发送注册短信demo
	mux.HandleFunc(sendmsg.TypeMemberRegisterMsg, l.SendmsgHandler.Handler)
	return mux
}
