package db

import (
	"fmt"
	"go-jwt/apps/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDb(dbConfig DbConfig){
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	dbConfig.DBHost,
	dbConfig.DBPort,
	dbConfig.DBUser,
	dbConfig.DBPassword,
	dbConfig.DBName,
)
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	log.Fatal("Failed to connect to the Database")
	}

	fmt.Println("konek ke pg")
	
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.AutoMigrate(
		&models.User{},
		&models.Post{},
	)

	DB = db
}