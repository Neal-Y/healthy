package service

import (
	"healthy/model/database"
	"healthy/repository"
	"os"
	"time"
)

type FileService interface {
	UploadFileLogic(filePath string) (*database.File, error)
	CleanupExpiredFiles(expirationDuration time.Duration) error
}

type fileService struct {
	fileRepo repository.FileRepository
}

func NewFileService(fileRepo repository.FileRepository) FileService {
	return &fileService{fileRepo: fileRepo}
}

func (s *fileService) UploadFileLogic(filePath string) (*database.File, error) {
	file := &database.File{
		FilePath:     filePath,
		UploadTime:   time.Now(),
		AccessedTime: time.Now(),
	}
	err := s.fileRepo.SaveFile(file)
	return file, err
}

func (s *fileService) CleanupExpiredFiles(expirationDuration time.Duration) error {
	files, err := s.fileRepo.GetExpiredFiles(expirationDuration)
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := os.Remove(file.FilePath); err != nil {
			return err
		}
		if err := s.fileRepo.DeleteFile(file.ID); err != nil {
			return err
		}
	}
	return nil
}
