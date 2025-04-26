package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"quizzes-service/config"
	httpservice "quizzes-service/internal/adapters/controller/http/service"
	mongorepo "quizzes-service/internal/adapters/repo/mongo"
	"quizzes-service/internal/usecase"
	mongocon "quizzes-service/pkg/mongo"
)

const serviceName = "quizzes-service"

type App struct {
	httpServer *httpservice.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v server", serviceName))

	log.Println("connecting to mongo", "database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	// Repository
	quizRepo := mongorepo.NewQuizRepository(mongoDB.Conn)

	// UseCase
	QuizUseCase := usecase.NewQuizUsecase(quizRepo)

	// http server
	httpServer := httpservice.New(cfg.Server, QuizUseCase)

	app := &App{
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Close() {
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("failed to shutdown server", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	a.httpServer.Run(errCh)

	log.Println(fmt.Sprintf("server %v started", serviceName))

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Println(fmt.Sprintf("received signal: %v. Running graceful shutdown...", s))

		a.Close()
		log.Println("graceful shutdown completed!")
	}

	return nil
}
