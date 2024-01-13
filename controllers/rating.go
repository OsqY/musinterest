package controllers

import (
	"context"
	"errors"
	"net/http"
	"os"
	"oscar/musinterest/initializers"
	"oscar/musinterest/models"
	"strconv"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRating(ginContext *gin.Context) {
	var input models.RatingInput

	authHeader := ginContext.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	provider, err := oidc.NewProvider(context.Background(), os.Getenv("AUTH0_DOMAIN"))
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}
	verifier := provider.Verifier(oidcConfig)

	idToken, err := verifier.Verify(context.Background(), tokenString)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var claims struct {
		Sub string `json:"sub"`
	}
	if err := idToken.Claims(&claims); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth0ID := claims.Sub
	if err := ginContext.ShouldBindJSON(&input); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
		return
	}

	albumId, err := strconv.Atoi(ginContext.Param("albumId"))

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album id"})
		return
	}

	input.UserId = auth0ID
	input.AlbumId = uint(albumId)
	album, err := models.FindAlbumById(uint(albumId))

	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "That album doesn't exists"})
		return
	}

	if album.ID == 0 {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "That album doesn't exists"})
		return
	}

	if err := initializers.DB.Preload("Ratings").First(album, uint(albumId)).Error; err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "That album doesn't exists"})
		return
	}

	rating := models.Rating{
		Rate:    input.Rate,
		Comment: input.Comment,
		Auth0ID: input.UserId,
		AlbumId: input.AlbumId,
	}

	_, err = rating.Save()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"data": rating})
}

func GetRatingById(ginContext *gin.Context) {

	ratingId, err := strconv.Atoi(ginContext.Param("ratingId"))
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}

	var rating models.Rating
	if err := initializers.DB.First(&rating, ratingId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ginContext.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{"data": rating})

}

func UpdateRating(ginContext *gin.Context) {

	authHeader := ginContext.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	provider, err := oidc.NewProvider(context.Background(), os.Getenv("AUTH0_DOMAIN"))
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}
	verifier := provider.Verifier(oidcConfig)

	idToken, err := verifier.Verify(context.Background(), tokenString)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var claims struct {
		Sub string `json:"sub"`
	}
	if err := idToken.Claims(&claims); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth0ID := claims.Sub
	ratingId, err := strconv.Atoi(ginContext.Param("ratingId"))
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}
	rating, err := models.FindRatingById(uint(ratingId), auth0ID)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}
	var newRating models.RatingInput
	if err := ginContext.ShouldBindJSON(&newRating); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid put request"})
		return
	}
	rating.Rate = newRating.Rate
	rating.Comment = newRating.Comment
	if err := initializers.DB.Model(&rating).Updates(models.Rating{Rate: rating.Rate, Comment: rating.Comment}).Error; err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong happened!"})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{"data": rating})
}

func DeleteRating(ginContext *gin.Context) {
	authHeader := ginContext.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	provider, err := oidc.NewProvider(context.Background(), os.Getenv("AUTH0_DOMAIN"))
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}
	verifier := provider.Verifier(oidcConfig)

	idToken, err := verifier.Verify(context.Background(), tokenString)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var claims struct {
		Sub string `json:"sub"`
	}
	if err := idToken.Claims(&claims); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	auth0ID := claims.Sub
	ratingId, err := strconv.Atoi(ginContext.Param("ratingId"))
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Id not valid"})
		return
	}
	rating, err := models.FindRatingById(uint(ratingId), auth0ID)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ratingId"})
		return
	}
	if err := initializers.DB.Where("id = ?", uint(ratingId)).Delete(&rating).Error; err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong happened!"})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{"status": "Deleted successfully!"})
}
