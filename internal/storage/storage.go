package storage

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	// Directory permissions (rwxr-xr-x)
	dirPermissions = 0755
	// File permissions (rw-r--r--)
	filePermissions = 0644
)

type ImageStorage interface {
	Save(data []byte, mimeType string) (string, error)
	Load(filePath string) (data []byte, mimeType string, err error)
}

type FileStorage struct {
	baseDir string
}

func NewFileStorage(baseDir string) (*FileStorage, error) {
	absDir, err := filepath.Abs(baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	if err := os.MkdirAll(absDir, dirPermissions); err != nil {
		return nil, fmt.Errorf("failed to create storage directory: %w", err)
	}

	info, err := os.Stat(absDir)
	if err != nil {
		return nil, fmt.Errorf("failed to stat storage directory: %w", err)
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("storage path is not a directory: %s", absDir)
	}

	return &FileStorage{
		baseDir: absDir,
	}, nil
}

func (fs *FileStorage) Save(data []byte, mimeType string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("image data is empty")
	}

	ext := getExtensionFromMIME(mimeType)

	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(fs.baseDir, filename)

	if err := os.WriteFile(filePath, data, filePermissions); err != nil {
		return "", fmt.Errorf("failed to write image file: %w", err)
	}

	return filePath, nil
}

func (fs *FileStorage) Load(filePath string) ([]byte, string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read image file: %w", err)
	}

	if len(data) == 0 {
		return nil, "", fmt.Errorf("image file is empty: %s", filePath)
	}

	mimeType := getMIMETypeFromExtension(filePath)

	return data, mimeType, nil
}

func getMIMETypeFromExtension(filePath string) string {
	ext := filepath.Ext(filePath)

	mimeType := mime.TypeByExtension(ext)
	if mimeType != "" {
		return mimeType
	}

	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".webp":
		return "image/webp"
	case ".gif":
		return "image/gif"
	case ".heic":
		return "image/heic"
	case ".heif":
		return "image/heif"
	default:
		return "image/jpeg"
	}
}

func getExtensionFromMIME(mimeType string) string {
	exts, err := mime.ExtensionsByType(mimeType)
	if err == nil && len(exts) > 0 {
		return exts[0]
	}

	switch mimeType {
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	case "image/gif":
		return ".gif"
	default:
		return ".png"
	}
}
