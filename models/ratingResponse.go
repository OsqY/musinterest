package models

type RatingResponse struct {
	Rate    int    `json:"rate"`
	Comment string `json:"comment"`
	UserId  uint   `json:"userId"`
	AlbumId uint   `json:"albumId"`
}
