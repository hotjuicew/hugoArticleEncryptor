package main

import (
	"fmt"
	"github.com/hotjuicew/hugoArticleEncryptor/html"
	"log"
	"regexp"

	"github.com/hotjuicew/hugoArticleEncryptor/crypto"
	"github.com/hotjuicew/hugoArticleEncryptor/data"
)

func main() {
	fmt.Println("begin")
	// 获取所有密码和内容
	passwords, contents, err := data.GetPasswords("./content")
	if err != nil {
		log.Fatal(err)
	}

	// 将密码加密
	encryptedPasswords := crypto.GetEncryptedPasswords(passwords)
	fmt.Printf("%T", encryptedPasswords)

	encryptedContents := make(map[string]string)
	//将encryptedPasswords 作为aes加密算法的密钥,对contents中的value进行加密,存放到encryptedContents
	for file, content := range contents {
		encrypted, err := crypto.AESEncrypt(content, encryptedPasswords[file])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("encryptedPasswords[file]1：", encryptedPasswords[file])

		encryptedContents[file] = encrypted
		fmt.Println("加密后的文本", encryptedContents[file])
		//	todo
		fmt.Println("file", file) // content\111.md

		r, _ := regexp.Compile("content\\\\|\\.md$")
		fileName := r.ReplaceAllString(file, "")

		html.AddCustomAttribute(fileName, encrypted)
	}

	////解密
	//for file, encrypted := range encryptedContents {
	//	//fmt.Printf("%s: %s\n", file, encrypted)
	//	plaintext, err := crypto.AESDecrypt(encrypted, encryptedPasswords[file])
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	//fmt.Println("encryptedPasswords[file]2", encryptedPasswords[file])
	//
	//	fmt.Printf("解密后的文本%s:\n%s\n\n", file, plaintext)
	//
	//}

}
