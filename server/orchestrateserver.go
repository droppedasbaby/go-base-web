package server

import (
	"context"
	"os"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"

	"go-base-web/api"
	"go-base-web/utils"
)

const (
	ConnReadIdleTimeoutS  = 10 * time.Second
	ConnWriteIdleTimeoutS = 10 * time.Second
)

func OrchestrateServer(wg *sync.WaitGroup, logger *zap.Logger) {
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8000"
	} else {
		addr = ":" + addr
	}

	var err error
	if logger == nil {
		logger, err = zap.NewProduction(zap.WithCaller(false))
		utils.PanicIfError(err)
		defer utils.PanicIfError(logger.Sync(), syscall.ENOTTY)
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := api.GetRouter(logger)

	wg.Add(1)
	go func() {
		defer wg.Done()
		shErr := utils.WaitForShutdown(ctx, logger)
		utils.PanicIfError(shErr, context.Canceled)
		cancel()
	}()

	wg.Add(1)
	defer wg.Done()
	ManageServerLifecycle(ctx, logger, addr, r)
}
