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
	rawDSN := os.Getenv("DB_URL")
	if rawDSN == "" {
		log.Fatal("❌ DB_URL environment variable not set")
	}

	// Convert mysql://user:pass@host:port/dbname -> GORM DSN
	trimmed := strings.TrimPrefix(rawDSN, "mysql://")
	parts := strings.SplitN(trimmed, "@", 2)
	if len(parts) != 2 {
		log.Fatal("❌ Invalid DB_URL format")
	}

	userPassParts := strings.SplitN(parts[0], ":", 2)
	if len(userPassParts) != 2 {
		log.Fatal("❌ Invalid user:password in DB_URL")
	}

	hostParts := strings.SplitN(parts[1], "/", 2)
	if len(hostParts) != 2 {
		log.Fatal("❌ Invalid host/dbname in DB_URL")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		userPassParts[0], userPassParts[1], hostParts[0], hostParts[1],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Post must come before Block so the foreign key resolves correctly
	if err := db.AutoMigrate(&models.Admin{}, &models.Post{}, &models.Block{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	fmt.Println("✅ Database connected and migrated!")
	DB = db
	return db
}
