package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type Config struct {
	App struct {
		Port string `yaml:"port" env:"APP_PORT"`
		Name string `yaml:"name" env:"APP_NAME"`
	} `yaml:"app"`
	CORS struct {
		AllowedMethods     []string `yaml:"allowed-methods" env:"HTTP_CORS_ALLOWED_METHODS"`
		AllowedOrigins     []string `yaml:"allowed-origins" env:"HTTP_CORS_ALLOWED_ORIGINS"`
		AllowCredentials   bool     `yaml:"allow_credentials" env:"HTTP_CORS_ALLOWED_CREDS"`
		AllowedHeaders     []string `yaml:"allowed_headers" env:"HTTP_CORS_ALLOWED_HEADERS"`
		OptionsPassthrough bool     `yaml:"options_passthrough" env:"HTTP_CORS_PASS_THROUGH"`
		ExposedHeaders     []string `yaml:"exposed_headers" env:"HTTP_CORS_EXPOSED_HEADERS"`
		Debug              bool     `yaml:"debug" env:"HTTP_CORS_DEBUG"`
	} `yaml:"cors"`
	DB struct {
		Host     string `yaml:"host" env:"DB_HOST"`
		Username string `yaml:"username" env:"DB_USERNAME"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Name     string `yaml:"name" env:"DB_NAME"`
		Port     int    `yaml:"port" env:"DB_PORT"`
	} `yaml:"db"`
	Log        zap.Config `yaml:"log"`
	Production bool       `yaml:"production" env:"PRODUCTION"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
