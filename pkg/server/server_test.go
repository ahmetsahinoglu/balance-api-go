package server_test

import (
	"balance-api-go/internal/config"
	"balance-api-go/pkg/server"
	"fmt"

	"net/http"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGivenServerConfigWhenICallRunThenItShouldRunOnSpecifiedPort(t *testing.T) {
	port, err := freeport.GetFreePort()
	assert.Nil(t, err)
	serverConfig := config.Server{Port: fmt.Sprintf(":%d", port)}

	s := server.New(serverConfig, []server.Handler{}, zap.NewExample())
	go func() {
		_ = s.Run()
	}()

	time.Sleep(50 * time.Millisecond)
	testEndpointURL := fmt.Sprintf("http://localhost%s/health", serverConfig.Port)
	req, err := http.NewRequest(http.MethodGet, testEndpointURL, http.NoBody)
	assert.Nil(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
