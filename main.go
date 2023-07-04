package main

import (
	"fmt"
	"github.com/hotjuicew/hugoArticleEncryptor/crypto"
	"github.com/hotjuicew/hugoArticleEncryptor/data"
	"github.com/hotjuicew/hugoArticleEncryptor/html"
	"log"
	"regexp"
)

func main() {

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
