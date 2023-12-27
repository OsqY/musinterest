package models

import (
	"oscar/musinterest/musinterest/database"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Rate    float32 `gorm:"size:1;not null" json:"rating"`
	UserID  uint
	AlbumID uint
}

func (rating *Rating) Save() (*Rating, error) {
	if err := database.Database.Create(&rating).Error; err != nil {
		return &Rating{}, err
	}

	return rating, nil
}
