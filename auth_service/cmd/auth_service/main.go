package main

import (
	"context"
	"time"

	"github.com/pavlechko/library/auth_service/internal/grpc"
	"github.com/pavlechko/library/auth_service/internal/routes"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	grpsClient := grpc.NewClient()
	server := routes.NewAPIClient(":8080", grpsClient)
	server.Run(ctx)
}
