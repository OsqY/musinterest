package models


type Comment struct {
        gorm.Model
        Username uint `json:"username" binding:"required"`
        Body string `json:"bodyOfTheComment" binding:"required"`
}
