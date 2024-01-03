package routes

import (
	"oscar/musinterest/controllers"

	"github.com/gin-gonic/gin"
)

type DiscussionForumRouteController struct {
  DiscussionForumController controllers.DiscussionForumController
} 

func NewDiscussionForumRouteController(discussionForumController controllers.DiscussionForumController) DiscussionForumRouteController {
 return DiscussionForumRouteController{discussionForumController} 
}

func (dfrc *DiscussionForumRouteController) DiscussionForumRoute(rg *gin.RouterGroup) {
  router := rg.Group("/discussions")

  router.POST("/create", dfrc.DiscussionForumController.CreateDiscussion)
}
