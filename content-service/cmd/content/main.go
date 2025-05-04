package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sherinur/doit-platform/content-service/config"
	"github.com/sherinur/doit-platform/content-service/internal/app"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		fmt.Println("Config Error:", err)
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
