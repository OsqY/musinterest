package models

type DiscussionForumInput struct {
  Title string `json:"title" binding:"required,min=10, max=100"`
  Description string `json:"description" binding:"required, min=10, max=255"`
  UserId uint `json:"userId" binding:"required"`
} 
