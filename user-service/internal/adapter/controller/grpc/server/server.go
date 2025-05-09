package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sherinur/doit-platform/user-service/config"
	"github.com/sherinur/doit-platform/user-service/internal/adapter/controller/grpc/server/frontend"

	svc "github.com/sherinur/doit-platform/apis/gen/user-service/service/frontend/user/v1"
)

type API struct {
	server *grpc.Server
	cfg    config.GRPCServer
	addr   string
	jwt    config.Jwt

	UserUseCase UserUsecase
}

func New(
	cfg config.Server,
	UserUseCase UserUsecase,
	jwt config.Jwt,
) *API {
	return &API{
		cfg:         cfg.GRPCServer,
		addr:        fmt.Sprintf("0.0.0.0:%d", cfg.GRPCServer.Port),
		jwt:         jwt,
		UserUseCase: UserUseCase,
	}
}

func (a *API) Run(ctx context.Context, errCh chan<- error) {
	go func() {
		log.Println("gRPC server starting listen", fmt.Sprintf("addr: %s", a.addr))

		if err := a.run(ctx); err != nil {
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
func (a *API) run(ctx context.Context) error {
	a.server = grpc.NewServer(a.setOptions(ctx, a.jwt.JwtRefreshSecret)...)

	// Register services
	svc.RegisterUserServiceServer(a.server, frontend.NewUser(a.UserUseCase))

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
