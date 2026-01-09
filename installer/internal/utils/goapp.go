package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CopyAppToConfigFolder() error {
	srcPath, err := GetAppSourcePath()
	if err != nil {
		return err
	}

	dest, err := GetAppPath()
	if err != nil {
		return err
	}

	// Ensure the destination directory exists (especially for macOS paths)
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	// Remove the existing file if it exists
	_ = os.Remove(dest)

	// Read and Write the file
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	// We use 0755 to ensure the copied file is executable
	return os.WriteFile(dest, data, 0755)
}

func VerifyInstallation() error {
	appPath, err := GetAppPath()
	if err != nil {
		return err
	}

	// Check if the source file exists and is not empty
	info, err := os.Stat(appPath)
	if err != nil {
		return fmt.Errorf("source application error: %w", err)
	}
	if info.Size() == 0 {
		return fmt.Errorf("source application file %q is empty", appPath)
	}
	return nil
}
