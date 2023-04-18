package middleware

import (
	"time"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/config"
)

func responseTimeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start)
		config.ResponseTimeHistogram.WithLabelValues(c.Request().Method, c.Path()).Observe(duration.Seconds())
		return err
	}
}

func EnableResponseTimeMiddleware(e *echo.Echo) {
	e.Use(responseTimeMiddleware)
}
