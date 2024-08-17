package main

import (
	"healthy/config"
	"healthy/route"
	"log"
)

func main() {
	config.LoadConfig()

	//dbErr := infrastructure.InitMySQL()
	//if dbErr != nil {
	//	log.Fatal(dbErr)
	//}

	_, err := route.InitGinServer()
	if err != nil {
		log.Fatal(err)
	}
}
