package server

import (
	"context"
	"fmt"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	quesvc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/question/v1"
	quizsvc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/quiz/v1"
	ressvc "github.com/sherinur/doit-platform/apis/gen/quiz-service/service/frontend/result/v1"
	"github.com/sherinur/doit-platform/quiz-service/config"
	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/controller/grpc/server/frontend"
)

type API struct {
	server *grpc.Server
	cfg    config.GRPCServer
	addr   string

	ResultUseCase   ResultUseCase
	QuestionUseCase QuestionUseCase
	QuizUseCase     QuizUseCase
}

func New(
	cfg config.Server,
	ResultUseCase ResultUseCase,
	QuizUseCase QuizUseCase,
	QuestionUseCase QuestionUseCase,
) *API {
	return &API{
		cfg:             cfg.GRPCServer,
		addr:            fmt.Sprintf("0.0.0.0:%d", cfg.GRPCServer.Port),
		ResultUseCase:   ResultUseCase,
		QuestionUseCase: QuestionUseCase,
		QuizUseCase:     QuizUseCase,
	}
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		log.Println("gRPC server starting listen", fmt.Sprintf("addr: %s", a.addr))

		if err := a.run(); err != nil {
			errCh <- fmt.Errorf("can't start grpc server: %w", err)

			return
		}
	}()
}

// Stop method gracefully stops grpc API server. Provide context to force stop on timeout.
func (a *API) Stop(ctx context.Context) error {
	if a.server == nil {
		return nil
	}

	stopped := make(chan struct{})
	go func() {
		a.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done(): // Stop immediately if the context is terminated
		a.server.Stop()
	case <-stopped:
	}

	return nil
}

// run starts and runs GRPCServer server.
func (a *API) run() error {
	a.server = grpc.NewServer()

	// Register services
	quesvc.RegisterQuestionServiceServer(a.server, frontend.NewQuestion(a.QuestionUseCase))
	quizsvc.RegisterQuizServiceServer(a.server, frontend.NewQuiz(a.QuizUseCase))
	ressvc.RegisterResultServiceServer(a.server, frontend.NewResult(a.ResultUseCase))

	// Register reflection service
	reflection.Register(a.server)

	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}

	err = a.server.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to serve grpc: %w", err)
	}

	return nil
}
