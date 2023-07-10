package html

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func convertToLower(str string) string {
	var result strings.Builder

	for _, char := range str {
		if unicode.IsUpper(char) && unicode.IsLetter(char) {
			char = unicode.ToLower(char)
		}
		result.WriteRune(char)
	}

	return result.String()
}
func WriteEncryptedContentToHTML(folderName string, encryptedText string) {
	folderNameLower := convertToLower(folderName)
	folderPath1 := filepath.Join("public", "post", folderNameLower)
	folderPath2 := filepath.Join("public", "posts", folderNameLower)

	// Check if folderPath1 exists
	if _, err := os.Stat(folderPath1); err == nil {
		updateHTMLFiles(folderPath1, encryptedText)
		return
	}

	// Check if folderPath2 exists
	if _, err := os.Stat(folderPath2); err == nil {
		updateHTMLFiles(folderPath2, encryptedText)
		return
	}

	fmt.Println("Folder does not exist:", folderNameLower)
}

func updateHTMLFiles(folderPath string, encryptedText string) {
	// Iterate through the HTML files in the folder
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			doc, err := goquery.NewDocumentFromReader(file)
			if err != nil {
				return err
			}

			// Find div elements with id=secret
			doc.Find("#secret").Each(func(i int, s *goquery.Selection) {
				s.SetText(encryptedText)
			})

			// Add a <script> tag to the <body> tag of an HTML file that references an external JavaScript file
			doc.Find("body").AppendHtml(`<script src="https://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.9-1/crypto-js.js"></script>`)

			// Add a <script> tag that references '/static/js/decrypt.js'
			doc.Find("body").AppendHtml(`<script src="../../js/AESDecrypt.js"></script>`)

			// Get the modified HTML content
			updatedHTML, err := doc.Html()
			if err != nil {
				return err
			}

			// Write the modified HTML content to the file
			err = os.WriteFile(path, []byte(updatedHTML), 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal("filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) gets error", err)
	}
}
