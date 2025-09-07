package api


import (
	"net/http"
	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"
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

	router.POST("/posts", func(c *gin.Context){
		var input CreatePostInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post := models.Post{
			Title:   input.Title,
			Content: input.Content,	
			Author: input.Author,
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
}