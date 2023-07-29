package controllers

import (
	"fmt"
	"go-jwt/apps/middlewares"
	"go-jwt/apps/models"
	"go-jwt/db"
	"go-jwt/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context){
	fmt.Println("this is login")
	var loginData *models.LoginInput

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "fail", 
			"message" : err.Error(),
		})
		return
	}

	var user models.User

	result := db.DB.First(&user, "email = ?", strings.ToLower(loginData.Email))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", 
			"message": "invalid email or password"})
		return 
	}

	if err := utils.VerifiedPassword(user.Password, loginData.Password); err !=nil{
		c.JSON(http.StatusForbidden, gin.H{
			"status": "fail", 
			"message": "invalid email or password"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"data" : user.ID,
	// })


	config, _:= utils.LoadConfig(".")

	access_token, err := middlewares.GenerateToken(
		config.AccessTokenExpiresIn,
		user.ID,
		config.AccessTokenPrivateKey,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "failed", 
			"message" : err.Error()})
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status" : "success", 
		"access_token" : access_token,
	})


// }

// func (ac *AuthController) RefreshToken(ctx *gin.Context) {
// 	message := "can not refresh access token"

// 	cookie, err := ctx.Cookie("refresh_token")

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"status" 	: "fail",
// 			"message"	: message,
// 		})
// 		return
// 	}
// 	config, _ := utils.LoadConfig(".")

// 	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)

// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"status" : "fail",
// 			"message" : err.Error(),
// 		})
// 		return
// 	}

// 	var user models.User
// 	result := db.DB.First(&user, "id = ?", fmt.Sprint(sub) )
// 	if result.Error != nil {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"status" : "fail",
// 			"message" : err.Error(),
// 		})
// 	}


// 	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
// 	if err != nil{
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 			"status" : "fail",
// 			"message" : err.Error(),
// 		})
// 	}

// 	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
// 	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"status" : "success",
// 		"access_token" : access_token,
// 	})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context){
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)
}
