package routes

import (
	"go-jwt/apps/controllers"

	"github.com/gin-gonic/gin"
)


type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController{
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(r *gin.RouterGroup){
	authRoute:= r.Group("auth")

	authRoute.POST("/register", rc.authController.Register)
	authRoute.POST("/login", rc.authController.Login)

	authRoute.POST("/forgotpassword", controllers.ForgotPassword)
	authRoute.PATCH("/resetpassword/:resetToken", controllers.ResetPassword)

}