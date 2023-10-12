package main

import (
	"github.com/cuida-me/mvp-backend/internal/infrastructure/server"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config := server.NewConfig()

	api := server.NewApi(config)

	err := api.Bootstrap()
	if err != nil {
		return err
	}

	reflection.Register(api.Server)

	lis, err := net.Listen(config.Network, config.Port)
	if err != nil {
		return err
	}

	err = api.Server.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
