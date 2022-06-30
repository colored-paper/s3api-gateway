package config

import (
	"github.com/coloered-paper/s3api-gateway/s3"
	"time"

	"github.com/coloered-paper/s3api-gateway/gateway"
)

type Config struct {
	Gateway gateway.Config
	S3      s3.Config
}

func NewWithDefaultValues() Config {
	return Config{
		Gateway: gateway.Config{
			GatewayPort: 8080,
			HTTP2: gateway.HTTP2Config{
				MaxConcurrentStreams: 250,
				MaxReadFrameSize:     1048576,
				IdleTimeout:          10 * time.Second,
			},
			Metric: gateway.MetricConfig{
				Enable: true,
				Port:   8090,
			},
		},
		S3: s3.Config{
			Name: "skeleton",
		},
	}
}
