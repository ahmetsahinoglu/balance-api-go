package server

import (
	"balance-api-go/internal/config"
	"balance-api-go/pkg/errors"
	appLogger "balance-api-go/pkg/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Handler interface {
	RegisterRoutes(app *fiber.App)
}

type Server struct {
	app    *fiber.App
	config config.Server
	logger *zap.Logger
}

func New(serverConfig config.Server, handlers []Handler, logger *zap.Logger) Server {
	app := fiber.New(fiber.Config{
		ErrorHandler: errors.Handler(),
	})
	server := Server{app: app, config: serverConfig, logger: logger}
	server.app.Use(cors.New())
	server.app.Use(appLogger.Middleware(logger))
	server.addRoutes()

	for _, handler := range handlers {
		handler.RegisterRoutes(server.app)
	}

	return server
}

func (s Server) Run() error {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		shutdownSignal := <-shutdownChan
		s.logger.Info("Received interrupt signal", zap.String("shutdownSignal", shutdownSignal.String()))
		if err := s.app.Shutdown(); err != nil {
			s.logger.Info("Failed to shutdown gracefully", zap.Error(err))
		}
		s.logger.Info("application shutdown gracefully")
	}()
	return s.app.Listen(s.config.Port)
}

func (s Server) addRoutes() {
	s.app.Get("/health", healthCheck)
}

func healthCheck(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return nil
}
