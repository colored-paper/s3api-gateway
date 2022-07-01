package gateway

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/net/http2"
	"net/http"
	"strconv"

	"github.com/coloered-paper/s3api-gateway/s3"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	engine    *echo.Echo
	s3Service s3.Actor
	config    Config
}

func NewRouter(ctx context.Context, s3Service s3.Actor, cfg Config) *Router {
	return &Router{
		engine:    echo.New(),
		s3Service: s3Service,
		config:    cfg,
	}
}

func (r *Router) Run(ctx context.Context) error {
	gatewayPort, err := convertPort(r.config.GatewayPort)
	if err != nil {
		return err
	}

	http2Server := &http2.Server{
		MaxConcurrentStreams: r.config.HTTP2.MaxConcurrentStreams,
		MaxReadFrameSize:     r.config.HTTP2.MaxReadFrameSize,
		IdleTimeout:          r.config.HTTP2.IdleTimeout,
	}

	if middlewareErr := r.Middleware(); middlewareErr != nil {
		return middlewareErr
	}
	if handlerErr := r.Handler(); handlerErr != nil {
		return handlerErr
	}

	return r.engine.StartH2CServer(gatewayPort, http2Server)
}

func (r *Router) Middleware() error {
	r.engine.Use(setRequestID)
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recover())
	if r.config.Metric.Enable {
		metricPort, err := convertPort(r.config.Metric.Port)
		if err != nil {
			return err
		}
		go r.runMetricServer(metricPort)
	}
	return nil
}

func setRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := uuid.New().String()
		c.Request().Header.Set(echo.HeaderXRequestID, requestID)
		c.Response().Header().Set(s3.XAMZRequestID, requestID)
		return next(c)
	}
}

func (r *Router) Handler() error {
	r.engine.GET("/api/v1/health", healthHandler)

	r.engine.DELETE("/:bucket/:key", r.s3Service.AbortMultipartUpload)
	r.engine.PUT("/:bucket", r.s3Service.CreateBucket)
	return nil
}

func healthHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (r *Router) runMetricServer(metricPort string) {
	engine := echo.New()
	engine.HideBanner = true
	engine.HidePort = true

	prom := prometheus.NewPrometheus("echo", nil)
	prom.SetMetricsPath(engine)

	r.engine.Use(prom.HandlerFunc)
	r.engine.Logger.Fatal(engine.Start(metricPort))
}

func convertPort(port int) (string, error) {
	if port < 0 {
		return "", InvalidPortNumber
	}
	return ":" + strconv.Itoa(port), nil
}
