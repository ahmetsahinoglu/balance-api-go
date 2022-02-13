package ethereum

import (
	"balance-api-go/internal/model"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type (
	Service struct {
		ethClient EthClient
		logger    *zap.Logger
	}

	EthClient interface {
		GetBalance(ctx *fiber.Ctx, address string) (*big.Int, error)
	}
)

func NewEthService(ethClient EthClient, logger *zap.Logger) *Service {
	return &Service{
		ethClient: ethClient,
		logger:    logger,
	}
}

func (s Service) GetBalance(ctx *fiber.Ctx, request *model.BalanceRequest) (model.BalanceResponseData, error) {
	balances := make(model.BalanceMap)
	var notValidAddress []interface{}

	for _, address := range request.Addresses {
		isValidAddress := common.IsHexAddress(fmt.Sprint(address))
		if !isValidAddress {
			notValidAddress = append(notValidAddress, address)
		} else {
			balance, err := s.ethClient.GetBalance(ctx, fmt.Sprint(address))
			if err != nil {
				return model.BalanceResponseData{}, err
			}
			s.logger.Info(fmt.Sprintf("Balance is %v Wei %f Eth\n", balance, WeiToEther(balance)))
			balances[fmt.Sprint(address)] = WeiToEther(balance)
		}
	}

	balanceResponseData := model.BalanceResponseData{
		Balances:         balances,
		InvalidAddresses: notValidAddress,
	}
	return balanceResponseData, nil
}

func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}
