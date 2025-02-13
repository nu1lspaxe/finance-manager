package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"

	"finance_manager/cmd/middleware"
	v1 "finance_manager/cmd/v1"
	"finance_manager/configs"
	"finance_manager/pkg/workers"
)

func main() {
	logger := middleware.Logger()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := InitDB(ctx)
	if err != nil {
		logger.Errorf("Failed to connect to postgres database: %v\n", err)
		return
	}

	router := v1.SetupRouter(pool)

	srv := &http.Server{
		Addr:              ":8989",
		Handler:           router,
		WriteTimeout:      configs.TimeOutSeconds * time.Second,
		ReadHeaderTimeout: configs.TimeOutSeconds * time.Second,
	}

	go func() {
		if err2 := srv.ListenAndServe(); err2 != nil && err2 != http.ErrServerClosed {
			logger.Errorf("Server error: %v\n", err2)
		}
	}()

	requests := make(chan *http.Request, configs.MaxWorkers)

	for range configs.MaxWorkers {
		go workers.Worker(ctx, requests, router)
	}

	<-ctx.Done()
	close(requests)
	shutdown(ctx, srv, pool, logger)
}

func InitDB(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, os.Getenv("PG_URL"))
	if err != nil {
		return nil, err
	}

	if connErr := pool.Ping(ctx); connErr != nil {
		return nil, connErr
	}

	return pool, nil
}

func shutdown(ctx context.Context, srv *http.Server, pool *pgxpool.Pool, logger *logrus.Logger) {
	logger.Info("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(ctx, configs.TimeOutSeconds*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Errorf("Server shutdown error: %v\n", err)
		return
	}

	pool.Close()
	logger.Info("Server shutdown successfully")
}
