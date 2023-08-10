package controllers

import (
	"go-jwt/apps/models"
	"go-jwt/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

func (pc *PostController) CreatePost(c *gin.Context){

	currentUser:= c.MustGet("currentUser").(models.User)

	var postData models.CreatePostInput

	err:= c.ShouldBindJSON(&postData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now		:= time.Now()
	newPost := models.Post{
		Title: postData.Title,
		Content: postData.Content,
		User: currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := db.DB.Create(&newPost)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : "failed",
			"message" : result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : "success",
		"data" : newPost,
	})
}