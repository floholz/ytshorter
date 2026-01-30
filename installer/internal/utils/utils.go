package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/floholz/ytshorter/installer/assets"
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

func GetChromePath() (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("chrome path not supported on %s", runtime.GOOS)
	}

	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, "/google-chrome"), nil
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
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(configPath, "ytshorter.app/Contents/MacOS"), nil
	case "windows":
		return filepath.Join(configPath, "ytshorter.exe"), nil
	}
	return filepath.Join(configPath, "ytshorter"), nil
}

func GetNativeHostManifestPath() (string, error) {
	switch runtime.GOOS {
	case "linux":
		chromePath, err := GetChromePath()
		if err != nil {
			return "", err
		}
		return filepath.Join(chromePath, "/NativeMessagingHosts", assets.NativeHostName+".json"), nil
	case "windows":
		configPath, err := GetConfigPath()
		if err != nil {
			return "", err
		}
		return filepath.Join(configPath, "/NativeMessagingHosts", assets.NativeHostName+".json"), nil
	}

	return "", fmt.Errorf("native host manifest not supported on %s", runtime.GOOS)
}

func GetAppSourcePath() (string, error) {
	srcPath := "./application/ytshorter_app"

	switch runtime.GOOS {
	case "windows":
		srcPath += ".exe"
	}

	// Check if the source file exists and is not empty
	info, err := os.Stat(srcPath)
	if err != nil {
		return srcPath, fmt.Errorf("source application error: %w", err)
	}
	if info.Size() == 0 {
		return srcPath, fmt.Errorf("source application file %q is empty", srcPath)
	}

	return srcPath, nil
}

func TruncatePath(path string, max int) string {
	if len(path) <= max {
		return path
	}
	splits := strings.Split(path, string(filepath.Separator))
	if len(splits) < 3 {
		return path
	}
	if splits[0] == "" {
		if len(splits) == 3 {
			return path
		}
		splits = splits[1:]
	}

	start := splits[0] + string(filepath.Separator) + "..."
	tail := string(filepath.Separator) + splits[len(splits)-1]
	idx := len(splits) - 2
	for len(start)+len(tail) < max && idx > 1 {
		tail = string(filepath.Separator) + splits[idx] + tail
		idx--
	}
	return start + tail
}
