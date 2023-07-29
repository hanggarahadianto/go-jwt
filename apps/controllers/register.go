package controllers

import (
	"fmt"
	"go-jwt/apps/models"
	"go-jwt/db"
	"time"

	"go-jwt/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type AuthController struct {
	DB *gorm.DB
}

func NewAuthContoller (DB *gorm.DB) AuthController{
	return AuthController{DB}
}


func Register(c *gin.Context){
	fmt.Println("this is register")
	var registerData models.RegisterInput

	err:= c.ShouldBindJSON(&registerData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	hashPassword, err := utils.HashPassword(registerData.Password)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status" : "password failed to hashed",
			"message" : err.Error(),
		})
	}

	if registerData.Password != registerData.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail", 
			"message" : "password do not match",
		})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name: registerData.Name,
		Email: registerData.Email,
		Password: hashPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := db.DB.Create(&newUser)

		if result.Error != nil && strings.Contains(result.Error.Error(), 
		"duplicate key value"){
		c.JSON(http.StatusConflict, gin.H{
			"status": "fail", 
			"message": "Email already exist"})
		return
	} 	else if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status": "error", 
			"message": "Something bad happened",
		})
		return
	}

	userResponse := models.UserResponse{
		ID : newUser.ID,
		Name : newUser.Name,
		Email: newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "success",
		"data" : userResponse,
	})
}

