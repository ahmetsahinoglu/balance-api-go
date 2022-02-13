package client

import (
	"context"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type EthClient struct {
	URL    string
	Logger *zap.Logger
	Client http.Client
}

func NewEthClient(url string, logger *zap.Logger) *EthClient {
	return &EthClient{
		URL:    url,
		Logger: logger,
		Client: http.Client{
			Timeout:   5 * time.Second,
			Transport: http.DefaultTransport,
		},
	}
}

func (e *EthClient) GetBalance(ctx *fiber.Ctx, address string) (*big.Int, error) {
	client, _ := ethclient.Dial(e.URL)

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
