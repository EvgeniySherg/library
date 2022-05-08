package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Configuration struct {
	DatabaseSource DatabaseSource
}

type DatabaseSource struct {
	Driver        string
	User          string
	Password      string
	ServerAddress string
	DatabaseName  string
}

func GetConfiguration() Configuration {
	s := flag.String("filePath", "", "")
	flag.Parse()
	file, _ := os.Open(*s)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
} //через конфигуратор
//func GetConfiguration() Configuration {
//	file, _ := os.Open("config.json")
//	defer file.Close()
//	decoder := json.NewDecoder(file)
//	configuration := Configuration{}
//	err := decoder.Decode(&configuration)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return configuration
//}
