package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gopi/internal/config"
	"gopi/internal/server/handlers/save"
	"gopi/internal/storage"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	db, err := sql.Open("mysql", cfg.Database.Dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Initialize the storage layer with the DB connection
	s := &storage.Storage{Db: db}

	// Initialize the save handler
	saveHandler := save.New(s)

	// Create a new Gin router
	r := gin.Default()

	// Register the route with the handler
	r.POST("/create", saveHandler)

	// Run the Gin server
	if err := r.Run(cfg.HTTPServer.Port); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
