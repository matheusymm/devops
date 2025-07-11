package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"example/backend/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	const maxRetries = 10
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", cfg.DB.DSN)
		if err != nil {
			return nil, fmt.Errorf("failed to open database connection: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = db.PingContext(ctx)
		if err == nil {
			fmt.Println("Successfully connected to the database.")
			break
		}

		db.Close()
		fmt.Printf("Could not connect to database: %v. Retrying in 10 seconds...\n", err)
		time.Sleep(10 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d retries: %w", maxRetries, err)
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	maxIdleTime := cfg.DB.MaxIdleTime
	if maxIdleTime == "" {
		maxIdleTime = "5m"
	}
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("invalid value for DB_MAX_IDLE_TIME: %w", err)
	}
	db.SetConnMaxIdleTime(duration)

	return db, nil
}
