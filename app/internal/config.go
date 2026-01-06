package internal

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Keybind []string `yaml:"keybind,flow"`
}

func NewDefaultConfig() Config {
	return Config{Keybind: []string{"ctrl", "shift", "u"}}
}

func getConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(dir, "YTShorter")
	// Ensure the directory exists
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "config.yaml"), nil
}

func LoadConfig() Config {
	conf := NewDefaultConfig()
	path, _ := getConfigPath()

	data, err := os.ReadFile(path)
	if err != nil {
		// If file doesn't exist, we just return defaults
		return conf
	}

	// Unmarshal will only overwrite fields present in the YAML file
	yaml.Unmarshal(data, &conf)
	return conf
}

func (c *Config) Save() error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
