package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port    int    `env:"PORT" env-default:"8080"`
	ApiURL  string `env:"SERVER_ADDRESS" env-default:"localhost:8080"`
	BaseURL string `env:"BASE_URL" env-default:"http://localhost:8080"`
}

var Cfg Config

func initFromCLI() {
	flag.StringVar(&Cfg.ApiURL, "a", Cfg.ApiURL, "api service address")
	flag.StringVar(&Cfg.BaseURL, "b", Cfg.BaseURL, "shortURL address")
	flag.Parse()
}

func initFromEnv() {
	err := cleanenv.ReadEnv(&Cfg)
	if err != nil {
		panic(fmt.Errorf("failed to read config %v", err))
	}
}

func InitConfig() {
	initFromCLI()
	initFromEnv()
}
