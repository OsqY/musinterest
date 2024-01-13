package controllers

import (
	"context"
	"net/http"
	"os"
	"oscar/musinterest/initializers"
	"oscar/musinterest/models"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

func CreateDiscussion(ginContext *gin.Context) {
	var input models.DiscussionForumInput

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
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discussionForum := models.DiscussionForum{
		Auth0ID:     auth0ID,
		Title:       input.Title,
		Description: input.Description,
	}

	savedDiscussion, err := discussionForum.Save()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ginContext.JSON(http.StatusCreated, gin.H{"data": savedDiscussion})
}

func GetDiscussions(context *gin.Context) {
	var discussions []models.DiscussionForum

	result := initializers.DB.Find(&discussions)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving discussions"})
		return
	}

	context.JSON(http.StatusOK, discussions)
}
