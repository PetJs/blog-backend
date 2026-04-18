package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PetJs/blog-backend/internal/services"
	"github.com/PetJs/blog-backend/pkg/utils"
)

func RegisterAuthRoutes(router *gin.Engine, adminService *services.AdminService) {
	router.POST("/admin/login", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		admin, err := adminService.LoginAdmin(req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		token, err := utils.GenerateToken(admin.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
