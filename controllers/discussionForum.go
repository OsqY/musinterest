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

  if err := context.ShouldBindJSON(input).Error; err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request!"})
    return
  }

  user, err := helpers.GetCurrentUser(context)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  } 

  discussionForum := models.DiscussionForum {
    DiscussionOwner: user,
    Title: input.Title,
    Description: input.Description, 
  }

  savedDiscussion, err := discussionForum.Save()
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  context.JSON(http.StatusCreated, gin.H{"data": savedDiscussion})
}
