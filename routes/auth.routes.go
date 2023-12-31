package routes

import (
	"oscar/musinterest/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
    authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
    return AuthRouteController{authController}
}

func (routeController *AuthRouteController) AuthRoute(routerGroup *gin.RouterGroup) {
    router := routerGroup.Group("/auth")

    router.POST("/register", routeController.authController.Register)
    router.POST("/login", routeController.authController.Login)
    router.GET("/verifyEmail/:verificationCode", routeController.authController.VerifyEmail)
}
