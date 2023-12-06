package config

import (
	"flag"
	"fmt"
	"os"
)

func getEnv(key string, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return defaultValue
}

var PORT = getEnv("PORT", "8080")
var APIURL = getEnv("SERVER_ADDRESS", fmt.Sprintf("localhost:%s", PORT))
var BASEURL = getEnv("BASE_URL", fmt.Sprintf("http://%s", APIURL))

func InitFromCLI() {
	flag.StringVar(&APIURL, "a", APIURL, "api service address")
	flag.StringVar(&BASEURL, "b", BASEURL, "shortURL address")
	flag.Parse()
}
