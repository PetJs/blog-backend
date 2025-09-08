package api

import (
	"fmt"
	"net/http"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type CreatePostInput struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
    Author  string `json:"author" binding:"required"`
}



func RegisterPostRoutes(router *gin.Engine, service *services.PostService){
	router.GET("/posts", func(c *gin.Context){
		posts, err := service.GetPosts()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusOK, posts)
	})

	// Protected routes
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/posts", func(c *gin.Context){
		var input CreatePostInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Get user_id from context set by AuthMiddleware
		user_id, exists := c.Get("user_id")
		if !exists {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
			return
		}

		post := models.Post{
			Title:   input.Title,
			Content: input.Content,	
			Author: input.Author,
			UserID:  user_id.(uint), 
		}


		createdPost, err := service.CreatePost(post)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, gin.H{
			"status": http.StatusCreated,
			"message": "Post created successfully",
			"data": createdPost,
		})
	})

	auth.DELETE("/posts/:id", func(c *gin.Context) {
		user_id, _ := c.Get("user_id")
		post_id := c.Param("id")

		if err := service.DeletePost(post_id, user_id.(uint)); err != nil {
			if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own posts"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
	})

	// Get posts by the authenticated user
	auth.GET("/users/posts", func(ctx *gin.Context) {
		user_id, _ := ctx.Get("user_id")
		fmt.Printf("This User id is %v\n", user_id)

		posts, err := service.GetUserPosts(user_id.(uint))
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, posts)
	})
}