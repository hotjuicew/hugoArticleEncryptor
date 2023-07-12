package data

import (
	"embed"
	"github.com/hotjuicew/hugoArticleEncryptor/crypto"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CopyFile 将嵌入文件复制到目标路径
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

// WalkHTMLFiles 遍历一遍public/posts(或post)下面的文件夹里面的html文件
func WalkHTMLFiles() error {
	err := filepath.WalkDir("public", func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 如果当前路径是文件夹，则继续遍历子文件夹
		if info.IsDir() {
			return nil
		}

		// 检查文件扩展名是否为.html
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".html" {
			// 处理HTML文件的逻辑
			getData(path)
		}

		return nil
	})

	return err
}

// getData 获取HTML文件中的密码和内容，修改HTML文件并写入
func getData(path string) {
	// 从嵌入文件中读取HTML内容
	file, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error reading file:", path)
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(file)))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return
	}

	doc.Find("body").AppendHtml(`<script src="https://cdn.jsdelivr.net/npm/marked/marked.min.js"></script>`)
	doc.Find("body").AppendHtml(`<script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.9-1/crypto-js.js"></script>`)
	doc.Find("body").AppendHtml(`<script src="/js/AESDecrypt.js"></script>`)
	secretElements := doc.Find("div#secret")
	if secretElements.Length() == 0 {
		return
	}
	passwordAttr, _ := secretElements.Attr("password")
	innerText := secretElements.Text()

	//在html变量中删除password属性，并且删除innerText
	secretElements.RemoveAttr("password")
	encryptedPassword := crypto.GetEncryptedPassword(passwordAttr)
	encryptedContent, err := crypto.AESEncrypt(innerText, encryptedPassword)
	if err != nil {
		log.Fatal("crypto.AESEncrypt(innerText, encryptedPassword) gets err", err)
	}
	secretElements.SetText(encryptedContent)
	//把修改后的HTML内容写入源文件
	newHtml, err := doc.Html()
	if err != nil {
		log.Fatal("doc.Html() gets err: ", err)
	}
	err = os.WriteFile(path, []byte(newHtml), 0644)
	if err != nil {
		log.Fatal("os.WriteFile(path, []byte(newHtml), 0644) gets err: ", err)
	}
}
