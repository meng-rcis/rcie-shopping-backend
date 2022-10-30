package context

import (
	"context"
	"time"
)

func WithTimeout(second time.Duration) (context.Context, context.CancelFunc) {
	duration := second * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)

	return ctx, cancel
}
