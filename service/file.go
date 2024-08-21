package service

import (
	"fmt"
	"healthy/config"
	"healthy/constant"
	"healthy/model/database"
	"healthy/repository"
	"healthy/util"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileService interface {
	UploadFileLogic(filePath string) (*database.File, error)
	CleanupExpiredFiles(expirationDuration time.Duration) error
	ProcessAndAnalyzeImage(file *multipart.FileHeader) (string, error)
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

	return s.fileRepo.UpsertFile(file)
}

func (s *fileService) CleanupExpiredFiles(expirationDuration time.Duration) error {
	files, err := s.fileRepo.GetExpiredFiles(expirationDuration)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(file.FilePath)
		if err != nil {
			return err
		}
		err = s.fileRepo.DeleteFile(file.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *fileService) ProcessAndAnalyzeImage(file *multipart.FileHeader) (string, error) {
	saveDir := "uploads"
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Error creating save directory: %v", err)
	}

	savePath := filepath.Join(saveDir, file.Filename)
	err = s.saveUploadedFile(file, savePath)
	if err != nil {
		return "", fmt.Errorf("Error saving file: %v", err)
	}

	fileModel, err := s.UploadFileLogic(savePath)
	if err != nil {
		return "", fmt.Errorf("Error saving file info: %v", err)
	}

	fileURL := fmt.Sprintf("%s/%s", config.AppConfig.WebURL, fileModel.FilePath)

	gptResponse, err := callChatGPT(fileURL)
	if err != nil {
		return "", fmt.Errorf("Error calling OpenAI API: %v", err)
	}

	return gptResponse, nil
}

func callChatGPT(fileURL string) (string, error) {
	request := util.GPTRequest{
		Model:         "gpt-4o",
		SystemMessage: constant.SystemMsg,
		UserMessage:   constant.UserMsg,
		ImageURL:      fileURL,
		MaxTokens:     300,
	}

	response, err := request.SendRequest()
	if err != nil {
		return "", err
	}

	return response, nil
}

func (s *fileService) saveUploadedFile(file *multipart.FileHeader, savePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
