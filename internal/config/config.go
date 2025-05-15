package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP struct {
		Host             string        `yaml:"host"`
		Port             string        `yaml:"port"`
		Network          string        `yaml:"network"`
		ReadTimeout      time.Duration `yaml:"read_timeout"`
		WriteTimeout     time.Duration `yaml:"write_timeout"`
		GracefullTimeout time.Duration `yaml:"gracefull_timeout"`
		CORS             struct {
			AllowedOrigins   []string `yaml:"allowedOrigins"`
			AllowedMethods   []string `yaml:"allowedMethods"`
			AllowedHeaders   []string `yaml:"allowedHeaders"`
			AllowCredentials bool     `yaml:"allowCredentials"`
		} `yaml:"cors"`
	} `yaml:"http"`
	PostgreSQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgre"`
}

const configPath = "configs/config.local.yaml"

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			log.Fatalf("error: ошибка при чтении конфига %v", err)
		}
	})

	return instance
}
