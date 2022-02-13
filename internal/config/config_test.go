package config_test

import (
	"testing"

	"balance-api-go/internal/config"

	"github.com/stretchr/testify/assert"
)

const configPath = "../../test/testdata"
const fakeConfigPath = "../test/fake"
const configName = "test-config"

func TestGivenTestConfigFileWhenICallNewThenItShouldReturnConfig(t *testing.T) {
	actualConfig, _ := config.New(configPath, configName)

	expectedConfig := config.Config{
		AppName:  "balance-api-go",
		Server:   config.Server{Port: ":3001"},
		LogLevel: "info",
		Clients:  config.Clients{Eth: "ethClientUrl"},
	}

	assert.Equal(t, expectedConfig, actualConfig)
}

func TestGivenNotExistingConfigFileWhenICallNewThenItShouldReturnError(t *testing.T) {
	_, err := config.New(fakeConfigPath, "nothing")

	assert.NotNil(t, err)
}
