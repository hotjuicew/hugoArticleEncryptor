package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

func GetEncryptedPassword(password string) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	encryptedPassword := hash.Sum(nil)
	return encryptedPassword
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
