package api

import (
	"net/http"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterAnalyticsRoutes(router *gin.Engine, analyticsService *services.AnalyticsService) {
	// Public: record a view when a post is opened
	// Uses :id param name to match the existing POST tree wildcard from block routes
	router.POST("/api/posts/:id/view", func(c *gin.Context) {
		slug := c.Param("id")
		if err := analyticsService.RecordView(slug); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "View recorded"})
	})

	// Protected: admin dashboard stats
	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())

	admin.GET("/analytics", func(c *gin.Context) {
		stats, err := analyticsService.GetDashboardStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stats)
	})
}
