package api


import (
	"net/http"
	"github.com/PetJs/blog-backend/internal/models"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"
)


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
		var post models.Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.CreatePost(post); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
	})
}