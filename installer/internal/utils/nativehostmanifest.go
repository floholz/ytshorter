package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/floholz/ytshorter/installer/assets"
)

type NativeHostManifest struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Path           string   `json:"path"`
	Type           string   `json:"type"`
	AllowedOrigins []string `json:"allowed_origins"`
}

func InstallManifest() error {
	appPath, err := GetAppPath()
	if err != nil {
		return err
	}

	manifest := NativeHostManifest{
		Name:        assets.NativeHostName,
		Description: "YTShorter native host application",
		Path:        appPath,
		Type:        "stdio",
		AllowedOrigins: []string{
			fmt.Sprintf("chrome-extension://%s/", assets.ExtensionId),
		},
	}

	manifestBytes, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}

	manifestPath, err := GetNativeHostManifestPath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(manifestPath), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(manifestPath, manifestBytes, 0644)
	if err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		err = InstallRegistryEntry()
		if err != nil {
			return err
		}
	}

	return nil
}
