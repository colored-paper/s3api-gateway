package config

import (
	"time"

	"github.com/coloered-paper/s3api-gateway/pkg/router"
	"github.com/coloered-paper/s3api-gateway/pkg/s3"
)

type Config struct {
	Router router.Config
	S3     s3.Config
}

func NewWithDefaultValues() Config {
	return Config{
		Router: router.Config{
			GatewayPort: 8080,
			HTTP2: router.HTTP2Config{
				MaxConcurrentStreams: 250,
				MaxReadFrameSize:     1048576,
				IdleTimeout:          10 * time.Second,
			},
			Metric: router.MetricConfig{
				Enable: true,
				Port:   8090,
			},
		},
		S3: s3.Config{
			ActorType: "skeleton",
		},
	}
}
