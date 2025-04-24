package server

import (
	"inventory_service/config"

	"google.golang.org/grpc"
)

const serverIPAddress = "0.0.0.0:8081"

type API struct {
	server *grpc.Server
	cfg    *config.GRPCServer
	addr   string

	productUsecase ProductUsecase
}

func New(
	cfg config.GRPCServer,
	productUsecase ProductUsecase,
) *API {
	return &API{
		cfg:  &cfg,
		addr: serverIPAddress,
	}
}

func (api *API) Run() error {
	return nil
}

func (api *API) Stop() {}
