package main

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gopi/internal/config"
	pl "gopi/internal/lib/handlers/prettyloger"
	l "gopi/internal/lib/middleware/logger"
	"gopi/internal/server/handlers/save"
	"gopi/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()
	log := setupPrettySlog()

	db, err := sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		log.Error("Failed to connect to the database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	s := &storage.Storage{Db: db}

	saveHandler := save.New(s)

	r := gin.Default()
	r.Use(l.New(log))
	r.POST("/create", saveHandler)

	go func() {
		if err := r.Run(cfg.HTTPServer.Port); err != nil {
			log.Error("Server startup error", slog.String("error", err.Error()))
		}
	}()

	log.Info("Server started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	log.Info("Stopping server")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("Server stopped")
}

func setupPrettySlog() *slog.Logger {
	opts := pl.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
