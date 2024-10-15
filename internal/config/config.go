package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)

type TomlConfig struct {
	Path       PathConfig `toml:"path"`
	HTTPServer HTTPServer `toml:"http_server"`
}

type PathConfig struct {
	Storage string `toml:"storage"`
}

type HTTPServer struct {
	Address     string        `toml:"address"`
	Timeout     time.Duration `toml:"timeout"`
	IdleTimeout time.Duration `toml:"idle_timeout"`
	User        string        `toml:"user"`
	Pass        string        `toml:"pass"`
}

func LoadConfig() *TomlConfig {
	configPath := "config.toml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg TomlConfig

	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
