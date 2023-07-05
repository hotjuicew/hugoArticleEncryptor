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
)

func getMetadata(filename string) (map[string]interface{}, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 将内容分割成字符串,FrontMatter在最前面
	splitContent := strings.SplitN(string(content), "---", 3)

	// 解析FrontMatter中的YAML信息
	metadata := make(map[string]interface{})
	if err = yaml.Unmarshal([]byte(splitContent[1]), &metadata); err != nil {
		return nil, err
	}

	// 校验metadata是否同时含有protected和password
	_, ok1 := metadata["protected"]
	_, ok2 := metadata["password"]
	if !ok1 || !ok2 {
		return nil, nil
	}

	return metadata, nil
}

// GetContent 获取需要加密的html代码并删除

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
	// 去除meta标签
	doc.Find("meta").Remove()
	encryptedDiv := doc.Find("#encrypted")
	verificationDiv := encryptedDiv.Find("#verification")
	encryptedDiv.Contents().Remove()
	encryptedDiv.AppendSelection(verificationDiv)

	// 获取最终的HTML内容
	result, err := doc.Html()
	if err != nil {
		return "", err
	}

	return result, nil
}

// GetPasswords 获取所有密码和html内容
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
			return err
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

	fmt.Println("passwords", passwords)
	return passwords, nil
}

// GetHTML receives the file name path and returns the content to be encrypted
func GetHTML(contentDir string) string {
	base := filepath.Base(contentDir)
	baseWithoutExt := strings.TrimSuffix(base, filepath.Ext(base))

	htmlDir1 := filepath.Join("public", "posts", baseWithoutExt, "index.html")
	htmlDir2 := filepath.Join("public", "post", baseWithoutExt, "index.html")

	// Check if the first possible htmlDir exists
	_, err := os.Stat(htmlDir1)
	if err == nil {
		// htmlDir1 exists
	} else if os.IsNotExist(err) {
		// htmlDir1 does not exist, use htmlDir2
		htmlDir1 = htmlDir2
	} else {
		log.Println(err)
		return ""
	}

	file, err := os.ReadFile(htmlDir1)
	if err != nil {
		log.Println(err)
		return ""
	}

	content, err := GetContent(string(file))
	if err != nil {
		log.Println(err)
		return ""
	}

	updatedHTML, err := UpdateHTML(string(file))
	if err != nil {
		log.Println(err)
		return ""
	}

	err = os.WriteFile(htmlDir1, []byte(updatedHTML), 0644)
	if err != nil {
		log.Println(err)
	}

	return content
}
func CopyFile(sourcePath, destinationPath string, content embed.FS) error {
	// 从嵌入的文件中读取内容
	file, err := content.Open(sourcePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建目标文件并写入内容
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
