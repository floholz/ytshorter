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

func GetExtensionPath() (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(configPath, "browser-extension"), nil
}

func GetAppPath() (string, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return "", err
	}
	switch os.Getenv("GOOS") {
	case "darwin":
		return filepath.Join(configPath, "ytshorter.app/Contents/MacOS"), nil
	case "windows":
		return filepath.Join(configPath, "ytshorter.exe"), nil
	}
	return filepath.Join(configPath, "ytshorter"), nil
}
