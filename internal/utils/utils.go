package utils

import (
	"os"
	"path/filepath"
)

func GetAccountsFolderPath(basePath string) (string, error) {
	if basePath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		basePath = filepath.Join(home, ".speckle")
	}
	return filepath.Join(basePath, "accounts"), nil
}
