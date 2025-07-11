package main

import (
	"example/backend/api"
	"example/backend/config"
	"example/backend/db"
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	cfg := config.NewConfig()
	if err := cfg.ParseFlags(); err != nil {
		logger.Fatal("failed to parse command-line flags", zap.Error(err))
	}

	database, err := db.Connect(cfg)
	if err != nil {
		logger.Fatal("failed to connect to the database", zap.Error(err))
	}
	defer database.Close()
	logger.Info("database connection established successfully")

	repos := cfg.InitializeRepositories(database)
	handlers := cfg.InitializeHandlers(repos)
	srv := api.NewAPI(cfg, handlers)

	logger.Info("starting server", zap.String("address", cfg.Port))
	if err := srv.Run(); err != nil {
		logger.Fatal("failed to start the server", zap.Error(err))
	}
}
