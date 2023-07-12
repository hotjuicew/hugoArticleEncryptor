package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func GetThemeFromConfig() (string, error) {
	configFile := findConfigFile()
	if configFile == "" {
		return "", fmt.Errorf("configuration file not found")
	}

	fileData, err := os.ReadFile(configFile)
	if err != nil {
		return "", fmt.Errorf("failed to read configuration file：%v", err)
	}

	ext := filepath.Ext(configFile)
	switch ext {
	case ".toml":
		theme, err := getThemeFromToml(fileData)
		if err != nil {
			return "", fmt.Errorf("failed to parse TOML configuration file：%v", err)
		}
		return theme, nil
	case ".yml", ".yaml":
		theme, err := getThemeFromYaml(fileData)
		if err != nil {
			return "", fmt.Errorf("failed to parse YAML configuration file：%v", err)
		}
		return theme, nil
	default:
		return "", fmt.Errorf("unsupported configuration file format：%s", ext)
	}
}

func findConfigFile() string {
	_, err := os.Stat("config.toml")
	if err == nil {
		return "config.toml"
	}

	_, err = os.Stat("config.yml")
	if err == nil {
		return "config.yml"
	}
	_, err = os.Stat("config.yaml")
	if err == nil {
		return "config.yaml"
	}
	return ""
}

func getThemeFromToml(data []byte) (string, error) {
	var config struct {
		Theme string `toml:"theme"`
	}
	err := toml.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}
	return config.Theme, nil
}

func getThemeFromYaml(data []byte) (string, error) {
	var config struct {
		Theme string `yaml:"theme"`
	}
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}
	return config.Theme, nil
}
func GetThemesFolderName() (string, error) {
	files, err := os.ReadDir("themes")
	if err != nil {
		return "", fmt.Errorf("failed to read themes directory: %v", err)
	}
	if len(files) != 1 || !files[0].IsDir() {
		return "", fmt.Errorf("themes directory should contain exactly one folder")
	}
	return files[0].Name(), nil
}
