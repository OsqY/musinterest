package router

import (
	"encoding/gob"
	"oscar/musinterest/callback"
	"oscar/musinterest/controllers"
	"oscar/musinterest/login"
	"oscar/musinterest/platform/authenticator"
	"oscar/musinterest/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

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

func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"} // Añade los métodos que necesites
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}                 // Añade los encabezados que necesites
	router.Use(cors.New(config))

	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	// router.LoadHTMLGlob("web/template/*")

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:albumId", controllers.GetAlbumById)

	router.GET("/rating/:ratingId", controllers.GetRatingById)
	router.POST("/rating/:albumId", controllers.AddRating)
	router.PUT("/rating/:ratingIid", controllers.UpdateRating)
	router.DELETE("/rating/:ratingId", controllers.DeleteRating)

	router.POST("/discussion", controllers.CreateDiscussion)

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", user.Handler)
	// router.GET("/logout", logout.Handler)

	return router
}
