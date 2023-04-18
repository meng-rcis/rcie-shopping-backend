package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/config"
)

func responseTimeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		status := strconv.Itoa(c.Response().Status)
		method := c.Request().Method
		path := c.Path()

		config.HttpDuration.WithLabelValues(method, path).Observe(duration.Seconds())
		config.HttpRequestsTotal.WithLabelValues(status, method, path).Inc()

		return err
	}
}

func EnableResponseTimeMiddleware(e *echo.Echo) {
	e.Use(responseTimeMiddleware)
}
