package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

func GetEncryptedPasswords(passwords map[string]string) map[string][]byte {
	encryptedPasswords := make(map[string][]byte)
	for file, password := range passwords {
		hash := sha256.New()
		hash.Write([]byte(password))
		encryptedPasswords[file] = hash.Sum(nil)
	}
	return encryptedPasswords
}

// AESEncrypt 使用AES-GCM算法加密明文
func AESEncrypt(plaintext string, password []byte) (string, error) {
	// 创建AES的BlockCipher
	block, err := aes.NewCipher(password)
	if err != nil {
		return "", err
	}

	// 创建AES-GCM
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成nonce
	hash := sha256.Sum256(password)
	nonce := hash[:aesgcm.NonceSize()]

	// 加密
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)

	// 拼接nonce和密文,返回
	return hex.EncodeToString(nonce) + "|" + hex.EncodeToString(ciphertext), nil
}

// AESDecrypt 使用AES-GCM算法解密密文
func AESDecrypt(ciphertext string, password []byte) (string, error) {
	// 创建AES的BlockCipher
	block, err := aes.NewCipher(password)
	if err != nil {
		return "", err
	}

	// 创建AES-GCM
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 将密文分割成nonce和密文部分
	parts := strings.Split(ciphertext, "|")
	if len(parts) != 2 {
		return "", errors.New("invalid ciphertext format")
	}

	// 解码nonce和密文
	decodedNonce, err := hex.DecodeString(parts[0])
	fmt.Println("decodedNonce", decodedNonce)
	if err != nil {
		return "", err
	}
	decodedCiphertext, err := hex.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	// 解密
	plaintext, err := aesgcm.Open(nil, decodedNonce, decodedCiphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
