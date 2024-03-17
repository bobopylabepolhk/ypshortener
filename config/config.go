package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port           int    `env:"PORT" env-default:"8080"`
	APIURL         string `env:"SERVER_ADDRESS" env-default:"localhost:8080"`
	BaseURL        string `env:"BASE_URL" env-default:"http://localhost:8080"`
	Debug          bool   `env:"DEBUG" env-default:"false"`
	URLStoragePath string `env:"FILE_STORAGE_PATH" env-default:"/tmp/short-url-db.json"`
}

var Cfg Config

func initFromCLI() {
	flag.StringVar(&Cfg.APIURL, "a", Cfg.APIURL, "api service address")
	flag.StringVar(&Cfg.BaseURL, "b", Cfg.BaseURL, "shortURL address")
	flag.StringVar(&Cfg.URLStoragePath, "f", Cfg.URLStoragePath, "short url db path")
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
