package grpcapp

import (
	"fmt"
	"net"

	grpcservice "github.com/pavlechko/library/data_service/internal/grpc"
	"google.golang.org/grpc"
)

type App struct {
	port   uint
	server *grpc.Server
}

func NewGRPCServer(port uint) *App {
	grpcServer := grpc.NewServer()

	grpcservice.Register(grpcServer)

	return &App{
		port:   port,
		server: grpcServer,
	}
}

func (a *App) ListenAndServe() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	fmt.Println("grpc server is running", lis.Addr().String())

	if err := a.server.Serve(lis); err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}

func (a *App) Close() error {
	fmt.Println("grpc server is stopping on port: ", a.port)

	a.server.GracefulStop()
	return nil
}
