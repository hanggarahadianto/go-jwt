package routes

import (
	"go-jwt/apps/controllers"
	"go-jwt/apps/middlewares"

	"github.com/gin-gonic/gin"
)


type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController{
	return AuthRouteController{authController}
}

func (rc *AuthRouteController)AuthRoute(rg *gin.RouterGroup){
	router := rg.Group("/auth")


	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/ogin", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshToken)
	router.GET("/logout", middlewares.DeserializeUser(), rc.authController.LogoutUser)
}