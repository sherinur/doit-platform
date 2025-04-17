package main

import (
	"api-gateway/config"
	"api-gateway/internal/app"
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app, err := app.New(ctx, cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = app.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
