//go:build windows

package utils

import (
	"golang.org/x/sys/windows/registry"
)

func InstallRegistryEntry() error {

	keyPath := `Software\Google\Chrome\NativeMessagingHosts\com.floholz.ytshorter`
	manifestPath, err := GetNativeHostManifestPath()
	if err != nil {
		return err
	}

	// Open HKCU
	k, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		keyPath,
		registry.SET_VALUE,
	)
	if err != nil {
		return err
	}
	defer k.Close()

	// Set default value (@)
	err = k.SetStringValue("", manifestPath)
	if err != nil {
		return err
	}
	return nil
}
