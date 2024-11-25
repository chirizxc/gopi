package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopi/internal/config"
	del "gopi/internal/server/handlers/delete"
	"gopi/internal/server/handlers/get/gif"
	"gopi/internal/server/handlers/get/gifs"
	"gopi/internal/server/handlers/save"
	auth "gopi/internal/server/middleware/auth"
	"gopi/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	time.Sleep(10 * time.Second) // Ожидание 10 секунд

	gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()

	db, err := storage.NewDB(cfg.Database.Dsn)
	if err != nil {
		fmt.Println("Failed to initialize database:", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	s := &storage.Storage{Db: db}
	authMiddleware := auth.New(cfg.HTTPServer.User, cfg.HTTPServer.Pass)

	router := gin.New()

	router.POST("/save", save.New(s))
	router.DELETE("/delete/:id", authMiddleware, del.New(s))
	router.GET("/gif/:id", gif.New(s))
	router.GET("/gifs", gifs.New(s))

	// Запуск сервера
	go func() {
		if err := router.Run(fmt.Sprintf(":%s", cfg.HTTPServer.Port)); err != nil {
			fmt.Println("Server startup error:", err.Error())
		}
	}()

	// Ожидание сигналов для завершения работы
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Завершение работы сервера
	fmt.Println("Server stopped")
}
