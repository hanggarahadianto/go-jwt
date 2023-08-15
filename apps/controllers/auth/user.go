package controllers

import (
	"go-jwt/apps/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func newUserController(DB *gorm.DB) UserController{
// 		return UserController{DB}
// }

func (ac *AuthController) GetMe(ctx *gin.Context){
	currentUser:= ctx.MustGet("currentUser").(models.User)

	userResponse:= &models.UserResponse{
		ID: currentUser.ID,
		Name: currentUser.Name,
		Email: currentUser.Email,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success", 
		"data": gin.H{"user": userResponse}})

}