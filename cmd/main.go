package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gopi/internal/config"
	pl "gopi/internal/lib/handlers/prettyloger"
	del "gopi/internal/server/handlers/delete"
	"gopi/internal/server/handlers/get/gif"
	"gopi/internal/server/handlers/get/gifs"
	"gopi/internal/server/handlers/save"
	auth "gopi/internal/server/middleware/auth"
	l "gopi/internal/server/middleware/logger"
	"gopi/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Waiting for dependencies to initialize...")
	time.Sleep(10 * time.Second) // Ожидание 10 секунд
	fmt.Println("Starting the application...")

	gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()
	fmt.Println("Imported config package successfully.") // Сообщение о успешном импорте

	log := setupPrettySlog()
	fmt.Println("Imported prettyloger package successfully.") // Сообщение о успешном импорте

	db, err := storage.NewDB(cfg.Database.Dsn)
	if err != nil {
		log.Error("Failed to initialize database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()
	fmt.Println("Imported storage package and initialized DB successfully.") // Сообщение о успешном импорте

	s := &storage.Storage{Db: db}
	authMiddleware := auth.New(cfg.HTTPServer.User, cfg.HTTPServer.Pass)
	fmt.Println("Imported auth middleware package successfully.") // Сообщение о успешном импорте

	router := gin.New()
	router.Use(l.New(log))
	fmt.Println("Imported logger middleware package successfully.") // Сообщение о успешном импорте

	router.POST("/save", save.New(s))
	router.DELETE("/delete/:id", authMiddleware, del.New(s))
	router.GET("/gif/:id", gif.New(s))
	router.GET("/gifs", gifs.New(s))
	fmt.Println("Imported server handler packages (save, delete, gif, gifs) successfully.") // Сообщение о успешном импорте

	go func() {
		if err := router.Run(fmt.Sprintf(":%s", cfg.HTTPServer.Port)); err != nil {
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
