package utils

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func newCloudinaryClient() (*cloudinary.Cloudinary, error) {
	return cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
}

// UploadFile uploads any io.Reader to Cloudinary and returns the secure URL.
// resourceType should be "image" for images/gifs and "video" for audio files.
func UploadFile(file io.Reader, resourceType string) (string, error) {
	cld, err := newCloudinaryClient()
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:       "blog",
		ResourceType: resourceType,
	})
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
