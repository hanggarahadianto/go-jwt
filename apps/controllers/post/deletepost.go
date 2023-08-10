package controllers

import (
	"go-jwt/apps/models"
	"go-jwt/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *PostController) DeletePost(c *gin.Context){
	postId:= c.Param("postId")

	result:= db.DB.Debug().Delete(&models.Post{}, "id = ?",  postId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status" : "faile",
			"message" : "post id not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : "success delete",
	})

}