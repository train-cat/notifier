package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/train-cat/notifier/helper"
	"github.com/train-cat/notifier/route"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	http.HandleFunc("/health_check", route.HealthCheck)
	http.HandleFunc("/issues", helper.MiddlewareSecurity(helper.MiddlewareNeedPost(route.Issue)))
	http.Handle("/", http.NotFoundHandler())

	h := http.Server{
		Addr: fmt.Sprintf(":%d", viper.GetInt("http.port")),
	}

	go func() {
		log.Infof("HTTP Server listening on %s", h.Addr)

		if err := h.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
			os.Exit(helper.ExitCodeErrorListenServer)
		}
	}()

	<-stop

	log.Info("Graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.Shutdown(ctx)

	if err != nil {
		log.Fatal(err.Error())
		os.Exit(helper.ExitCodeErrorStopServer)
	}
}
