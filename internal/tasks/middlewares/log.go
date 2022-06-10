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
	"fmt"
	"github.com/hibiken/asynq"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"time"
)

func LoggingHandler(log *zap.Logger) asynq.MiddlewareFunc {
	var tracer = otel.Tracer("asynq/tasks")
	return func(h asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
			ctx, span := tracer.Start(ctx, fmt.Sprintf("middleware-task-%s", t.Type()))
			defer span.End()

			start := time.Now()
			log.Info("asynq-processing", zap.String("type", t.Type()))

			err := h.ProcessTask(ctx, t)
			end := time.Since(start)

			if err != nil {
				log.Error("asynq-error",
					zap.String("type", t.Type()),
					zap.String("payload", string(t.Payload())),
					zap.Duration("elapsed", end),
					zap.Error(err),
				)
				return err
			}

			log.Info("asynq-success",
				zap.String("type", t.Type()),
				zap.String("payload", string(t.Payload())),
				zap.Duration("elapsed", end),
			)
			return nil
		})
	}
}
