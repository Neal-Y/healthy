package cron

import (
	"healthy/infrastructure"
	"healthy/repository"
	"healthy/service"
	"log"
	"time"
)

func Cleanup() (err error) {
	fileRepo := repository.NewFileRepository(infrastructure.Db)
	fileService := service.NewFileService(fileRepo)

	err = fileService.CleanupExpiredFiles(24 * time.Hour)
	if err != nil {
		log.Fatalf("Error during cleanup: %v", err)
	}

	log.Println("Cleanup successful")
	return
}
