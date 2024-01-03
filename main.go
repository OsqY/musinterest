package main

import (
	"log"
	"oscar/musinterest/controllers"
	"oscar/musinterest/initializers"
	"oscar/musinterest/routes"
	"oscar/musinterest/models"

	"github.com/gin-gonic/gin"

    )
var   (

  server		  *gin.Engine
  AuthController	  controllers.AuthController
  AuthRouteController	  routes.AuthRouteController
	
  AlbumController	  controllers.AlbumController
  AlbumRouteController  routes.AlbumRouteController
	
  RatingController 	  controllers.RatingController
  RatingRouteController routes.RatingRouteController

  DiscussionForumController controllers.DiscussionForumController
  DiscussionForumRouteController routes.DiscussionForumRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables ", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	AlbumController = controllers.NewAlbumController(initializers.DB)
	AlbumRouteController = routes.NewRouteAlbumController(AlbumController)

	RatingController = controllers.NewRatingController(initializers.DB)
	RatingRouteController = routes.NewRouteRatingController(RatingController)

  DiscussionForumController = controllers.NewDiscussionForumController(initializers.DB)
  DiscussionForumRouteController = routes.NewDiscussionForumRouteController(DiscussionForumController)

	server = gin.Default()
}

func main() {
    _, err := initializers.LoadConfig(".")
    
    if err != nil {
	log.Fatal("Couldn't load env variables")
    }
	initializers.DB.AutoMigrate(&models.Rating{})
	initializers.DB.AutoMigrate(&models.Album{})
	initializers.DB.AutoMigrate(&models.User{})
  initializers.DB.AutoMigrate(&models.DiscussionForum{})
  initializers.DB.AutoMigrate(&models.Comment{})

    router := server.Group("/")
    AuthRouteController.AuthRoute(router)
    AlbumRouteController.AlbumRoute(router)
    RatingRouteController.RatingRoute(router)

    server.Run(":8000")
}

