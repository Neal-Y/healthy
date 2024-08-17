package main

import (
	"healthy/infrastructure"
	"healthy/repository"
	"healthy/service"
	"log"
	"time"
)

func main() {
	dbErr := infrastructure.InitMySQL()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	fileRepo := repository.NewFileRepository(infrastructure.Db)
	fileService := service.NewFileService(fileRepo)

	err := fileService.CleanupExpiredFiles(24 * time.Hour)
	if err != nil {
		log.Fatalf("Error during cleanup: %v", err)
	}

	log.Println("Cleanup successful")
}
