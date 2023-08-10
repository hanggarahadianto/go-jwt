package controllers

import (
	"go-jwt/apps/models"
	"go-jwt/db"
	"go-jwt/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
)

func (rc *AuthController)ForgotPassword(c *gin.Context){
	var forgotPasswordData models.ForgotPasswordInput

	err:= c.ShouldBindJSON(&forgotPasswordData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "failed",
			"message" : err.Error(),
		})
	}
	message := "You will receive reset email if email exist"


	var user models.User
	result:= db.DB.First(&user, "email = ?", strings.ToLower(forgotPasswordData.Email))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "failed",
			"message" : "email not found",
		})
		return
	}

	config, _:= utils.LoadConfig(".")

	// config, err := utils.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("could not load config in forgot email ", err)
	// }

	resetToken:= randstr.String(20)

	passwordResetToken:= utils.Encode(resetToken)
	user.PasswordResetToken = passwordResetToken
	user.PasswordResetAt = time.Now().Add(time.Minute * 15)
	db.DB.Save(&user)

	var firstName = user.Name

	if strings.Contains(firstName, " "){
		firstName = strings.Split(firstName, " ")[1]
	}

	emailData := utils.EmailData{
		URL		    : config.ClientOrigin + "/forgotPassword/" + resetToken,
		FirstName	: firstName,
		Subject		: "Your password reset token (valid for 10 minutes)",
	}

	utils.SendEmail(&user, &emailData, "resetPassword.html")

	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"message" : message,
	})

}