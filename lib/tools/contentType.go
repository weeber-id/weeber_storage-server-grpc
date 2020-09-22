package tools

import (
	"path"
)

// GetContentTypeFromExt for save to Minio server
// Example .jpg to image/jpg
func GetContentTypeFromExt(ext string) string {
	// get only ext from .jpg to jpg
	ext = ext[1:]

	switch ext {
	case "jpeg", "jpg", "png":
		return path.Join("image", ext)
	}

	return ""
}
