package router

import (
	"context"
	"golang.org/x/net/http2"
	"strconv"

	"github.com/coloered-paper/s3api-gateway/pkg/s3"
	"github.com/labstack/echo/v4"
)

type Router struct {
	engine    *echo.Echo
	server    *http2.Server
	s3Service s3.Actor
	config    Config
}

func New(ctx context.Context, s3Service s3.Actor, cfg Config) *Router {
	return &Router{
		engine: echo.New(),
		server: &http2.Server{
			MaxConcurrentStreams: cfg.HTTP2.MaxConcurrentStreams,
			MaxReadFrameSize:     cfg.HTTP2.MaxReadFrameSize,
			IdleTimeout:          cfg.HTTP2.IdleTimeout,
		},
		s3Service: s3Service,
		config:    cfg,
	}
}

func (r *Router) Run(ctx context.Context) error {
	r.Middleware()
	r.Handler()
	return r.engine.StartH2CServer(convertPort(r.config.GatewayPort), r.server)
}

func convertPort(port int) string {
	return ":" + strconv.Itoa(port)
}
