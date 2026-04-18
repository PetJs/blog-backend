package api

import (
	"net/http"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.Engine, postService *services.PostService) {
	api := router.Group("/api")

	// Public routes
	api.GET("/posts", func(c *gin.Context) {
		posts, err := postService.GetPublishedPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	api.GET("/posts/:slug", func(c *gin.Context) {
		post, err := postService.GetPostBySlug(c.Param("slug"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	// Admin-only routes
	admin := api.Group("/")
	admin.Use(middleware.AuthMiddleware())

	admin.POST("/posts", func(c *gin.Context) {
		post, err := postService.CreatePost()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, post)
	})

	admin.PATCH("/posts/:id", func(c *gin.Context) {
		var input struct {
			Title      string `json:"title"`
			Excerpt    string `json:"excerpt"`
			CoverImage string `json:"cover_image"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post, err := postService.UpdatePost(c.Param("id"), input.Title, input.Excerpt, input.CoverImage)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	admin.PATCH("/posts/:id/publish", func(c *gin.Context) {
		post, err := postService.PublishPost(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	admin.DELETE("/posts/:id", func(c *gin.Context) {
		if err := postService.DeletePost(c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	})
}
