package main

import (
	"healthy/config"
	"healthy/cron"
	"healthy/infrastructure"
	"log"
)

func main() {
	config.LoadConfig()

	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	err := cron.Cleanup()
	if err != nil {
		log.Fatal(err)
	}
}
