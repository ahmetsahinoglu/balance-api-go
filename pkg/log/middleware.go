package log

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	RequestIDHeader     = "x-request-id"
	ClientTraceIDHeader = "x-client-trace-id"
)

var loggerKey = "logger"

func Middleware(logger *zap.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var fields []zap.Field
		requestID := c.Get(RequestIDHeader, uuid.New().String())
		clientTraceID := c.Get(ClientTraceIDHeader)

		if requestID != "" {
			fields = append(fields, RequestID(requestID))
		}
		if clientTraceID != "" {
			fields = append(fields, ClientTraceID(clientTraceID))
		}
		c.Locals(loggerKey, logger.With(fields...))

		return c.Next()
	}
}

func FromContext(ctx context.Context, fields ...zap.Field) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		if len(fields) > 0 {
			return l.With(fields...)
		}

		return l
	}

	return zap.New(zap.L().Core())
}

func ClientTraceID(requestID string) zap.Field {
	return zap.String("clientTraceID", requestID)
}

func RequestID(requestID string) zap.Field {
	return zap.String("requestID", requestID)
}
