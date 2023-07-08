package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
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
func ChangeSingleHTML(themeName string) {
	var filename string
	if fileExists(filepath.Join("themes/", themeName, "/layouts/post/single.html")) {
		filename = filepath.Join("themes/", themeName, "/layouts/post/single.html")
	} else if fileExists(filepath.Join("themes/", themeName, "/layouts/_default/single.html")) {
		filename = filepath.Join("themes/", themeName, "/layouts/_default/single.html")
	} else {
		log.Fatal("can't find single.html")
	}
	// Read HTML files
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to read single.html")
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Fail to read single html")
	}
	contentStr := string(content)
	// Check if "secret.html" already exists in contentStr
	if strings.Contains(contentStr, "secret.html") {
		return
	}

	// Find the first "<div"
	divIndex := strings.Index(contentStr, "<div")
	if divIndex == -1 {
		log.Fatal("can't find '<div'")
	}

	// After the first "<div" insert "id="encrypted""
	updatedHTML := contentStr[:divIndex+5] + ` id="encrypted"` + contentStr[divIndex+5:]

	// Find the first ">"
	endDivIndex := strings.Index(updatedHTML, ">")

	// After the first "</div>" insert "\n{{ partial "secret.html" . }}\n"
	updatedHTML = updatedHTML[:endDivIndex+6] + "\n{{ partial \"secret.html\" . }}\n" + updatedHTML[endDivIndex+6:]
	// Write the modified content to the original file
	err = os.WriteFile(filename, []byte(updatedHTML), 0644)
	if err != nil {
		log.Fatal("Failed to write modified single HTML")
	}
}
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
