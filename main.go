package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hotjuicew/hugoArticleEncryptor/data"
)

//go:embed AESDecrypt.js
var aesDecryptScript embed.FS

//go:embed secret.html
var secretHtml embed.FS

func main() {

	dir1 := "static/js"
	err := os.MkdirAll(dir1, os.ModePerm)
	if err != nil {
		log.Println(" ", err)
		return
	}
	dir2 := "layouts/shortcodes"
	err = os.MkdirAll(dir2, os.ModePerm)
	if err != nil {
		log.Println(" \"layouts/shortcodes\" create fail:", err)
		return
	}
	err = data.CopyFile("AESDecrypt.js", filepath.Join("static/js/AESDecrypt.js"), aesDecryptScript)
	if err != nil {
		log.Fatalf("data.CopyFile: AESDecrypt.js gets error %v", err)
	}
	err = data.CopyFile("secret.html", filepath.Join("layouts/shortcodes/secret.html"), secretHtml)
	if err != nil {
		log.Fatal("data.CopyFile: secret.html gets error", err)
	}

	output, err := exec.Command("hugo").Output()
	if err != nil {
		log.Fatalln("cmd.Output() gets error", err)
	}

	// Output command execution results
	fmt.Println(string(output))

	// solve html files in public folder
	err = data.WalkHTMLFiles()
	if err != nil {
		log.Fatal("Error:", err)
	}

}
