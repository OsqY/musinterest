package models

type DiscussionForumInput struct {
        Title string `json:"title" binding:"required,min=10, max=100"`
} 
