package html

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// WriteEncryptedContentToHTML 将加密后的内容写入html文件
func WriteEncryptedContentToHTML(folderName string, encryptedText string) {
	// 构建文件夹路径
	folderPath := filepath.Join("public", folderName)

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

			// 查找具有id=secret的div元素
			doc.Find("#secret").Each(func(i int, s *goquery.Selection) {
				s.SetText(encryptedText)
			})
			// 向HTML文件的<body>标签添加一个引用外部JavaScript文件的<script>标签
			doc.Find("body").AppendHtml(`<script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.9-1/crypto-js.js"></script>`)

			// 添加一个引用'/static/js/decrypt.js'的<script>标签
			doc.Find("body").AppendHtml(`<script src="../../js/AESDecrypt.js"></script>`)

			// 获取修改后的HTML内容
			updatedHTML, err := doc.Html()
			if err != nil {
				return err
			}

			// 向 <body> 标签添加引用外部 JavaScript 文件的 <script> 标签
			scriptTag1 := fmt.Sprintf(`<script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.9-1/crypto-js.js"></script>`)
			doc.Find("body").AppendHtml(scriptTag1)

			// 向 <body> 标签添加引用 /static/js/AESDecrypt.js 的 <script> 标签
			scriptTag2 := fmt.Sprintf(`<script src="/static/js/AESDecrypt.js"></script>`)
			doc.Find("body").AppendHtml(scriptTag2)

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
