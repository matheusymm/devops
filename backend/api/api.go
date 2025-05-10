package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"example/backend/api/handlers"
	"example/backend/api/middleware"
	"example/backend/api/routes"
	"example/backend/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
)

type API struct {
	Router *chi.Mux
	Config *config.Config
	Logger *zap.SugaredLogger
	Wg     *sync.WaitGroup
}

func NewAPI(cfg *config.Config, h *handlers.Handlers) *API {
	router := chi.NewRouter()

	router.Use(
		middleware.RequestLogger,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")}, // Frontend origin
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value for preflight request cache
		}),
	)
	routes.SetupRoutes(router, h)

	return &API{
		Router: router,
		Config: cfg,
	}
}

func (a *API) Run() error {
	srv := &http.Server{
		Addr:         "0.0.0.0:" + a.Config.Port,
		Handler:      a.Router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sign := <-quit

		log.Print("Caught signal", "signal", sign.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		log.Print("Completing background tasks", "addr", srv.Addr)

		a.Wg.Wait()
		shutdownError <- nil
	}()

	log.Print("Starting server on ", srv.Addr)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	log.Print("Stopped server", "addr", srv.Addr)

	return nil
}
