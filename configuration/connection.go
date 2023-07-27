package configuration

import (
	"fmt"
	"go-jwt/apps/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var pl = fmt.Printf

var DB *gorm.DB

func ConnectDatabase(){
	
	db, err := gorm.Open(mysql.Open("root:12345678@/go"), &gorm.Config{})
	pl("connected to database")
	if err != nil {
	  panic("failed to connect database")
	}
	// db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.User{}, &models.Post{})

	DB = db
}

