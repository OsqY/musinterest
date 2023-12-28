package main

import (
	"fmt"
	"log"
	"oscar/musinterest/controllers"
	"oscar/musinterest/database"
	"oscar/musinterest/middleware"
	"oscar/musinterest/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	servApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Album{})
	database.Database.AutoMigrate(&models.Rating{})
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func servApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:albumId", controllers.GetAlbumById)
	router.GET("/albums/search", controllers.GetAlbumByName)
	router.GET("/rating/:ratingId", controllers.GetRatingById)

	authenticatedRoutes := router.Group("/api")
	authenticatedRoutes.Use(middleware.JWTAuthMiddleware())
	authenticatedRoutes.POST("/rateAlbum/:albumId", controllers.AddRating)
	authenticatedRoutes.PUT("/rateAlbum/:ratingId", controllers.UpdateRating)
	authenticatedRoutes.DELETE("/rateAlbum/:ratingId", controllers.DeleteRating)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
