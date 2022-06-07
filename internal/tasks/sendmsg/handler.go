/**
 * @Author: ekin
 * @Description:发送短信demo
 * @File: taskshandler
 * @Version: 1.0.0
 * @Date: 2022/6/6 18:02
 */

package sendmsg

import (
	"context"
	"encoding/json"
	"ghub/internal/data/account"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Handler struct {
	log         *zap.Logger
	accountRepo *account.Repo
}

func NewSendMsgHandler(log *zap.Logger,
	accountRepo *account.Repo,
) *Handler {
	return &Handler{
		log:         log,
		accountRepo: accountRepo,
	}
}

func (l *Handler) Handler(ctx context.Context, t *asynq.Task) error {
	l.log.Info("coming SendMsgHandler")

	var p MemberRegisterMsgPayLoad
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Errorf("SendMsgHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	//业务操作
	//l.accountRepo.Save

	return nil
}
