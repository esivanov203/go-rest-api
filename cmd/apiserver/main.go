package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/esivanov203/go-rest-api/internal/app/apiserver"
	"log"
)

var cfgPath string

func init() {
	flag.StringVar(&cfgPath, "cfgfile", "./configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse() // Параметры приложения при запуске

	cfg := apiserver.NewConfig()
	_, err := toml.DecodeFile(cfgPath, cfg)
	if err != nil {
		log.Fatal("Config reading error: " + err.Error())
	}
	s := apiserver.New(cfg)
	if err = s.Run(); err != nil {
		log.Fatal("Server runnig error:" + err.Error())
	}
}
