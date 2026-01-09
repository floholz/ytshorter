package utils

import (
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(dir, "YTShorter")
	// Ensure the directory exists
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}
	return appDir, nil
}

func MustGetConfigPath() string {
	path, err := GetConfigPath()
	if err != nil {
		return ""
	}
	return path
}
