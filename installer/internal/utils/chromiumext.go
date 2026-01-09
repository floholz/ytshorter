package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Preferences struct {
	Extensions struct {
		Settings map[string]Extension `json:"settings"`
	} `json:"extensions"`
}

type Extension struct {
	Location       int    `json:"location"`
	Path           string `json:"path"`
	DisableReasons []int  `json:"disable_reasons,omitempty"`
}

func DetectExtension(extensionId string, profiles ...string) (*Extension, error) {
	if len(profiles) == 0 {
		var err error
		profiles, err = findProfiles()
		if err != nil {
			return nil, err
		}
	}

	for _, profile := range profiles {
		ext, err := detectForProfile(extensionId, profile)
		if err != nil {
			continue
		}
		if ext != nil {
			return ext, nil
		}
	}

	return nil, nil
}

func findProfiles() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	chromeDir := path.Join(homeDir, ".config/google-chrome")
	entries, err := os.ReadDir(chromeDir)
	if err != nil {
		return nil, err
	}

	var profiles []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()
		if name == "Default" || strings.HasPrefix(name, "Profile") {
			profiles = append(profiles, name)
		}
	}

	return profiles, nil
}

func detectForProfile(extensionId, profile string) (*Extension, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path.Join(homeDir, ".config/google-chrome", profile, "Preferences"))
	if err != nil {
		return nil, err
	}

	var prefs Preferences
	if err = json.Unmarshal(data, &prefs); err != nil {
		return nil, err
	}

	ext, ok := prefs.Extensions.Settings[extensionId]
	if !ok {
		fmt.Println("Extension not registered")
		return nil, nil
	}

	return &ext, nil
}

func CopyExtensionToConfigFolder() error {
	srcPath := "./browser-extension"

	// Check if the source directory exists and has files
	entries, err := os.ReadDir(srcPath)
	if err != nil {
		return fmt.Errorf("source directory error: %w", err)
	}
	if len(entries) == 0 {
		return fmt.Errorf("source directory %q is empty", srcPath)
	}

	dest, err := GetExtensionPath()
	if err != nil {
		return err
	}

	err = os.RemoveAll(dest)
	if err != nil {
		return err
	}

	return os.CopyFS(dest, os.DirFS(srcPath))
}
