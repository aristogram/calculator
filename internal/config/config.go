package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env  string     `yaml:"env"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file does not exists: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Failed to read configs: " + path + ";" + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches configs path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty.
func fetchConfigPath() string {
	var path string

	// --configs="path/to/config_file.yaml"
	flag.StringVar(&path, "config_path", "", "path to configs file")
	flag.Parse()

	// Reads path to configs from env variable
	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
