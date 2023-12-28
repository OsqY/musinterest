package models

import (
	"oscar/musinterest/database"
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title       string    `gorm:"size:255;not null" json:"title"`
	Artist      string    `gorm:"size:255;not null" json:"artist"`
	ReleaseDate time.Time `gorm:"not null" json:"releaseDate"`
	Ratings     []Rating
}

func (album *Album) Save() (*Album, error) {
	var existingAlbum Album
	if err := database.Database.First(&existingAlbum, album.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := database.Database.Create(&album).Error; err != nil {
				return &Album{}, err
			}
		} else {
			return &Album{}, err
		}
	} else {
		if err := database.Database.Save(&album).Error; err != nil {
			return &Album{}, err
		}
	}
	return album, nil
}

func FindAlbumById(albumId uint) (*Album, error) {
	var album Album
	if err := database.Database.Where("id = ?", albumId).First(&album).Error; err != nil {
		return &Album{}, err
	}
	return &album, nil
}
