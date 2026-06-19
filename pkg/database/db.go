package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/PetJs/blog-backend/internal/models"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	rawDSN := os.Getenv("DB_URL")
	if rawDSN == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	// Convert mysql://user:pass@host:port/dbname -> GORM DSN
	trimmed := strings.TrimPrefix(rawDSN, "mysql://")
	parts := strings.SplitN(trimmed, "@", 2)
	if len(parts) != 2 {
		log.Fatal("Invalid DB_URL format")
	}

	userPassParts := strings.SplitN(parts[0], ":", 2)
	if len(userPassParts) != 2 {
		log.Fatal("Invalid user:password in DB_URL")
	}

	hostParts := strings.SplitN(parts[1], "/", 2)
	if len(hostParts) != 2 {
		log.Fatal("Invalid host/dbname in DB_URL")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		userPassParts[0], userPassParts[1], hostParts[0], hostParts[1],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error), // only log actual errors
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Post must come before Block/PostView so foreign keys resolve correctly
	if err := db.AutoMigrate(&models.Admin{}, &models.Post{}, &models.Block{}, &models.PostView{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Drop stale columns that were removed from models but survived AutoMigrate
	dropStaleColumn(db, "posts", "content")

	fmt.Println("Database connected and migrated!")
	DB = db
	return db
}

// dropStaleColumn removes a column that no longer exists in the model but is
// still present in the DB (AutoMigrate never drops columns automatically).
func dropStaleColumn(db *gorm.DB, table, column string) {
	var count int64
	db.Raw(
		"SELECT COUNT(*) FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ? AND COLUMN_NAME = ?",
		table, column,
	).Scan(&count)
	if count > 0 {
		if err := db.Exec("ALTER TABLE `" + table + "` DROP COLUMN `" + column + "`").Error; err != nil {
			log.Printf("Could not drop stale column %s.%s: %v", table, column, err)
		} else {
			log.Printf("Dropped stale column %s.%s", table, column)
		}
	}
}
