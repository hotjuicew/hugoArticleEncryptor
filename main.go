package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/hotjuicew/hugoArticleEncryptor/crypto"
	"github.com/hotjuicew/hugoArticleEncryptor/data"
	"github.com/hotjuicew/hugoArticleEncryptor/html"
)

// go:embed AESDecrypt.js
var aesDecryptScript embed.FS

// go:embed secret.html
var secretHtml embed.FS

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide the theme name as a parameter.〒▽〒")
		return
	}

	themeName := os.Args[1]
	fmt.Println("Theme Name: ", themeName)

	err := data.CopyFile("AESDecrypt.js", fmt.Sprintf("/themes/%s/static/js/AESDecrypt.js", themeName), aesDecryptScript)
	if err != nil {
		log.Fatal(err)
	}
	err = data.CopyFile("secret.html", fmt.Sprintf("/themes/%s/layouts/partials/secret.html", themeName), secretHtml)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("11111111111111111111111")
	// 创建一个命令对象
	cmd := exec.Command("hugo")

	// 执行命令并等待其完成
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// 输出命令执行结果
	fmt.Println(string(output))
	fmt.Println("11111111111111111111111")
	// 获取所有密码和内容
	passwords, err := data.GetPasswords("./content")
	if err != nil {
		log.Fatal(err)
	}

	// 将密码加密
	encryptedPasswords := crypto.GetEncryptedPasswords(passwords)

	encryptedContents := make(map[string]string)
	//将encryptedPasswords 作为aes加密算法的密钥,对contents中的value进行加密,存放到encryptedContents
	for file := range passwords {
		content := data.GetHTML(file)
		encrypted, err := crypto.AESEncrypt(content, encryptedPasswords[file])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(file, "中要加密的部分：", content)

		encryptedContents[file] = encrypted
		r, _ := regexp.Compile("content\\\\|\\.md$")
		fileName := r.ReplaceAllString(file, "")

		html.WriteEncryptedContentToHTML(fileName, encrypted)
	}
}
