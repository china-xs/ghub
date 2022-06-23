/**
 * @Author: ekin
 * @Description:任务
 * @File: provider
 * @Version: 1.0.0
 * @Date: 2022/6/6 18:14
 */

package tasks

import (
	"ghub/internal/tasks/routes"
	"ghub/internal/tasks/sendmsg"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	routes.HandlersSet,
	NewTaskClient,
	sendmsg.NewSendMsgHandler, //发送短信demo
)
