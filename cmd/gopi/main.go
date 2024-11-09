package main

import (
	"gopi/internal/config"
	"gopi/internal/server"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	r := server.NewServer()

	if err := r.Run(cfg.HTTPServer.Port); err != nil {
		log.Fatal("Server startup error:", err)
	}
}
