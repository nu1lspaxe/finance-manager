package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	v1 "finance_manager/cmd/v1"
	"finance_manager/configs"
	"finance_manager/pkg/workers"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := InitDB(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to PostgreSQL: %v\n", err)
		<-ctx.Done()
		return
	}
	defer pool.Close()

	router := v1.SetupRouter(pool)

	srv := &http.Server{
		Addr:              ":8989",
		Handler:           router,
		WriteTimeout:      configs.TimeOutSeconds * time.Second,
		ReadHeaderTimeout: configs.TimeOutSeconds * time.Second,
	}

	go func() {
		if err2 := srv.ListenAndServe(); err2 != nil && err2 != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Unable to start server: %v\n", err2)
		}
	}()

	requests := make(chan *http.Request, configs.MaxWorkers)
	var wg sync.WaitGroup

	for range configs.MaxWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			workers.Worker(ctx, requests, router)
		}()
	}

	<-ctx.Done()
	close(requests)
	wg.Wait()

	shutdown(ctx, srv)
}

func InitDB(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, os.Getenv("PG_URL"))
	if err != nil {
		return nil, err
	}

	if connErr := pool.Ping(ctx); connErr != nil {
		return nil, connErr
	}

	fmt.Fprintln(os.Stdout, "Successfully connected to PostgreSQL")
	return pool, nil
}

func shutdown(ctx context.Context, srv *http.Server) {
	fmt.Fprint(os.Stdout, "Shutting down gracefully, press Ctrl+C again to force shutdown\n")

	shutdownCtx, cancel := context.WithTimeout(ctx, configs.TimeOutSeconds*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Fprintf(os.Stderr, "Server forced to shutdown: %v\n", err)
	}

	fmt.Fprintln(os.Stdout, "Server exiting")
}
