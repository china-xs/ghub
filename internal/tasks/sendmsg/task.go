/**
 * @Author: ekin
 * @Description:发送短信异步任务demo
 * @File: sendMsg
 * @Version: 1.0.0
 * @Date: 2022/6/6 17:38
 */

package sendmsg

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

const (
	TypeMemberRegisterMsg = "member:register:send:code"
)

// 任务数据载体
type MemberRegisterMsgPayLoad struct {
	Code string
}

//返回新任务
func NewMemberRegisterSendMsgTask(code string) (*asynq.Task, error) {
	payload, err := json.Marshal(MemberRegisterMsgPayLoad{Code: code})
	if err != nil {
		return nil, errors.New("发送任务失败")
	}
	return asynq.NewTask(TypeMemberRegisterMsg, payload), nil
}
