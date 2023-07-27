package authcontroller

// import (
// 	"go-jwt/apps/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// func login(w http.ResponseWriter, r *http.Request){

// }
// func Register(c *gin.Context){

// 	var body struct{
// 		Email		string
// 		Password	string
// 	}

// 	if c.Bind(&body) !=nil{
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "failed to get body",
// 		})
// 		return
// 	}

// 	hash, err := bcrypt.GenerateFromPassword([] byte(body.Password), 10)

// 	if err!= nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error" : "failed to hash password",
// 		})

// 		return
// 	}

// 	// user:= models.User{Email : body.Email, Password: string(hash)}
// 	// result := models.ConnectDatabase().DB.Create(&user)

// 	// if result.Error != nil{
// 	// 	c.JSON(http.StatusBadRequest, gin.H{
// 	// 		"error" : "failed to create User",
// 	// 	})
// 	// 	return
// 	// }

// 	// c.JSON(http.StatusOK, gin.H{})

// }
// func Logout(w http.ResponseWriter, r *http.Request){

// }