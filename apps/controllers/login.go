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


func (ac *AuthController)Login(c *gin.Context){
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
			"message": "email not found"})
		return 
	}

	if err := utils.VerifiedPassword(user.Password, loginData.Password); err !=nil{
		c.JSON(http.StatusForbidden, gin.H{
			"status": "fail", 
			"message": "wrong password"})
		return
	}
	
	config, _:= utils.LoadConfig(".")

	token, err := middlewares.GenerateToken(
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

	c.SetCookie("token", token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status" : "success", 
		"access_token" : token,
		"data" : loginData.Email,
	})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context){
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)
}
