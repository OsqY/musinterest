package models

type AuthenticationInput struct {
	Email    string `json:"email" binding:"required,min=3,max=255"`
	Username string `json:"username" binding:"required,min=6,max=50"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}
