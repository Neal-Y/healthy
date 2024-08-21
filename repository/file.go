package repository

import (
	"gorm.io/gorm"
	"healthy/model/database"
	"time"
)

type FileRepository interface {
	UpsertFile(file *database.File) (*database.File, error)
	GetExpiredFiles(expirationDuration time.Duration) ([]*database.File, error)
	DeleteFile(fileID uint) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) UpsertFile(file *database.File) (*database.File, error) {
	var existingFile database.File
	err := r.db.Where("file_path = ?", file.FilePath).First(&existingFile).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = r.db.Create(file).Error
			return file, err
		}
		return nil, err
	}

	existingFile.AccessedTime = file.AccessedTime
	err = r.db.Save(&existingFile).Error
	return &existingFile, err
}

func (r *fileRepository) GetExpiredFiles(expirationDuration time.Duration) ([]*database.File, error) {
	var files []*database.File
	expirationTime := time.Now().Add(-expirationDuration)
	err := r.db.Where("accessed_time < ?", expirationTime).Find(&files).Error
	return files, err
}

func (r *fileRepository) DeleteFile(fileID uint) error {
	return r.db.Delete(&database.File{}, fileID).Error
}
