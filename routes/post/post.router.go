package routes

import (
	controllers "go-jwt/apps/controllers/post"
	"go-jwt/apps/middlewares"

	"github.com/gin-gonic/gin"
)

type PostRouteController struct{
	postController controllers.PostController
}

func NewRoutePostController(postController controllers.PostController) PostRouteController{
	return PostRouteController{postController}
}

func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup){
	
	r := rg.Group("post")

	r.Use(middlewares.DeserializeUser())
	r.POST("/", pc.postController.CreatePost)
	r.GET("/", pc.postController.GetPosts)
	r.DELETE("/:postId", pc.postController.DeletePost)

}