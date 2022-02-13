package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EthAddress common.Address

type BalanceMap map[string]*big.Float

type BalanceResponse struct {
	Data    BalanceResponseData `json:"data"`
	Status  int                 `json:"status"`
	Message string              `json:"message"`
}

type BalanceResponseData struct {
	Balances         BalanceMap    `json:"balances"`
	InvalidAddresses []interface{} `json:"invalidAddresses"`
}

type BalanceRequest struct {
	Addresses []interface{} `json:"addresses"`
}
