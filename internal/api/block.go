package api

import (
	"net/http"
	"strconv"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterBlockRoutes(router *gin.Engine, blockService *services.BlockService) {
	admin := router.Group("/api")
	admin.Use(middleware.AuthMiddleware())

	admin.POST("/posts/:id/blocks", func(c *gin.Context) {
		postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
			return
		}

		var input struct {
			Type             string `json:"type" binding:"required"`
			Content          string `json:"content"`
			OriginalAudioURL string `json:"original_audio_url"`
			Position         int    `json:"position"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		block, err := blockService.AddBlock(uint(postID), input.Type, input.Content, input.OriginalAudioURL, input.Position)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, block)
	})

	admin.PATCH("/blocks/:id", func(c *gin.Context) {
		var input struct {
			Content  *string `json:"content"`
			Position *int    `json:"position"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updates := map[string]interface{}{}
		if input.Content != nil {
			updates["content"] = *input.Content
		}
		if input.Position != nil {
			updates["position"] = *input.Position
		}

		block, err := blockService.UpdateBlock(c.Param("id"), updates)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
			return
		}
		c.JSON(http.StatusOK, block)
	})

	admin.DELETE("/blocks/:id", func(c *gin.Context) {
		if err := blockService.DeleteBlock(c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Block not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Block deleted successfully"})
	})
}
