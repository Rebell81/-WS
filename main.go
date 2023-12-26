package main

import (
	"context"
	"cws/config"
	"cws/sweek"
	"fmt"
	"log"
	"os"
)

func main() {
	ctx, cancelFunction := context.WithCancel(context.Background())
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()

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
