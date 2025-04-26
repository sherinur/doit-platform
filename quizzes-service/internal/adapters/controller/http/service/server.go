package service

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

	"github.com/gin-gonic/gin"
	"quizzes-service/config"
	"quizzes-service/internal/adapters/controller/http/service/handler"
)

const serverIPAddress = "0.0.0.0:%s" // Changed to 0.0.0.0 for external access

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	QuizHandler *handler.QuizHandler
}

func New(cfg config.Server, quizUseCase QuizUseCase) *API {
	// Setting the Gin mode
	gin.SetMode(cfg.HttpServer.Mode)
	// Creating a new Gin Engine
	server := gin.New()

	// Applying middleware
	server.Use(gin.Recovery())

	// Binding clients
	quizHandler := handler.NewQuizHandler(quizUseCase)

	api := &API{
		server:      server,
		cfg:         cfg.HttpServer,
		addr:        fmt.Sprintf(serverIPAddress, cfg.HttpServer.Port),
		QuizHandler: quizHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("/api/v1")
	{
		tv := v1.Group("/quizzes")
		{
			tv.POST("/", a.QuizHandler.CreateQuiz)
			tv.GET("/:id", a.QuizHandler.GetQuizById)
			tv.GET("/", a.QuizHandler.GetQuizAll)
			tv.PUT("/:id", a.QuizHandler.UpdateQuiz)
			tv.DELETE("/:id", a.QuizHandler.DeleteQuiz)
		}
	}
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		log.Printf("HTTP server starting on: %v", a.addr)

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
