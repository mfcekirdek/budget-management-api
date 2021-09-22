package main

import (
	"gitlab.com/mfcekirdek/budget-management-api/config"
	"log"
	"os"
)

func main() {
	configFileName := os.Getenv("APP_ENV")
	if configFileName == "" {
		configFileName = "local"
	}

	conf, err := config.New(".conf", configFileName)
	checkFatalError(err)
	conf.Print()

	server := NewServer(conf)
	err = server.Start()
	checkFatalError(err)
}

func checkFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
