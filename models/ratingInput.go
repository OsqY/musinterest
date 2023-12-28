package models

type RatingInput struct {
	Rate    float32 `json:"rating" binding:"required,min=1,max=5"`
	Comment string  `json:"comment" binding:"required,min=10,max=255"`
	UserId  uint    `json:"userId"`
	AlbumId uint    `json:"albumId"`
}
