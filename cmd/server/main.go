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

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, relying on environment variables")
	}

	cfg := config.LoadConfig()

	db := database.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	repo := repository.NewPostRepository(db)
	service := services.NewPostService(repo)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api.RegisterPostRoutes(router, service)
	api.RegisterUserRoutes(router, userService)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port // fallback for local dev
	}
	log.Println("Server running on port " + port)

	router.Run(":" + port)
}