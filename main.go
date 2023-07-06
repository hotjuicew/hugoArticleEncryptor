package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/hotjuicew/hugoArticleEncryptor/crypto"
	"github.com/hotjuicew/hugoArticleEncryptor/data"
	"github.com/hotjuicew/hugoArticleEncryptor/html"
)

//go:embed AESDecrypt.js
var aesDecryptScript embed.FS

//go:embed secret.html
var secretHtml embed.FS

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the theme name as a parameter.〒▽〒")
		return
	}
	themeName := os.Args[1]
	fmt.Println("Theme Name: ", themeName)
	err := data.CopyFile("AESDecrypt.js", filepath.Join("themes/", themeName, "/static/js/AESDecrypt.js"), aesDecryptScript)
	if err != nil {
		log.Fatal(err)
	}
	err = data.CopyFile("secret.html", filepath.Join("themes/", themeName, "/layouts/partials/secret.html", themeName), secretHtml)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("hugo")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// Output command execution results
	fmt.Println(string(output))

	// Get all passwords and content
	passwords, err := data.GetPasswords("./content")
	if err != nil {
		log.Fatal(err)
	}

	// Encrypt the password
	encryptedPasswords := crypto.GetEncryptedPasswords(passwords)

	encryptedContents := make(map[string]string)

	for file := range passwords {
		content := data.GetHTML(file)
		encrypted, err := crypto.AESEncrypt(content, encryptedPasswords[file])
		if err != nil {
			log.Fatal(err)
		}

		encryptedContents[file] = encrypted
		r, _ := regexp.Compile("content\\\\|\\.md$")
		fileName := r.ReplaceAllString(file, "")

		html.WriteEncryptedContentToHTML(fileName, encrypted)
	}
}
