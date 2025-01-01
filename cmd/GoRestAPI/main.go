package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Akashkarmokar/GoRestAPI/internal/config"
	"github.com/Akashkarmokar/GoRestAPI/internal/http/handlers/student"
	middleware "github.com/Akashkarmokar/GoRestAPI/internal/http/middlewars"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.Handle("/api/students", middleware.CORS(http.HandlerFunc(student.New())))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))

	exitDone := make(chan os.Signal, 1)
	signal.Notify(exitDone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-exitDone
	slog.Info("Shutting down ther server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Faild to shutdown the server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
