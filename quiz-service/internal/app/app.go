package app

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"github.com/sherinur/doit-platform/quiz-service/config"
	grpcserver "github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server"
	httpservice "github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/http/service"
	mongorepo "github.com/sherinur/doit-platform/quiz-service/internal/adapters/repo/mongo"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/repo/redis"
	"github.com/sherinur/doit-platform/quiz-service/internal/usecase"
	mongocon "github.com/sherinur/doit-platform/quiz-service/pkg/mongo"
	rediscon "github.com/sherinur/doit-platform/quiz-service/pkg/redis"
)

const serviceName = "quiz-service"

type App struct {
	log *zap.Logger

	httpServer *httpservice.API
	grpcServer *grpcserver.API

	telemetry *Telemetry
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("starting %v server", serviceName))
	fmt.Println(fmt.Sprintf("starting %v server", serviceName))

	logger.Info("connecting to mongo database", zap.String("db_name", cfg.Mongo.Database))
	fmt.Println("connecting to mongo database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	logger.Info("connecting to Redis", zap.String("redis_uri", cfg.Redis.URI))
	fmt.Println("connecting to Redis", cfg.Redis.URI)
	red := rediscon.NewRedis(cfg.Redis)
	err = red.PingRedis(ctx)
	if err != nil {
		logger.Error("failed to connect to redis", zap.Error(err))
		return nil, fmt.Errorf("redis: %w", err)
	}

	// Repository
	resultRepo := mongorepo.NewResultRepository(mongoDB.Conn)
	quizRepo := mongorepo.NewQuizRepository(mongoDB.Conn)
	questionRepo := mongorepo.NewQuestionRepository(mongoDB.Conn)

	// Redis
	resultRedis := redis.NewResultRedis(red.Conn)
	quizRedis := redis.NewQuizRedis(red.Conn)
	questionRedis := redis.NewQuestionRedis(red.Conn)

	// UseCase
	ResultUseCase := usecase.NewResultUsecase(resultRepo, quizRepo, questionRepo, resultRedis)
	QuizUseCase := usecase.NewQuizUsecase(quizRepo, questionRepo, quizRedis)
	QuestionUseCase := usecase.NewQuestionUsecase(quizRepo, questionRepo, questionRedis)

	// Servers
	httpServer := httpservice.New(cfg.Server, ResultUseCase, QuizUseCase, QuestionUseCase)
	grpcServer := grpcserver.New(cfg.Server, logger, ResultUseCase, QuizUseCase, QuestionUseCase)

	// Telemetry
	//telemetry, err := InitTelemetry(ctx, cfg.Telemetry, logger)
	//if err != nil {
	//	return nil, err
	//}

	app := &App{
		log:        logger,
		httpServer: httpServer,
		grpcServer: grpcServer,
		//telemetry:  telemetry,
	}

	return app, nil
}

func (a *App) Close() {
	err := a.grpcServer.Stop(context.Background())
	if err != nil {
		a.log.Error("failed to shutdown server", zap.Error(err))
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	a.grpcServer.Run(errCh)

	a.log.Info(fmt.Sprintf("server %v started", serviceName))

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		a.log.Info(fmt.Sprintf("received signal: %v. Running graceful shutdown...", s))

		a.Close()
		a.log.Info("graceful shutdown completed!")
	}

	return nil
}
