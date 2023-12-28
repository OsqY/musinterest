package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"oscar/musinterest/database"
	"oscar/musinterest/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAlbums(context *gin.Context) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	sortBy := context.DefaultQuery("sort_by", "title")
	order := context.DefaultQuery("order", "asc")

	offset := (page - 1) * limit

	var albums []models.Album

	result := database.Database.Order(fmt.Sprintf("%s %s", sortBy, order)).
		Offset(offset).Limit(limit).Find(&albums)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving albums"})
		return
	}

	context.JSON(http.StatusOK, albums)
}

func GetAlbumByName(context *gin.Context) {
	var albums []models.Album
	title := context.Query("title")
	result := database.Database.Where("title LIKE ?", "%"+title+"%").Find(&albums)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "that album doesn't exists"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": albums})
}

func GetAlbumById(context *gin.Context) {
	albumId, err := strconv.Atoi(context.Param("albumId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id!"})
		return
	}

	album, err := models.FindAlbumById(uint(albumId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "That id doesn't exists"})
		return
	}

	if err := database.Database.Preload("Ratings").First(&album, albumId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": album})

}
