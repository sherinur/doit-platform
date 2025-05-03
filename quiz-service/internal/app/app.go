package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sherinur/doit-platform/quiz-service/config"
	grpcserver "github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server"
	httpservice "github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/http/service"
	mongorepo "github.com/sherinur/doit-platform/quiz-service/internal/adapters/repo/mongo"
	"github.com/sherinur/doit-platform/quiz-service/internal/usecase"
	mongocon "github.com/sherinur/doit-platform/quiz-service/pkg/mongo"
)

const serviceName = "quiz-service"

type App struct {
	httpServer *httpservice.API
	grpcServer *grpcserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v server", serviceName))

	log.Println("connecting to mongo", "database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	// Repository
	resultRepo := mongorepo.NewResultRepository(mongoDB.Conn)
	quizRepo := mongorepo.NewQuizRepository(mongoDB.Conn)
	questionRepo := mongorepo.NewQuestionRepository(mongoDB.Conn)

	// UseCase
	ResultUseCase := usecase.NewResultUsecase(resultRepo, quizRepo, questionRepo)
	QuizUseCase := usecase.NewQuizUsecase(quizRepo, questionRepo)
	QuestionUseCase := usecase.NewQuestionUsecase(quizRepo, questionRepo)

	// http server
	httpServer := httpservice.New(cfg.Server, ResultUseCase, QuizUseCase, QuestionUseCase)
	grpcServer := grpcserver.New(cfg.Server, ResultUseCase, QuizUseCase, QuestionUseCase)

	app := &App{
		httpServer: httpServer,
		grpcServer: grpcServer,
	}

	return app, nil
}

func (a *App) Close() {
	err := a.grpcServer.Stop(context.Background())
	if err != nil {
		log.Println("failed to shutdown server", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	a.grpcServer.Run(errCh)

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
