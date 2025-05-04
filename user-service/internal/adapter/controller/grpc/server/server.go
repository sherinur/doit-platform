package server

import (
	"context"
	"fmt"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sherinur/doit-platform/user-service/config"
)

type API struct {
	server *grpc.Server
	cfg    config.GRPCServer
	addr   string

	UserUseCase UserUsecase
}

func New(
	cfg config.Server,
	UserUseCase UserUsecase,
) *API {
	return &API{
		cfg:         cfg.GRPCServer,
		addr:        fmt.Sprintf("0.0.0.0:%d", cfg.GRPCServer.Port),
		UserUseCase: UserUseCase,
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
	// quesvc.RegisterQuestionServiceServer(a.server, frontend.NewQuestion(a.QuestionUseCase))
	// quizsvc.RegisterQuizServiceServer(a.server, frontend.NewQuiz(a.QuizUseCase))
	// ressvc.RegisterResultServiceServer(a.server, frontend.NewResult(a.ResultUseCase))

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
