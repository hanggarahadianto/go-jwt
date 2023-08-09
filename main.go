package main

import (

	// "go-jwt/apps/controllers/postcontroller"

	"fmt"
	"go-jwt/apps/controllers"
	"go-jwt/db"
	"go-jwt/routes"
	"go-jwt/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server					*gin.Engine
	AuthController      	controllers.AuthController
	AuthRouteController		routes.AuthRouteController
)

func init(){
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()
}

func main(){

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	db.Run()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:9090", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	
	router := server.Group("/api")
	router.GET("/testing", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)

	var serverPort = "9090"
	fmt.Print("server running on port " + serverPort)
	log.Fatal(server.Run(":"+serverPort))

	// handleRouter()

	// config, err := initializer.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("? Could not load environment variables", err)
	// }

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	// corsConfig.AllowCredentials = true

	// server.Use(cors.New(corsConfig))

	// router := server.Group("/api")
	// router.GET("/healthchecker", func(ctx *gin.Context) {
	// 	message := "welcome"
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"status" : "success",
	// 		"message" : message,
	// 	})
	// })

	// AuthRouteController.AuthRoute(router)
	// log.Fatal(server.Run(":" + config.ServerPort))

	
}

// func handleRouter(){
// 	r := gin.Default()
// 	configuration.ConnectDatabase()
// 	gin.SetMode(gin.ReleaseMode)

// 	 r.POST("/post", postcontroller.CreatePost)

// 	 r.GET("/post", postcontroller.FindPosts)

// 	 r.GET("/post/:id", postcontroller.ById)

// 	 r.POST("/user", AuthController.SignUpUser)
// 	log.Fatal(http.ListenAndServe(":9090", r))
// }
