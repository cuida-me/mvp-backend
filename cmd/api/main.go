package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cuida-me/mvp-backend/internal/infrastructure/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	config := server.NewConfig()

	api := server.NewApi(config)

	err := api.Bootstrap()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         api.Cfg.Port,
		WriteTimeout: api.Cfg.WriteTimeout,
		ReadTimeout:  api.Cfg.ReadTimeout,
		IdleTimeout:  api.Cfg.IdleTimeout,
		Handler:      api.Router,
	}

	fmt.Println("Starting server at port ", api.Cfg.Port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
