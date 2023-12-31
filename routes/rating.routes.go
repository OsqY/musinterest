package routes

import (
	"oscar/musinterest/controllers"
	"oscar/musinterest/middleware"

	"github.com/gin-gonic/gin"
)

type RatingRouteController struct {
    ratingController controllers.RatingController
}

func NewRouteRatingController(ratingController controllers.RatingController) RatingRouteController {
    return RatingRouteController{ratingController}
}

func (rc *RatingRouteController) RatingRoute(rg *gin.RouterGroup) {
    protectedRouter := rg.Group("/api")
    protectedRouter.Use(middleware.JWTAuthMiddleware())
    protectedRouter.POST("/rating/:albumId", rc.ratingController.AddRating)
    protectedRouter.PUT("/rating/:ratingId", rc.ratingController.UpdateRating)
    protectedRouter.DELETE("/rating/:ratingId", rc.ratingController.DeleteRating)

    router := rg.Group("/ratings")
    router.GET("/:ratingId", rc.ratingController.GetRatingById)

}
