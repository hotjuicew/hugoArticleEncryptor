package data

import (
	"embed"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func getMetadata(filename string) (map[string]interface{}, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	splitContent := strings.SplitN(string(content), "---", 3)
	metadata := make(map[string]interface{})
	if err = yaml.Unmarshal([]byte(splitContent[1]), &metadata); err != nil {
		return nil, err
	}
	_, ok1 := metadata["protected"]
	_, ok2 := metadata["password"]
	if !ok1 || !ok2 {
		return nil, nil
	}
	return metadata, nil
}

// GetContent Get the html code that needs to be encrypted and delete it
func GetContent(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}
	encryptedDiv := doc.Find("#encrypted")
	verificationDiv := encryptedDiv.Find("#verification")
	verificationDiv.Remove()
	others := encryptedDiv.Children()
	othersHTML, _ := goquery.OuterHtml(others)
	return strings.TrimSpace(othersHTML), nil
}
func UpdateHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}
	// Remove meta tags
	doc.Find("meta").Remove()
	encryptedDiv := doc.Find("#encrypted")
	verificationDiv := encryptedDiv.Find("#verification")
	encryptedDiv.Contents().Remove()
	encryptedDiv.AppendSelection(verificationDiv)
	// Get the final HTML content
	result, err := doc.Html()
	if err != nil {
		return "", err
	}
	return result, nil
}

// GetPasswords Get all passwords and html content
func GetPasswords(contentDir string) (map[string]string, error) {
	passwords := make(map[string]string)

	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		metadata, err := getMetadata(path)
		fmt.Println(path)
		if err != nil {
			log.Println("getMetadata(path) gets err", err)
		}

		password, ok := metadata["password"].(string)
		if ok {
			passwords[path] = password
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return passwords, nil
}

// GetHTML receives the file name path and returns the content to be encrypted
func GetHTML(contentDir string) string {
	base := filepath.Base(contentDir)
	baseWithoutExt := strings.TrimSuffix(base, filepath.Ext(base))
	baseWithoutExtLower := strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return r
	}, baseWithoutExt)
	htmlDir1 := filepath.Join("public", "posts", baseWithoutExtLower, "index.html")
	htmlDir2 := filepath.Join("public", "post", baseWithoutExtLower, "index.html")

	// Check if the first possible htmlDir exists
	_, err := os.Stat(htmlDir1)
	if err == nil {
		// htmlDir1 exists
	} else if os.IsNotExist(err) {
		// htmlDir1 does not exist, use htmlDir2
		htmlDir1 = htmlDir2
	} else {
		log.Println("htmlDir gets error", err)
		return ""
	}

	file, err := os.ReadFile(htmlDir1)
	if err != nil {
		log.Println("os.ReadFile(htmlDir1) gets error", err)
		return ""
	}

	content, err := GetContent(string(file))
	if err != nil {
		log.Println("GetContent(string(file)) gets error", err)
		return ""
	}

	updatedHTML, err := UpdateHTML(string(file))
	if err != nil {
		log.Println("UpdateHTML(string(file)) gets error", err)
		return ""
	}

	err = os.WriteFile(htmlDir1, []byte(updatedHTML), 0644)
	if err != nil {
		log.Println("os.WriteFile(htmlDir1, []byte(updatedHTML), 0644) gets error", err)
	}

	return content
}
func CopyFile(sourcePath, destinationPath string, content embed.FS) error {
	// Read content from embedded files
	file, err := content.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create the target file and write the contents
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		return err
	}

	return nil
}
