package router

import "time"

type Config struct {
	GatewayPort int
	HTTP2       HTTP2Config
	Metric      MetricConfig
}

type HTTP2Config struct {
	MaxConcurrentStreams uint32
	MaxReadFrameSize     uint32
	IdleTimeout          time.Duration
}

type MetricConfig struct {
	Enable bool
	Port   int
}
