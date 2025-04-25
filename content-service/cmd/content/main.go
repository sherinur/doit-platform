package main

import (
	"context"
	"fmt"
	"os"

	"content-service/config"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)

	cfg, err := config.New()
	if err != nil {
		fmt.Println("Config Error:", err)
		os.Exit(1)
	}

	fmt.Println("--------CONFIG:--------")
	fmt.Println("CONN_STR: ", cfg.Postgres.ConnStr)
	fmt.Println(cfg.Server.GRPCServer.Port)

	// app, err := app.New(ctx, cfg)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// err = app.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}
