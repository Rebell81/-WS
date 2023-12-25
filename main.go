package main

import (
	"context"
	"cws/config"
	"cws/sweek"
	"log"
	"os"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	cfg, err := config.ReadConfig(ctx)
	if err != nil {
		log.Fatalf("err during reading config enviroment: %s", err)
	}

	code, err := sweek.Process(ctx, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	os.Exit(code)
}
