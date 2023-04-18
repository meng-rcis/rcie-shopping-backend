package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/config"
)

func prometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		status := strconv.Itoa(c.Response().Status)
		method := c.Request().Method
		path := c.Path()
		duration := time.Since(start)

		config.HttpRequestsTotal.WithLabelValues(status, method, path).Inc()
		config.HttpDuration.WithLabelValues(method, path).Observe(duration.Seconds())

		return err
	}
}

func EnableResponseTimeMiddleware(e *echo.Echo) {
	e.Use(prometheusMiddleware)
}
