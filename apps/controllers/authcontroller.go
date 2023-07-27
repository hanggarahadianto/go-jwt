package controllers

import (
	"fmt"
	"go-jwt/apps/models"

	"go-jwt/configuration"
	"go-jwt/initializer"
	"go-jwt/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type AuthController struct {
	DB *gorm.DB
}

func NewAuthContoller (DB *gorm.DB) AuthController{
	return AuthController{DB}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context){
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status" : "fail" , "message" : err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message" : "password do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now:= time.Now()
	newUser := models.User{
		Name		: payload.Name,
		Email		: strings.ToLower(payload.Email),
		Password	: hashedPassword,
		Role		: "user",
		Verified	: true,
		Photo		: payload.Photo,
		Provider	: "local",
		CreatedAt	: now,
		UpdatedAt	: now ,
	}

	result := configuration.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value"){
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Email already exist"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	userResponse := &models.UserResponse{
		ID: newUser.ID,
		Name: newUser.Name,
		Email: newUser.Email,
		Photo: newUser.Photo,
		Role: newUser.Role,
		Provider: newUser.Provider,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status" : "success", "data" : gin.H{"user" : userResponse}})


}


func (ac *AuthController) SignInUser (ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status" : "fail", "message" : err.Error()})
		return
	}

	var user models.User
	result := configuration.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid email or password"})
		return
	}

	if err := utils.VerifiedPassword(user.Password, payload.Password); err !=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "invalid email or password"})
		return
	}

	config, _:= initializer.LoadConfig(".")

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status" : "failed", "message" : err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message" : err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status" : "success", "access_token" : access_token})


}

func (ac *AuthController) RefreshToken(ctx *gin.Context) {
	message := "can not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")


	
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status" 	: "fail",
			"message"	: message,
		})
		return
	}
	config, _ := initializer.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	var user models.User
	result := configuration.DB.First(&user, "id = ?", fmt.Sprint(sub) )
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
	}


	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
	}

	

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"access_token" : access_token,
	})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context){
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)
}

