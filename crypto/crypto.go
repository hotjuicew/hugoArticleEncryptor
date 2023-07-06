package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
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

// AESEncrypt Encrypt plaintext using AES-GCM algorithm
func AESEncrypt(plaintext string, password []byte) (string, error) {
	block, err := aes.NewCipher(password)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(password)
	nonce := hash[:aesgcm.NonceSize()]
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return hex.EncodeToString(nonce) + "|" + hex.EncodeToString(ciphertext), nil
}
