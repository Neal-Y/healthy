package repository

import (
	"gorm.io/gorm"
	"healthy/model/database"
	"time"
)

type FileRepository interface {
	SaveFile(file *database.File) error
	GetExpiredFiles(expirationDuration time.Duration) ([]*database.File, error)
	DeleteFile(fileID uint) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) SaveFile(file *database.File) error {
	return r.db.Save(file).Error
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
