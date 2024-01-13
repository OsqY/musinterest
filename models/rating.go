package models

import (
	"oscar/musinterest/initializers"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Rate    float32 `gorm:"size:1;not null" json:"rating"`
	Comment string  `gorm:"size255;not null" json:"comment"`
	Auth0ID string  `gorm:"type:varchar(255)"`
	UserID  uint
	AlbumId uint
}

func (rating *Rating) Save() (*Rating, error) {
	if err := initializers.DB.Create(&rating).Error; err != nil {
		return &Rating{}, err
	}

	return rating, nil
}

func FindRatingById(ratingId uint, auth0Id string) (*Rating, error) {
	var rating Rating
	if err := initializers.DB.Where("id = ? AND user_id = ?", ratingId, auth0Id).First(&rating).Error; err != nil {
		return nil, err
	}
	return &rating, nil
}
