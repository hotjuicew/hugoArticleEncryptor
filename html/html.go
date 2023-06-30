package html

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// AddCustomAttribute 将加密后的内容写入html文件
func AddCustomAttribute(folderName string, encryptedText string) {
	// 构建文件夹路径
	folderPath := filepath.Join("public/post", folderName)

	// 遍历文件夹中的HTML文件
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理HTML文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			// 打开HTML文件
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// 解析HTML文档
			doc, err := goquery.NewDocumentFromReader(file)
			if err != nil {
				return err
			}

			// 查找具有id=verification的div元素，并为其添加自定义属性
			doc.Find("div#verification").Each(func(i int, selection *goquery.Selection) {
				selection.SetAttr("ciphertext", encryptedText)
			})

			// 获取修改后的HTML内容
			updatedHTML, err := doc.Html()
			if err != nil {
				return err
			}

			// 将修改后的HTML内容写入文件
			err = os.WriteFile(path, []byte(updatedHTML), 0644)
			if err != nil {
				return err
			}

			// 打印成功消息
			fmt.Printf("Custom attribute added to div in file: %s\n", path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
