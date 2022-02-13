package ethereum

import (
	"fmt"
	"net/http"

	"balance-api-go/internal/model"
	"balance-api-go/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type (
	EthService interface {
		GetBalance(ctx *fiber.Ctx, request *model.BalanceRequest) (model.BalanceResponseData, error)
	}

	EthHandler struct {
		service EthService
		logger  *zap.Logger
	}
)

func NewEthHandler(service EthService, logger *zap.Logger) *EthHandler {
	return &EthHandler{
		service: service,
		logger:  logger,
	}
}

func (h *EthHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/v1/ethereum")

	api.Post("/", h.GetBalance)
}

func (h *EthHandler) GetBalance(ctx *fiber.Ctx) error {
	request := model.BalanceRequest{}

	if err := ctx.BodyParser(&request); err != nil {
		return errors.BadRequest("").Err(err)
	}

	if err := validateRequest(request); err != nil {
		return errors.BadRequest(err.Error())
	}

	balanceResponseData, err := h.service.GetBalance(ctx, &request)

	balanceResponse := model.BalanceResponse{
		Data:    balanceResponseData,
		Status:  http.StatusOK,
		Message: "OK",
	}

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(balanceResponse)
}

func validateRequest(request model.BalanceRequest) error {
	if len(request.Addresses) == 0 {
		return fmt.Errorf("addresses can not be empty")
	}
	return nil
}
