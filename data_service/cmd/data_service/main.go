package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pavlechko/library/data_service/internal/app"
	"github.com/pavlechko/library/data_service/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)

	application := app.NewGRPCServer(cfg.GRPC.Port)
	go application.Server.ListenAndServe()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	application.Server.Close()

	fmt.Println("Application was stopped. Signal: ", sign)
}
