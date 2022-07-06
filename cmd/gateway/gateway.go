package gateway

import (
	"context"

	"github.com/coloered-paper/s3api-gateway/config"
	"github.com/coloered-paper/s3api-gateway/pkg/router"
	"github.com/coloered-paper/s3api-gateway/pkg/s3"
)

type Gateway struct {
	router *router.Router
	actor  s3.Actor
}

func New(ctx context.Context, cfg config.Config) (*Gateway, error) {
	actor, err := newActor(ctx, cfg.S3)
	if err != nil {
		return nil, err
	}

	return &Gateway{
		router: router.New(ctx, actor, cfg.Router),
		actor:  actor,
	}, nil
}

func (g *Gateway) Run(ctx context.Context) error {
	return g.router.Run(ctx)
}
