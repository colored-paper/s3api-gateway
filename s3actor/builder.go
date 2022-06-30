package s3actor

import (
	"context"

	"github.com/coloered-paper/s3api-gateway/s3"
	"github.com/coloered-paper/s3api-gateway/s3/dummy"
	"github.com/coloered-paper/s3api-gateway/s3/skeleton"
)

func New(ctx context.Context, cfg s3.Config) (s3.Actor, error) {
	switch cfg.Name {
	case "skeleton":
		return skeleton.New(ctx, cfg)
	case "dummy":
		return dummy.New(ctx, cfg)
	default:
		return nil, s3.InvalidActorName
	}
}
