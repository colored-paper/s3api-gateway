package main

import (
	"context"
	"net/http"

	"github.com/coloered-paper/s3api-gateway/config"
	"github.com/coloered-paper/s3api-gateway/gateway"
	"github.com/coloered-paper/s3api-gateway/s3actor"
)

func main() {
	ctx := context.Background()
	cfg := config.NewWithDefaultValues()

	s3Actor, actorErr := s3actor.New(ctx, cfg.S3)
	if actorErr != nil {
		panic(actorErr)
	}

	router := gateway.NewRouter(ctx, s3Actor, cfg.Gateway)
	if err := router.Run(ctx); err != http.ErrServerClosed {
		panic(err)
	}
}
