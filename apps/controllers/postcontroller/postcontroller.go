package postcontroller

// import (
// 	"fmt"
// 	"go-jwt/apps/models"
// 	"go-jwt/configuration"

// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type PostController struct {
// 	DB *gorm.DB
// }

// func CreatePost(c *gin.Context) {
//     var input models.Post
//     if err := c.ShouldBindJSON(&input);
//     err != nil {
//         c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     post := models.Post{Title: input.Title, Content: input.Content}
//     configuration.DB.Create(&post)

//     c.JSON(http.StatusOK, gin.H{"data": post})
// }

// func FindPosts(c *gin.Context) {
//     var posts []models.Post
//     result := configuration.DB.Debug().Find(&posts).Error

//     if result == nil {
//        fmt.Println("data kosong")
//     }

//     c.JSON(http.StatusOK, gin.H{"message": "success","data": posts})
// }

// func ById(c *gin.Context) {  // Get model if exist
//     var post models.Post

//     id:= c.Param("id")
//     configuration.DB.First(&post, id)

//     c.JSON(http.StatusOK, gin.H{"data": post})

//   }
