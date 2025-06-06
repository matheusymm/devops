package main

import (
	"example/backend/api"
	"example/backend/config"
	"example/backend/db"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()

	err := cfg.ParseFlags()
	if err != nil {
		fmt.Println("Failed to parse command-line flags", zap.Error(err))
	}

	db, err := db.Connect(cfg)
	if err != nil {
		fmt.Println("Failed to connect to the database", zap.Error(err))
		panic(err)
	}
	defer db.Close()

	hr := cfg.InitializeHandlers(cfg.InitializeRepositories(db))
	srv := api.NewAPI(cfg, hr)

	err = srv.Run()
	if err != nil {
		fmt.Println("Failed to start the server", zap.Error(err))
	}
}
