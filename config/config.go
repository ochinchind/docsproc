package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type (
	// Config -.
	Config struct {
		App               `yaml:"app"`
		HTTP              `yaml:"http"`
		Log               `yaml:"logger"`
		PG                `yaml:"postgres"`
		GoogleLoginConfig oauth2.Config
		RMQ               `yaml:"rabbitmq"`
		Redis             `yaml:"redis"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		//PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		PGURL string `env-required:"true"                 env:"PG_URL"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env-required:"true" yaml:"rpc_server_exchange" env:"RMQ_RPC_SERVER"`
		ClientExchange string `env-required:"true" yaml:"rpc_client_exchange" env:"RMQ_RPC_CLIENT"`
		RMQURL         string `env-required:"true"                            env:"RMQ_URL"`
	}

	// Redis -.
	Redis struct {
		RedisURL      string `env-required:"true" yaml:"url" env:"REDIS_URL"`
		RedisPassword string `yaml:"password" env:"REDIS_PASSWORD"`
	}
)

var Cfg Config

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	AppConfig := &Config{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}

	err = cleanenv.ReadConfig("./config/config.yml", AppConfig)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(AppConfig)
	if err != nil {
		return nil, err
	}

	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	Cfg = *AppConfig

	return AppConfig, nil
}

func GoogleConfig() oauth2.Config {
	return Cfg.GoogleLoginConfig
}
