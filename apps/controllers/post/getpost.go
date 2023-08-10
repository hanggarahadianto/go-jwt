package controllers

import (
	"go-jwt/apps/models"
	"go-jwt/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *PostController) GetPosts(c *gin.Context){
	var postList []models.Post

	result := db.DB.Debug().Find(&postList)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "error",
			"message" : result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"data" : postList,

	})	


}