package utils

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

// WaitForShutdown waits until the SIGINT signal is received or the context is cancelled.
func WaitForShutdown(ctx context.Context, logger *zap.Logger) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	select {
	case <-c:
		logger.Info("Received SIGINT, no longer waiting for shutdown.")
		return nil
	case <-ctx.Done():
		logger.Info("Context cancelled, no longer waiting for shutdown.")
		return fmt.Errorf("WaitForShutdown: context cancelled")
	}
}
