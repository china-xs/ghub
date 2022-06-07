/**
 * @Author: ekin
 * @Description:日志
 * @File: log
 * @Version: 1.0.0
 * @Date: 2022/6/7 13:33
 */

package middlewares

import (
	"context"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"time"
)

func LoggingHandler(log *zap.Logger) asynq.MiddlewareFunc {
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			start := time.Now()
			log.Info("processing", zap.String("task type", t.Type()))
			err := h.ProcessTask(ctx, t)
			if err != nil {
				return err
			}
			log.Info("processed",
				zap.String("type", t.Type()),
				zap.String("payload", string(t.Payload())),
				zap.Error(err),
				zap.Duration("elapsed", time.Since(start)),
			)
			return nil
		})
	}
}
