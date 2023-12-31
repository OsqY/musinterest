package routes

import (
	"oscar/musinterest/controllers"

	"github.com/gin-gonic/gin"
)

type AlbumRouteController struct {
    albumController controllers.AlbumController
}

func NewRouteAlbumController(albumController controllers.AlbumController) AlbumRouteController {
    return AlbumRouteController{albumController}
}

func (ac *AlbumRouteController) AlbumRoute(rg *gin.RouterGroup) {
    router := rg.Group("/albums")

    router.GET("/", ac.albumController.GetAlbums, ac.albumController.GetAlbumByName)
    router.GET("/:albumId", ac.albumController.GetAlbumById)
}
