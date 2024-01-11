package main

import (
	"log"
	"oscar/musinterest/controllers"
	"oscar/musinterest/initializers"
	"oscar/musinterest/models"
	"oscar/musinterest/routes"

	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	AlbumController      controllers.AlbumController
	AlbumRouteController routes.AlbumRouteController

	RatingController      controllers.RatingController
	RatingRouteController routes.RatingRouteController

	DiscussionForumController      controllers.DiscussionForumController
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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

	server.Use(CORSMiddleware())
	router := server.Group("/")
	router.Use(CORSMiddleware())

	AuthRouteController.AuthRoute(router)
	AlbumRouteController.AlbumRoute(router)
	RatingRouteController.RatingRoute(router)
	DiscussionForumRouteController.DiscussionForumRoute(router)

	server.Run(":8000")
}
