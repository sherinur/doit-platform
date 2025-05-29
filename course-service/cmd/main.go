package main

import (
	"context"
	"fmt"
	"os"

	"course-service/internal/app"

	"google.golang.org/api/config/v1"
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
