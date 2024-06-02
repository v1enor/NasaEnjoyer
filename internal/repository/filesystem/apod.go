package filesystem

import (
	"NasaEnjoyer/domain"
	"NasaEnjoyer/internal/repository"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileRepository interface {
	SaveImage(url, saveDir string) (string, error)
}

type fileRepository struct{}

func NewFileRepository() FileRepository {
	return &fileRepository{}
}

func (r *fileRepository) SaveImage(url, saveDir string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if !repository.IsImage(contentType) {
		return "", domain.ErrNotAnImage
	}

	filename := filepath.Base(url)

	filepath := filepath.Join(saveDir, filename)

	err = repository.EnsureDirExists(saveDir)
	if err != nil {
		return "", domain.ErrDirectoryCreation

	}

	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filepath, nil
}
