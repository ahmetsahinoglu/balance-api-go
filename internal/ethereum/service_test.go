package ethereum_test

import (
	"math/big"
	"testing"

	"balance-api-go/internal/ethereum"
	"balance-api-go/internal/mocks"
	"balance-api-go/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGivenValidRequestWhenGetBalanceIsCalledThenItShouldReturnProperResponse(t *testing.T) {
	validEthAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	expectedClientResponse := big.NewInt(82000000000000000)
	expectedBalanceResponseData := model.BalanceResponseData{
		Balances: model.BalanceMap{
			"0x323b5d4c32345ced77393b3530b1eed0f346429d": ethereum.WeiToEther(expectedClientResponse),
		},
		InvalidAddresses: []interface{}{"1", 22, map[string]string{"name": "ahmet"}},
	}

	request := model.BalanceRequest{
		Addresses: []interface{}{
			"0x323b5d4c32345ced77393b3530b1eed0f346429d",
			"1",
			22,
			map[string]string{"name": "ahmet"},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ethClient := mocks.NewMockEthClient(ctrl)

	ctx := &fiber.Ctx{}

	ethClient.EXPECT().
		GetBalance(ctx, gomock.Any()).
		Return(expectedClientResponse, nil)

	service := ethereum.NewEthService(ethClient, zap.NewExample())

	actualBalanceResponseData, err := service.GetBalance(ctx, &request)

	assert.Nil(t, err)
	assert.Equal(t, expectedBalanceResponseData.Balances[validEthAddress], actualBalanceResponseData.Balances[validEthAddress])
	assert.Equal(t, len(expectedBalanceResponseData.InvalidAddresses), len(actualBalanceResponseData.InvalidAddresses))
}
