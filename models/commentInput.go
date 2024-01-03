package models

type CommentInput struct {
  DiscussionForumId uint `json:"discussionId" binding:"required" `
  UserId uint `json:"userId" binding:"required"`
  Body string `json:"bodyOfTheComment" binding:"required"`
}
