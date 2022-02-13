package main

import (
	"balance-api-go/internal/client"
	"balance-api-go/internal/config"
	"balance-api-go/internal/ethereum"
	appLogger "balance-api-go/pkg/log"
	"balance-api-go/pkg/server"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	appEnv := "local"
	conf, err := config.New(".config", appEnv)
	if err != nil {
		return err
	}
	conf.Print()

	logger := appLogger.New(conf.LogLevel)

	defer func() {
		logErr := logger.Sync()
		if logErr != nil {
			fmt.Println(logErr)
		}
	}()

	ethClient := client.NewEthClient(conf.Clients.Eth, logger)
	ethService := ethereum.NewEthService(ethClient, logger)
	ethHandler := ethereum.NewEthHandler(ethService, logger)

	handlers := []server.Handler{
		ethHandler,
	}

	s := server.New(conf.Server, handlers, logger)

	return s.Run()
}
