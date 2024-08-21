package main

import (
	"healthy/cron"
	"healthy/infrastructure"
	"log"
)

func main() {
	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	err := cron.Cleanup()
	if err != nil {
		log.Fatal(err)
	}
}
