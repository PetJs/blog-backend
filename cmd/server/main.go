package main

import (
	"log"
	"os"
	"time"

	"github.com/PetJs/blog-backend/internal/api"
	"github.com/PetJs/blog-backend/internal/repository"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/PetJs/blog-backend/pkg/config"
	"github.com/PetJs/blog-backend/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, relying on environment variables")
	}

	cfg := config.LoadConfig()

	db := database.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Admin
	adminRepo := repository.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)

	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminEmail == "" || adminPassword == "" {
		log.Fatal("❌ ADMIN_EMAIL and ADMIN_PASSWORD must be set")
	}
	if err := adminService.SeedAdmin(adminEmail, adminPassword); err != nil {
		log.Fatal("❌ Failed to seed admin:", err)
	}

	// Posts & Blocks
	postRepo := repository.NewPostRepository(db)
	postService := services.NewPostService(postRepo)

	blockRepo := repository.NewBlockRepository(db)
	blockService := services.NewBlockService(blockRepo)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api.RegisterAuthRoutes(router, adminService)
	api.RegisterPostRoutes(router, postService)
	api.RegisterBlockRoutes(router, blockService)
	api.RegisterUploadRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}
	log.Println("Server running on port " + port)
	router.Run(":" + port)
}
