package models

import (
	"html"
	"oscar/musinterest/musinterest/database"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Username string `gorm:"size:40;not null;unique" json:"username"`
	Password string `gorm:"size 60;not null" json:"-"`
}

func (user *User) Save() (*User, error) {
	if err := database.Database.Create(&user).Error; err != nil {
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
	if err := database.Database.Where("id = ?", id).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
