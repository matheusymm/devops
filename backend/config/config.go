package config

import (
	"database/sql"
	"flag"
	"os"

	"example/backend/api/handlers"
	"example/backend/db/repositories"
)

type Config struct {
	Port string
	DB   struct {
		DSN		  string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.Port, "port", os.Getenv("PORT"), "API server port")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	// flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	// flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	// flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	return nil
}

func (c *Config) InitializeHandlers(r *repositories.Repositories) *handlers.Handlers {
	return handlers.NewHandlers(r.UserRepository, r.MoodRepository)
}

func (c *Config) InitializeRepositories(db *sql.DB) *repositories.Repositories {
	return repositories.NewRepositories(db)
}