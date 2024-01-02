package models

type CommentInput struct {
        UserId uint `json:"username" binding:"required"`
        Body string `json:"bodyOfTheComment" binding:"required"`
}
