package db

import (
	"context"
	"database/sql"
	"time"

	"example/backend/config"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	fmt.Printf("Connecting to database with DSN: %s\n", cfg.DB.DSN)
	for {
		db, err := sql.Open("postgres", cfg.DB.DSN)
		if err == nil {
			if err = db.Ping(); err == nil {
				break
			}
		}
		fmt.Printf("Waiting for postgresql...")
		time.Sleep(2 * time.Second)
	}
	db, err := sql.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)

	maxIdleTime := cfg.DB.MaxIdleTime
	if maxIdleTime == "" {
		maxIdleTime = "5m" // Default to 5 minutes if not set
	}
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
