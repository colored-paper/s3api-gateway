package gateway

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/coloered-paper/s3api-gateway/config"
	"github.com/coloered-paper/s3api-gateway/s3"
	"github.com/coloered-paper/s3api-gateway/s3actor"
)

var testConfig config.Config

func TestS3API(t *testing.T) {
	ctx := context.Background()
	testConfig = config.Config{
		Gateway: Config{
			GatewayPort: 8080,
			HTTP2: HTTP2Config{
				MaxConcurrentStreams: 250,
				MaxReadFrameSize:     1048576,
				IdleTimeout:          10 * time.Second,
			},
			Metric: MetricConfig{
				Enable: true,
				Port:   8090,
			},
		},
		S3: s3.Config{
			Name: "dummy",
		},
	}

	go runGateway(ctx)

	// Test APIs
}

func runGateway(ctx context.Context) {
	s3Actor, actorErr := s3actor.New(ctx, testConfig.S3)
	if actorErr != nil {
		panic(actorErr)
	}

	router := NewRouter(ctx, s3Actor, testConfig.Gateway)
	if err := router.Run(ctx); err != http.ErrServerClosed {
		panic(err)
	}
}
