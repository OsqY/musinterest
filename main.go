package main

import (
	"log"
	"oscar/musinterest/initializers"
	"oscar/musinterest/models"
	"oscar/musinterest/platform/authenticator"
	"oscar/musinterest/platform/router"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables ", err)
	}

	initializers.ConnectDB(&config)

}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(203)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func main() {
	_, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Couldn't load env variables")
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatal("Couldn't load the authenticator", err)
	}
	router := router.New(auth)

	initializers.DB.AutoMigrate(&models.Rating{})
	initializers.DB.AutoMigrate(&models.Album{})
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.DiscussionForum{})
	initializers.DB.AutoMigrate(&models.Comment{})

	// router.Use(CORSMiddleware())

	router.Run(":8000")
}
