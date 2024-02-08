package app

import (
	grpcapp "github.com/pavlechko/library/data_service/internal/app/grpc"
	"github.com/pavlechko/library/data_service/internal/services/book"
	"github.com/pavlechko/library/data_service/internal/storage/postgres"
	"github.com/pavlechko/library/data_service/internal/utils"
)

type App struct {
	Server *grpcapp.App
}

func NewGRPCServer(port uint) *App {
	utils.InfoLogger.Println("NewGRPCServer is started")
	storage, err := postgres.ConnectDB()
	utils.InfoLogger.Println("NewGRPCServer after ConnectDB")
	if err != nil {
		utils.ErrorLogger.Println("Some Error: ", err)
	}
	bookService := book.NewBookService(storage)
	utils.InfoLogger.Println(bookService)

	grpsApp := grpcapp.NewGRPCServer(port, bookService)
	utils.InfoLogger.Println(grpsApp)

	return &App{
		Server: grpsApp,
	}
}
