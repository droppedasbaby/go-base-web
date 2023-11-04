package main

import (
	"sync"

	"go-base-web/server"
)

// main starts/runs the server.
func main() {
	wg := &sync.WaitGroup{}
	server.OrchestrateServer(wg, nil)
	wg.Wait()
}
