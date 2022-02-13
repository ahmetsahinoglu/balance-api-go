package errors

import (
	"balance-api-go/pkg/log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Handler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if err != nil {
			l := log.FromContext(c.Context())

			res := buildErrorResponse(err)
			l.Error("encountered an error", zap.Error(getError(err, res)), zap.Int("status code", res.StatusCode()))

			if err = c.Status(res.StatusCode()).JSON(res); err != nil {
				l.Error("failed writing error response", zap.Error(err))
				return err
			}
		}
		return nil
	}
}
