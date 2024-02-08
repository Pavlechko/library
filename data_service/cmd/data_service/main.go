package main

import (
	"fmt"

	"github.com/pavlechko/library/data_service/internal/app"
	"github.com/pavlechko/library/data_service/internal/config"
	"github.com/pavlechko/library/data_service/internal/utils"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println(err)
	}
	utils.InfoLogger.Println(cfg)
	fmt.Println(cfg)

	application := app.NewGRPCServer(cfg.GRPC.Port)
	utils.InfoLogger.Println("application: ", application)

	application.Server.ListenAndServe()

	defer application.Server.Close()
}
