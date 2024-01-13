package models

import (
	"oscar/musinterest/initializers"

	"gorm.io/gorm"
)

type DiscussionForum struct {
	gorm.Model
	Auth0ID     string `gorm:"type:varchar(255);unique"`
	UserID      uint
	Title       string    `gorm:"size:100; not null" json:"title" `
	Description string    `gorm:"size:255; not null" json:"description" `
	Comments    []Comment `json:"comments"`
}

func (discussion *DiscussionForum) Save() (*DiscussionForum, error) {
	if err := initializers.DB.Create(&discussion).Error; err != nil {
		return &DiscussionForum{}, err
	}
	return discussion, nil
}

func FindDiscussionById(discussionId uint) (*DiscussionForum, error) {
	var discussion DiscussionForum
	if err := initializers.DB.Where("id = ?", discussionId).Find(&discussion).Error; err != nil {
		return &DiscussionForum{}, err
	}
	return &discussion, nil
}
