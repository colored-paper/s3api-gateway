package router

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/coloered-paper/s3api-gateway/pkg/s3"
	"github.com/coloered-paper/s3api-gateway/pkg/s3/dummy"
)

var testConfig Config

func TestS3API(t *testing.T) {
	ctx := context.Background()
	testConfig = Config{
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
	}

	go runRouter(ctx)

	// Test APIs
}

func runRouter(ctx context.Context) {
	router := New(ctx, dummy.New(ctx, s3.Config{}), testConfig)
	if err := router.Run(ctx); err != http.ErrServerClosed {
		panic(err)
	}
}
