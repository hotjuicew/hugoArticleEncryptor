package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hotjuicew/hugoArticleEncryptor/config"
	"github.com/hotjuicew/hugoArticleEncryptor/crypto"
	"github.com/hotjuicew/hugoArticleEncryptor/data"
	"github.com/hotjuicew/hugoArticleEncryptor/html"
)

//go:embed AESDecrypt.js
var aesDecryptScript embed.FS

//go:embed secret.html
var secretHtml embed.FS

func main() {
	//get themeName
	themeName, err := config.GetThemeFromConfig()
	if err != nil {
		fmt.Println("GetThemeFromConfig gets err", err)
		return
	}
	if themeName == "" {
		themeName, _ = config.GetThemesFolderName()
	}
	//在single.html中插入代码
	config.ChangeSingleHTML(themeName)
	if _, err := os.Stat(filepath.Join("themes", themeName, "static", "js")); os.IsNotExist(err) {
		// 路径不存在，创建路径
		err := os.MkdirAll(filepath.Join("themes", themeName, "static", "js"), os.ModePerm)
		if err != nil {
			log.Fatalf("Unable to create path：%v\n\n", err)
			return
		}
	}
	err = data.CopyFile("AESDecrypt.js", filepath.Join("themes/", themeName, "/static/js/AESDecrypt.js"), aesDecryptScript)
	if err != nil {
		log.Fatalf("data.CopyFile: AESDecrypt.js gets error %v", err)
	}
	err = data.CopyFile("secret.html", filepath.Join("themes/", themeName, "/layouts/partials/secret.html"), secretHtml)
	if err != nil {
		log.Fatal("data.CopyFile: secret.html gets error", err)
	}

	output, err := exec.Command("hugo").Output()
	if err != nil {
		log.Fatalln("cmd.Output() gets error", err)
	}

	// Output command execution results
	fmt.Println(string(output))

	// Get all passwords and content
	passwords, err := data.GetPasswords("./content")
	if err != nil {
		log.Fatalln("data.GetPasswords gets error", err)
	}

	// Encrypt the password
	encryptedPasswords := crypto.GetEncryptedPasswords(passwords)

	encryptedContents := make(map[string]string)

	for file := range passwords {
		content := data.GetHTML(file)
		encrypted, err := crypto.AESEncrypt(content, encryptedPasswords[file])
		if err != nil {
			log.Fatal("crypto.AESEncrypt gets error", err)
		}
		encryptedContents[file] = encrypted
		filename := filepath.Base(file)
		extension := filepath.Ext(filename)
		name := filename[:len(filename)-len(extension)]
		html.WriteEncryptedContentToHTML(filepath.Base(name), encrypted)
	}
}
