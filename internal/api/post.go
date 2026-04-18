package api

import (
	"net/http"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type PostInput struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ImageURL string `json:"image_url"`
}

func RegisterPostRoutes(router *gin.Engine, service *services.PostService) {
	// Public
	router.GET("/posts", func(c *gin.Context) {
		posts, err := service.GetPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	router.GET("/posts/:id", func(c *gin.Context) {
		post, err := service.GetPost(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	// Admin-only
	admin := router.Group("/")
	admin.Use(middleware.AuthMiddleware())

	admin.POST("/posts", func(c *gin.Context) {
		var input PostInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post := models.Post{
			Title:    input.Title,
			Content:  input.Content,
			ImageURL: input.ImageURL,
		}

		created, err := service.CreatePost(post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Post created successfully",
			"data":    created,
		})
	})

	admin.PUT("/posts/:id", func(c *gin.Context) {
		var input PostInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updates := map[string]interface{}{
			"title":     input.Title,
			"content":   input.Content,
			"image_url": input.ImageURL,
		}

		updated, err := service.UpdatePost(c.Param("id"), updates)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Post updated successfully",
			"data":    updated,
		})
	})

	admin.DELETE("/posts/:id", func(c *gin.Context) {
		if err := service.DeletePost(c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	})
}
