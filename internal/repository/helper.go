package repository

import (
	"os"
	"strings"
)

func EnsureDirExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

func IsImage(contentType string) bool {
	return strings.HasPrefix(contentType, "image/")
}
