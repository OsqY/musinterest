package models

import (
	"oscar/musinterest/initializers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"size:255;not null;unique" json:"email"`
	Auth0ID     string `gorm:"unique"`
	Name        string
	Discussions []DiscussionForum `gorm:"foreignKey:UserID"`
	Comments    []Comment         `gorm:"foreignKey:UserID"`
	Ratings     []Rating          `gorm:"foreignKey:UserID"`
}

func (user *User) Save() (*User, error) {
	if err := initializers.DB.Create(&user).Error; err != nil {
		return &User{}, err
	}

	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	if err := initializers.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func FindByUsername(username string) (User, error) {
	var user User
	if err := initializers.DB.Where("username = ?", username).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
