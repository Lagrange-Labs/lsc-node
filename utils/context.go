package utils

import (
	"context"
	"time"
)

const (
	// DefaultContextTimeout is the default timeout for context.
	DefaultContextTimeout = 10 * time.Second
)

// GetContext returns a context with a default timeout.
func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), DefaultContextTimeout)
	return ctx
}

// GetContextWithTimeout returns a context with a timeout.
func GetContextWithTimeout(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	return ctx
}

// GetContextWithCancel returns a context with a timeout and a cancel function.
func GetContextWithCancel() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultContextTimeout)
}
