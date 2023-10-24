package main

import (
	"go-base-web/server"
	"sync"
)

// main starts/runs the server.
func main() {
	wg := &sync.WaitGroup{}
	server.OrchestrateServer(wg, nil)
	wg.Wait()
}
