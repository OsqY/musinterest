package controllers

import (
	"errors"
	"net/http"
	"oscar/musinterest/helpers"
	"oscar/musinterest/models"	
	"oscar/musinterest/initializers"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RatingController struct {
	DB *gorm.DB
}

func NewRatingController(DB *gorm.DB) RatingController {
	return RatingController{DB}
}


func (rc *RatingController) AddRating(context *gin.Context) {
	var input models.RatingInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	user, err := helpers.GetCurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumId, err := strconv.Atoi(context.Param("albumId"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album id"})
		return
	}

	input.UserId = user.ID
	input.AlbumId = uint(albumId)
	album, err := models.FindAlbumById(uint(albumId))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "That album doesn't exists"})
		return
	}

	if album.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "That album doesn't exists"})
		return
	}

	if err := initializers.DB.Preload("Ratings").First(album, uint(albumId)).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "That album doesn't exists"})
		return
	}

	rating := models.Rating{
		Rate:    input.Rate,
		Comment: input.Comment,
		UserId:  input.UserId,
		AlbumId: input.AlbumId,
	}

	_, err = rating.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": rating})
}

func (rc *RatingController) GetRatingById(context *gin.Context) {

	ratingId, err := strconv.Atoi(context.Param("ratingId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}

	var rating models.Rating
	if err := initializers.DB.First(&rating, ratingId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": rating})

}

func (rc *RatingController) UpdateRating(context *gin.Context) {
	currentUser, err := helpers.GetCurrentUser(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	ratingId, err := strconv.Atoi(context.Param("ratingId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}
	rating, err := models.FindRatingById(uint(ratingId), currentUser.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}
	var newRating models.RatingInput
	if err := context.ShouldBindJSON(&newRating); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid put request"})
		return
	}
	rating.Rate = newRating.Rate
	rating.Comment = newRating.Comment
	if err := initializers.DB.Model(&rating).Updates(models.Rating{Rate: rating.Rate, Comment: rating.Comment}).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong happened!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": rating})
}

func (rc *RatingController) DeleteRating(context *gin.Context) {
	currentUser, err := helpers.GetCurrentUser(context)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}
	ratingId, err := strconv.Atoi(context.Param("ratingId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Id not valid"})
		return
	}
	rating, err := models.FindRatingById(uint(ratingId), currentUser.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}
	if err := initializers.DB.Where("id = ?", uint(ratingId)).Delete(&rating).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong happened!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "Deleted successfully!"})
}
