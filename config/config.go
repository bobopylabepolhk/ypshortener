package config

import (
	"flag"
	"fmt"
)

var PORT = 8080
var APIURL = fmt.Sprintf("localhost:%d", PORT)
var BASEURL = fmt.Sprintf("http://%s", APIURL)

func InitFromCLI() {
	flag.StringVar(&APIURL, "a", APIURL, "api service address")
	flag.StringVar(&BASEURL, "b", BASEURL, "shortURL address")
	flag.Parse()
}
