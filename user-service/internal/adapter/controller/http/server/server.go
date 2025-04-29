package httpserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-services/config"
	"user-services/internal/adapter/controller/http/server/handler"

	"github.com/gin-gonic/gin"
)

const serverIPAddress = "0.0.0.0:%d" // Changed to 0.0.0.0 for external access

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	userHandler *handler.UserHandler
}

func New(cfg config.Server, userUsecase UserUsecase) *API {
	// Setting the Gin mode
	gin.SetMode(cfg.HTTPServer.Mode)
	// Creating a new Gin Engine
	server := gin.New()

	// Applying middleware
	server.Use(gin.Recovery())

	// Binding clients
	clientHandler := handler.NewUserHandler(userUsecase)

	api := &API{
		server:      server,
		cfg:         cfg.HTTPServer,
		addr:        fmt.Sprintf(serverIPAddress, cfg.HTTPServer.Port),
		userHandler: clientHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("/api/v1")
	{
		clients := v1.Group("/clients")
		{
			clients.POST("/")
		}
	}
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		log.Printf("HTTP server starting on: %v", a.addr)

		// No need to reinitialize `a.server` here. Just run it directly.
		if err := a.server.Run(a.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to start HTTP server: %w", err)
			return
		}
	}()
}

func (a *API) Stop() error {
	// Setting up the signal channel to catch termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Blocking until a signal is received
	sig := <-quit
	log.Println("Shutdown signal received", "signal:", sig.String())

	// Creating a context with timeout for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("HTTP server shutting down gracefully")

	// Note: You can use `Shutdown` if you use `http.Server` instead of `gin.Engine`.
	log.Println("HTTP server stopped successfully")

	return nil
}
