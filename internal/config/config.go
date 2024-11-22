package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Application ApplicationConfig `yaml:"application" env-required:"true"`
	Database    DatabaseConfig    `yaml:"database" env-required:"true"`
	Bot         BotConfig         `yaml:"bot" env-required:"true"`
	Cryptomus   CryptomusConfig   `yaml:"cryptomus" env-required:"true"`
}

type ApplicationConfig struct {
	Port         uint32        `yaml:"port" env-required:"true"`
	Host         string        `yaml:"host" env-required:"true"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env-default:"10s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env-default:"2m"`
	Environment  string        `yaml:"environment" env-required:"true"`
	ClientUrl    string        `yaml:"client_url" env-required:"true"`
}

type BotConfig struct {
	Token             string        `yaml:"token" env-required:"true"`
	TelegramUrl       string        `yaml:"telegram_url" env-required:"true"`
	BotApiSecretToken string        `yaml:"bot_api_secret_token" env-required:"true"`
	RequestTimeout    time.Duration `yaml:"request_timeout" env-default:"6s"`
}

type CryptomusConfig struct {
	Url            string        `yaml:"url" env-required:"true"`
	MerchantId     string        `yaml:"merchant_id" env-required:"true"`
	RequestTimeout time.Duration `yaml:"request_timeout" env-default:"6s"`
	ApiKey         string        `yaml:"api_key" env-required:"true"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     uint32 `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	SslMode  *bool  `yaml:"ssl_mode" env-required:"true"`
}

func (c *DatabaseConfig) ConnectOptions() *pgxpool.Config {

	var sslMode string

	if *c.SslMode {
		sslMode = "require"
	} else {
		sslMode += "disable"
	}

	database_url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Database, sslMode)

	dbConfig, err := pgxpool.ParseConfig(database_url)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = int32(4)
	dbConfig.MinConns = int32(0)
	dbConfig.MaxConnLifetime = time.Hour
	dbConfig.MaxConnIdleTime = time.Minute * 30
	dbConfig.HealthCheckPeriod = time.Minute
	dbConfig.ConnConfig.ConnectTimeout = time.Second * 5

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
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
