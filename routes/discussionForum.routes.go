package routes

import (
	"oscar/musinterest/controllers"
	"oscar/musinterest/middleware"

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

	router.GET("/", dfrc.DiscussionForumController.GetDiscussions)
	router.Use(middleware.JWTAuthMiddleware())
	router.POST("/create", dfrc.DiscussionForumController.CreateDiscussion)
}
