package data

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getMetadata(filename string) (map[string]interface{}, string, error) {
	content, err := ioutil.ReadFile(filename)
	fmt.Println("content:")
	if err != nil {
		return nil, "", err
	}

	// 将内容分割成字符串,FrontMatter在最前面
	splitContent := strings.SplitN(string(content), "---", 3)

	// 解析FrontMatter中的YAML信息
	var metadata map[string]interface{}
	if err = yaml.Unmarshal([]byte(splitContent[1]), &metadata); err != nil {
		return nil, "", err
	}
	return metadata, string(content), nil
}

func GetPasswords(contentDir string) (map[string]string, map[string]string, error) {
	passwords := make(map[string]string)
	contents := make(map[string]string)

	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		metadata, content, err := getMetadata(path)
		if err != nil {
			return err
		}

		password, ok := metadata["password"].(string)
		if ok {
			passwords[path] = password
			contents[path] = content
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return passwords, contents, nil
}
