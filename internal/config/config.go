package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Application ApplicationConfig `yaml:"application" env-required:"true"`
	Database    DatabaseConfig    `yaml:"database" env-required:"true"`
}

type ApplicationConfig struct {
	Port uint32 `yaml:"port" env-required:"true"`
	Host string `yaml:"host" env-required:"true"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     uint32 `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	SslMode  *bool  `yaml:"ssl_mode" env-required:"true"`
}

func New() *Config {
	environment := os.Getenv("ENV")

	if strings.TrimSpace(environment) == "" {
		environment = "local"
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current working directory")
	}

	configPath := path.Join(cwd, "configs", fmt.Sprintf("%s.yaml", environment))

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Config file not found")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("Failed to read config:", err)
	}

	return &cfg
}
