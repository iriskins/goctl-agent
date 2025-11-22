package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/HaizorWill/goctl-agent/core"
)

func main() {
	app := core.NewApplication()
	signal.Notify(
		app.ShouldClose,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	go func() {
		sig := <-app.ShouldClose
		fmt.Printf("\nReceived a signal %v, stopping...\n", sig)
		time.Sleep(2 * time.Second)
		app.Stop()
	}()
	app.Run()
}
