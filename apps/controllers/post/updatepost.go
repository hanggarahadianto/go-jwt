package controllers

import (
	"go-jwt/apps/models"
	"go-jwt/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (pc *PostController) UpatePost(c *gin.Context){

	postId := c.Param("postId")
	currentUser := c.MustGet("currentUser").(models.User)

	var postData *models.UpdatePost
	err := c.ShouldBindJSON(&postData)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status" : "failed",
			"message" : err.Error(),
		})
		return
	}

	var updatePost models.Post
	result := db.DB.Debug().First(&updatePost, "id = ?", postId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status" : "failed",
			"message" : "post id not found",
		})
		return
	}

	now := time.Now()
	postToUpdate := models.Post{
		Title: postData.Title,
		Content: postData.Content,
		User: currentUser.ID,
		CreatedAt: updatePost.CreatedAt,
		UpdatedAt : now,
	}

	db.DB.Debug().Model(&updatePost).Updates(postToUpdate)
	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"data" : updatePost,
	})


}