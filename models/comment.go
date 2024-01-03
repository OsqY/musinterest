package models

import (
	"oscar/musinterest/initializers"

	"gorm.io/gorm"
)

type Comment struct {
  gorm.Model
  DiscussionForumId uint 
  UserId uint
  Body string `json:"bodyOfTheComment" binding:"required"`
}

func (comment *Comment) Save() (*Comment, error) {
  if err := initializers.DB.Create(&comment).Error; err != nil {
    return &Comment{}, err
  }
  return comment, nil
}

func FindCommentById(commentId uint) (*Comment, error) {
  var comment Comment
  if err := initializers.DB.Where("id = ?", commentId).First(&comment).Error; err != nil {
    return &Comment{}, err
  }
  return &comment, nil
}

