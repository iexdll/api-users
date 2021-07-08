package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"restAPI/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {

	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

}

//Создание миграции
//C:\Users\SmirnovA\go\migrate.exe create -ext sql -dir migrations create_users
//C:\Users\SmirnovA\go\migrate.exe -path migrations -database "postgres://192.168.99.100:5432/restapi_dev?sslmode=disable&user=postgres&password=qwerty" up
