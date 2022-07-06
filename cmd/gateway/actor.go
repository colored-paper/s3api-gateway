package gateway

import (
	"context"
	"errors"

	"github.com/coloered-paper/s3api-gateway/pkg/s3"
	"github.com/coloered-paper/s3api-gateway/pkg/s3/dummy"
	"github.com/coloered-paper/s3api-gateway/pkg/s3/skeleton"
)

var (
	InvalidActorName = errors.New("invalid actor name")
)

func newActor(ctx context.Context, cfg s3.Config) (s3.Actor, error) {
	switch cfg.ActorType {
	case "skeleton":
		return skeleton.New(ctx, cfg), nil
	case "dummy":
		return dummy.New(ctx, cfg), nil
	default:
		return nil, InvalidActorName
	}
}
