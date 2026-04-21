package api

import (
	"net/http"
	"strings"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var allowedMimeTypes = map[string]string{
	"image/jpeg": "image",
	"image/png":  "image",
	"image/gif":  "image",
	"audio/mpeg": "video",
	"audio/mp3":  "video",
	"audio/webm": "video",
	"audio/ogg":  "video",
}

func RegisterUploadRoutes(router *gin.Engine) {
	admin := router.Group("/api")
	admin.Use(middleware.AuthMiddleware())

	admin.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
			return
		}
		defer file.Close()

		mimeType := header.Header.Get("Content-Type")
		resourceType, allowed := allowedMimeTypes[mimeType]
		if !allowed {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File type not allowed"})
			return
		}

		url, err := utils.UploadFile(file, resourceType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"url":  url,
			"type": mimeType,
		})
	})

	admin.POST("/transcribe", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No audio file provided"})
			return
		}
		defer file.Close()

		mimeType := header.Header.Get("Content-Type")
		if !strings.HasPrefix(mimeType, "audio/") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File must be an audio file"})
			return
		}

		transcript, err := utils.TranscribeAudio(file, mimeType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transcription failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"transcript": transcript})
	})
}
