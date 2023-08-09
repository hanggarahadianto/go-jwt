package routes

import (
	"go-jwt/apps/controllers"
	"go-jwt/apps/middlewares"

	"github.com/gin-gonic/gin"
)


type userRouteController struct {
	userController controllers.UserController
}

func newRouteUserController(userController controllers.UserController) userRouteController{
	return userRouteController{userController}
}

func (uc *userRouteController) userRoute(rg * gin.RouterGroup){
		router := rg.Group("users")
		router.GET("/me", middlewares.DeserializeUser(), uc.userController.GetMe)
}