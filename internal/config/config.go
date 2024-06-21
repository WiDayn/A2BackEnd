package config

import (
	"A2BackEnd/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func SetupDatabase() *gorm.DB {
	dsn := "root:localmysql@tcp(127.0.0.1:3306)/a2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return db
}
