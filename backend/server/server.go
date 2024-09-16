package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"scrapquiz/config"

	"github.com/labstack/echo/v4"
)

func Run(ctx context.Context, conf *config.Config) {
	e := echo.New()
	InitRoute(e, conf.Server.APIKey)

	address := conf.Server.Address + ":" + conf.Server.Port
	log.Printf("Starting server on %v...\n", address)

	srv := &http.Server{
		Addr:              address,
		Handler:           e,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %v: %v\n", address, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(ctx, time.Duration(conf.Server.GracefulShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
