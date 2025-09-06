package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models" 
)

var DB *gorm.DB

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto Migrate models
	err = db.AutoMigrate(&models.Post{}, &models.User{})
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	fmt.Println("✅ Database connected and migrated!")
	DB = db
	return db
}
