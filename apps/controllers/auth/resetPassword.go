package controllers

import (
	"fmt"
	"go-jwt/apps/models"
	"go-jwt/db"
	"go-jwt/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func ResetPassword(c *gin.Context){
	fmt.Println("change password")
	var resetPasswordData models.ResetPassword
	resetToken := c.Params.ByName("resetToken")

	if err:= c.ShouldBindJSON(&resetPasswordData)
	err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status" : "fail",
			"message" : err.Error(),
		})
		return
	}

	if resetPasswordData.Password != resetPasswordData.PasswordConfirm{
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "failed",
			"message" : "Password do not match",
		})
	}

	hashPassword, _ := utils.HashPassword(resetPasswordData.Password)
	passwordResetToken := utils.Encode(resetToken)


	var updateUser models.User
	result := db.DB.First(
		&updateUser, 
		"password_reset_token = ? AND password_reset_at > ?",
		passwordResetToken,
		time.Now(),
		)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message" : "reset token has expired",
		})
		return
	}

	updateUser.Password = hashPassword
	updateUser.PasswordResetToken = ""
	db.DB.Save(&updateUser)

	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK,gin.H{
		"status" : "success",
		"message" : "update password success",
	})
}