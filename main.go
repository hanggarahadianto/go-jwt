package main

import (
	"go-jwt/apps/controllers"
	"go-jwt/apps/controllers/postcontroller"
	"go-jwt/apps/routes"
	"go-jwt/initializer"

	"go-jwt/configuration"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var pf = fmt.Printf

var (
	server				*gin.Engine
	AuthController		controllers.AuthController
	AuthRouteController	routes.AuthRouteController
)


func init(){

	AuthController = controllers.NewAuthContoller(configuration.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)



	server = gin.Default()


}


func main(){

	
	configuration.ConnectDatabase()

	handleRouter()

	config, err := initializer.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "welcome"
		ctx.JSON(http.StatusOK, gin.H{
			"status" : "success",
			"message" : message,
		})
	})

	AuthRouteController.AuthRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))

	
}

func handleRouter(){
	r := gin.Default()
	configuration.ConnectDatabase()
	gin.SetMode(gin.ReleaseMode)

	 r.POST("/post", postcontroller.CreatePost)

	 r.GET("/post", postcontroller.FindPosts)

	 r.GET("/post/:id", postcontroller.ById)

	 r.POST("/user", AuthController.SignUpUser)
	log.Fatal(http.ListenAndServe(":9090", r))
}
