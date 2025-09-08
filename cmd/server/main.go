package main

import (
	"github.com/PetJs/blog-backend/pkg/config"
	"github.com/PetJs/blog-backend/pkg/database"
	"github.com/PetJs/blog-backend/internal/repository"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/PetJs/blog-backend/internal/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	cfg := config.LoadConfig()
	
	
	db := database.ConnectDB(cfg.DB_URL)
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	repo := repository.NewPostRepository(db)
	service := services.NewPostService(repo)

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)


	router := gin.Default()
	api.RegisterPostRoutes(router, service)
	api.RegisterUserRoutes(router, userService)

	log.Println("Server running on port " + cfg.Port)

	router.Run(":" + cfg.Port)
}