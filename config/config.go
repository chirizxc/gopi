package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
)

type TomlConfig struct {
	HTTPServer HTTPServer `toml:"http_server"`
	Database   Database   `toml:"database"`
}

type HTTPServer struct {
	Port string `toml:"PORT"`
	User string `toml:"USER"`
	Pass string `toml:"PASS"`
}

type Database struct {
	Dsn string `toml:"MYSQL_DSN"`
}

func LoadConfig() *TomlConfig {
	path := filepath.Join("config.toml")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", path)
	}

	var cfg TomlConfig

	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
