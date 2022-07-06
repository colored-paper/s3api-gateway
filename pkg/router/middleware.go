package router

import (
	"github.com/coloered-paper/s3api-gateway/pkg/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (r *Router) Middleware() {
	r.engine.Use(setRequestID)
	r.engine.Use(middleware.Logger())
	r.engine.Use(middleware.Recover())
	if r.config.Metric.Enable {
		go r.runMetricServer(convertPort(r.config.Metric.Port))
	}
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

func setRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := uuid.New().String()
		c.Request().Header.Set(echo.HeaderXRequestID, requestID)
		c.Response().Header().Set(s3.XAMZRequestID, requestID)
		return next(c)
	}
}
