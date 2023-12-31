package models

import (
	"html"
	"oscar/musinterest/initializers"	
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Username string `gorm:"size:40;not null;unique" json:"username"`
	Password string `gorm:"size 60;not null" json:"-"`
	Verified bool   `gorm:"not null" json:"verified"`
	Ratings  []Rating
}

func (user *User) Save() (*User, error) {
	if err := initializers.DB.Create(&user).Error; err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
