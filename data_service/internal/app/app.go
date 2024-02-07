package app

import (
	grpcapp "github.com/pavlechko/library/data_service/internal/app/grpc"
)

type App struct {
	Server *grpcapp.App
}

func NewGRPCServer(port uint) *App {
	grpsApp := grpcapp.NewGRPCServer(port)

	return &App{
		Server: grpsApp,
	}
}
