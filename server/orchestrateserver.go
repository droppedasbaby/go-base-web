package server

import (
	"context"
	"go-base-web/api"
	"go-base-web/utils"
	"go.uber.org/zap"
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	ConnReadIdleTimeoutS  = 10 * time.Second
	ConnWriteIdleTimeoutS = 10 * time.Second
)

func OrchestrateServer(wg *sync.WaitGroup, logger *zap.Logger) {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8000"
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
		err := utils.WaitForShutdown(ctx, logger)
		utils.PanicIfError(err, context.Canceled)
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ManageServerLifecycle(ctx, logger, addr, r)
		utils.PanicIfError(err)
	}()
}
