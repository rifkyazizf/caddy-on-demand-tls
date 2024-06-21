package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"caddy-on-demand-tls/server"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	httpServer := server.RunHttp()

	gracefulShutdown(httpServer, sigChan)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("failed to start http-server: %v", err))
	}
}

func gracefulShutdown(server *http.Server, sigChan chan os.Signal) {
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sigChan
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}
		cancel()
	}()
}
