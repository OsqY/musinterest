package controllers

import (
	"net/http"
	"oscar/musinterest/helpers"
	"oscar/musinterest/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DiscussionForumController struct {
	DB *gorm.DB
}

func NewDiscussionForumController(DB *gorm.DB) DiscussionForumController {
	return DiscussionForumController{DB}
}

func (dc *DiscussionForumController) CreateDiscussion(context *gin.Context) {
	var input models.DiscussionForumInput

	user, err := helpers.GetCurrentUser(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discussionForum := models.DiscussionForum{
		UserId:      user.ID,
		Title:       input.Title,
		Description: input.Description,
	}

	savedDiscussion, err := discussionForum.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"data": savedDiscussion})
}

func (dc *DiscussionForumController) GetDiscussions(context *gin.Context) {
	var discussions []models.DiscussionForum

	result := dc.DB.Find(&discussions)

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving discussions"})
		return
	}

	context.JSON(http.StatusOK, discussions)
}
