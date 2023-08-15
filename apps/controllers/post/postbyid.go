package controllers

import (
	"go-jwt/apps/models"
	"go-jwt/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *PostController) PostById(c *gin.Context){
	postId := c.Param("postId")

	var post models.Post

	result := db.DB.Debug().First(&post, "id = ?", postId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "failed",
			"message" : "post id doesn't exist",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"data" : post, 
	})
}