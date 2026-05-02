package api

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/PetJs/blog-backend/internal/middleware"
	"github.com/PetJs/blog-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var extMimeTypes = map[string]string{
	".mp3":  "audio/mpeg",
	".m4a":  "audio/mp4",
	".mp4":  "audio/mp4",
	".webm": "audio/webm",
	".ogg":  "audio/ogg",
	".wav":  "audio/wav",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
}

// resolveMimeType returns the declared MIME type, falling back to extension-based detection.
func resolveMimeType(filename, declared string) string {
	if declared != "" && declared != "application/octet-stream" {
		return declared
	}
	ext := strings.ToLower(filepath.Ext(filename))
	if m, ok := extMimeTypes[ext]; ok {
		return m
	}
	return declared
}

var allowedMimeTypes = map[string]string{
	"image/jpeg":  "image",
	"image/png":   "image",
	"image/gif":   "image",
	"audio/mpeg":  "video",
	"audio/mp3":   "video",
	"audio/mp4":   "video", // .m4a files
	"audio/x-m4a": "video",
	"audio/webm":  "video",
	"audio/ogg":   "video",
	"audio/wav":   "video",
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

		mimeType := resolveMimeType(header.Filename, header.Header.Get("Content-Type"))
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

		mimeType := resolveMimeType(header.Filename, header.Header.Get("Content-Type"))
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
