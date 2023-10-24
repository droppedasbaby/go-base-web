package utils

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"time"
)

const (
	ConnReadIdleTimeoutS  = 10 * time.Second
	ConnWriteIdleTimeoutS = 10 * time.Second
)

// RunServer runs the http server used to accept and handle requests.
// The requests are routed to the correct handler.
func RunServer(ctx context.Context, logger *zap.Logger, addr string, r *chi.Mux) {
	wg := sync.WaitGroup{}

	logger.Info("Starting server on:", zap.String("addr", addr))
	server := &http.Server{Addr: addr, Handler: r, ReadTimeout: ConnReadIdleTimeoutS, WriteTimeout: ConnWriteIdleTimeoutS}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Info("Server no longer listening.")
			} else {
				logger.Error("Failed to start server.", zap.Error(err))
			}
		}
	}()

	logger.Info("Server running...")
	<-ctx.Done()
	err := server.Shutdown(ctx)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("Error while shutting down server. %w", zap.Error(err))
	}

	wg.Wait()
	logger.Info("Server shutdown.")
}
