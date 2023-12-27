package models

import (
	"oscar/musinterest/musinterest/database"
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title       string    `gorm:"size:255;not null" json:"title"`
	Artist      string    `gorm:"size:255;not null" json:"artist"`
	ReleaseDate time.Date `gorm:"not null" json:"releaseDate"`
}

func (album *Album) Save() (*Album, error) {
	if err := database.Database.Create(&album).Error; err != nil {
		return &Album{}, err
	}
	return album, nil
}
