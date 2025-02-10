package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"

	v1 "finance_manager/cmd/v1"
	"finance_manager/configs"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conn, err := setupDBConnection(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to PostgreSQL: %v\n", err)
		<-ctx.Done()
		return
	}
	defer conn.Close(ctx)

	router := v1.SetupRouter(conn)

	srv := &http.Server{
		Addr:              ":8989",
		Handler:           router,
		WriteTimeout:      configs.TimeOutSeconds * time.Second,
		ReadHeaderTimeout: configs.TimeOutSeconds * time.Second,
	}

	go func() {
		if err2 := srv.ListenAndServe(); err2 != nil && err2 != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err2)
		}
	}()

	<-ctx.Done()

	shutdown(ctx, srv)
}

func setupDBConnection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("PG_URL"))
	if err != nil {
		return nil, err
	}

	if connErr := conn.Ping(ctx); connErr != nil {
		return nil, connErr
	}

	fmt.Fprintln(os.Stdout, "Successfully connected to PostgreSQL")
	return conn, nil
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
