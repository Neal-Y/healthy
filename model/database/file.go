package database

import (
	"time"
)

type File struct {
	ID           uint      `gorm:"primaryKey"`
	FilePath     string    `gorm:"type:varchar(255);not null"`
	UploadTime   time.Time `gorm:"not null"`
	AccessedTime time.Time `gorm:"not null"`
}

func (File) TableName() string {
	return "uploaded_files"
}
