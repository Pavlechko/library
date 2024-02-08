package grpcapp

import (
	"fmt"
	"net"

	grpcservice "github.com/pavlechko/library/data_service/internal/grpc"
	"github.com/pavlechko/library/data_service/internal/utils"
	"google.golang.org/grpc"
)

type App struct {
	port   uint
	server *grpc.Server
}

func NewGRPCServer(port uint, bookService grpcservice.IBook) *App {
	server := grpc.NewServer()

	grpcservice.Register(server, bookService)

	return &App{
		port:   port,
		server: server,
	}
}

func (a *App) ListenAndServe() error {
	utils.InfoLogger.Println("ListenAndServe")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		utils.ErrorLogger.Println("Error: %w", err)
		return fmt.Errorf("Error: %w", err)
	}

	utils.InfoLogger.Println("grpc server is running", lis.Addr().String())

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
