package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/PetJs/blog-backend/internal/models"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	rawDSN := os.Getenv("DB_URL") // set this in your env or Render dashboard
	if rawDSN == "" {
		log.Fatal("❌ DB_URL environment variable not set")
	}

	// Convert mysql://user:pass@host:port/dbname -> GORM DSN
	trimmed := strings.TrimPrefix(rawDSN, "mysql://")
	parts := strings.SplitN(trimmed, "@", 2)
	if len(parts) != 2 {
		log.Fatal("❌ Invalid DB_URL format")
	}

	userPass := parts[0]
	hostDB := parts[1]

	userPassParts := strings.SplitN(userPass, ":", 2)
	if len(userPassParts) != 2 {
		log.Fatal("❌ Invalid user:password in DB_URL")
	}
	user := userPassParts[0]
	pass := userPassParts[1]

	hostParts := strings.SplitN(hostDB, "/", 2)
	if len(hostParts) != 2 {
		log.Fatal("❌ Invalid host/dbname in DB_URL")
	}
	hostPort := hostParts[0]
	dbname := hostParts[1]

	// Final GORM DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, hostPort, dbname)

	// Connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&models.Post{}, &models.User{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	fmt.Println("✅ Database connected and migrated!")
	DB = db
	return db
}
