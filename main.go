package main

import (
	"fmt"
	authcontrollers "go-jwt/apps/controllers/auth"
	postcontrollers "go-jwt/apps/controllers/post"
	"go-jwt/db"
	authroutes "go-jwt/routes/auth"
	postroutes "go-jwt/routes/post"
	"go-jwt/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server					*gin.Engine

	AuthController      	authcontrollers.AuthController
	AuthRouteController		authroutes.AuthRouteController
		

	PostController			postcontrollers.PostController
	PostRouteController		postroutes.PostRouteController

)

func init(){
	AuthRouteController = authroutes.NewAuthRouteController(AuthController)
	PostRouteController = postroutes.NewRoutePostController(PostController)
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()
}

func main(){

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}

	db.InitializeDb(&config)

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
	PostRouteController.PostRoute(router)

///////////-----------------------------


	perkalian := utils.Perkalian(4,5)
	fmt.Println(perkalian)

	title:= "belajar looping"
	for _, letter := range title{
		fmt.Println("letter :", string(letter) )
	}

	for i := 0; i < 5; i++ {
		fmt.Println("saya belajar golang")
	}

	//////////-----------------------

	fmt.Println("server running on port " + config.ServerPort)
	log.Fatal(server.Run(":" + config.ServerPort))


}


func cekumurbocil(age int){
	
	if age >= 12 {
		fmt.Println("boleh main game")
	} else if age <= 12 {
		fmt.Println("lu masih bocil tolol")
	}

}