package main

import (
	"context"
	"net/http"

	"github.com/coloered-paper/s3api-gateway/cmd/gateway"
	"github.com/coloered-paper/s3api-gateway/config"
)

func main() {
	ctx := context.Background()
	cfg := config.NewWithDefaultValues()

	s3APIGateway, initErr := gateway.New(ctx, cfg)
	if initErr != nil {
		panic(initErr)
	}

	if runErr := s3APIGateway.Run(ctx); runErr != http.ErrServerClosed {
		panic(runErr)
	}
}
