package main

import (
	"context"
	"go-base-web/api"
	"go-base-web/utils"
	"go.uber.org/zap"
	"os"
	"sync"
	"syscall"
)

func main() {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8000"
	}

	logger, err := zap.NewProduction(zap.WithCaller(false))
	utils.PanicIfError(err)
	defer utils.PanicIfError(logger.Sync(), syscall.ENOTTY)

	wg := &sync.WaitGroup{}
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
		utils.RunServer(ctx, logger, addr, r)
		utils.PanicIfError(err)
	}()

	wg.Wait()
}
