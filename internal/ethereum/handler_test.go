package ethereum_test

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"balance-api-go/internal/ethereum"
	"balance-api-go/internal/mocks"
	"balance-api-go/internal/model"
	"balance-api-go/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGivenValidArrayWhenGetBalanceIsCalledThenItShouldReturnProperResponseBalance(t *testing.T) {
	const validEthAddress = "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	expectedBalanceResponseData := model.BalanceResponseData{
		Balances: model.BalanceMap{
			validEthAddress: big.NewFloat(82),
		},
		InvalidAddresses: []interface{}{"1", 22, map[string]string{"name": "ahmet"}},
	}

	request := model.BalanceRequest{
		Addresses: []interface{}{
			validEthAddress,
			"1",
			22,
			map[string]string{"name": "ahmet"},
		},
	}

	expectedBalanceResponse := model.BalanceResponse{
		Data:    expectedBalanceResponseData,
		Status:  http.StatusOK,
		Message: "OK",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockEthService(ctrl)

	service.EXPECT().
		GetBalance(gomock.Any(), gomock.Any()).
		Return(expectedBalanceResponseData, nil)

	app := createTestApp()

	userHandler := ethereum.NewEthHandler(service, zap.NewExample())
	userHandler.RegisterRoutes(app)

	req := NewHTTPRequestWithJSONBody(http.MethodPost, "/v1/ethereum", request)

	resp, err := app.Test(req)

	assert.Nil(t, err)

	defer resp.Body.Close()

	actualResponse := model.BalanceResponse{}
	_ = json.NewDecoder(resp.Body).Decode(&actualResponse)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, expectedBalanceResponse.Data.Balances[validEthAddress].String(), actualResponse.Data.Balances[validEthAddress].String())
	assert.Equal(t, len(expectedBalanceResponse.Data.InvalidAddresses), len(actualResponse.Data.InvalidAddresses))
	assert.Equal(t, expectedBalanceResponse.Status, actualResponse.Status)
	assert.Equal(t, expectedBalanceResponse.Message, actualResponse.Message)
}

func TestGivenInvalidRequestWhenGetBalanceIsCalledThenItShouldReturnBadRequest(t *testing.T) {
	expectedBalanceResponse := model.BalanceResponse{
		Status:  http.StatusBadRequest,
		Message: "addresses can not be empty",
	}

	request := model.BalanceRequest{
		Addresses: []interface{}{},
	}

	app := createTestApp()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := mocks.NewMockEthService(ctrl)

	ethHandler := ethereum.NewEthHandler(service, zap.NewExample())
	ethHandler.RegisterRoutes(app)

	req := NewHTTPRequestWithJSONBody(http.MethodPost, "/v1/ethereum", request)

	resp, err := app.Test(req)

	assert.Nil(t, err)

	defer resp.Body.Close()

	actualResponse := model.BalanceResponse{}
	_ = json.NewDecoder(resp.Body).Decode(&actualResponse)

	assert.Equal(t, expectedBalanceResponse.Status, resp.StatusCode)
	assert.Equal(t, expectedBalanceResponse.Message, actualResponse.Message)
}

func NewHTTPRequestWithJSONBody(method, url string, requestBody interface{}) *http.Request {
	request, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(method, url, bytes.NewReader(request))
	req.Header.Add("Content-type", "application/json")

	return req
}

func createTestApp() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler: errors.Handler(),
	})
}
