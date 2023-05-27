package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func EncryptAES(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Generate a random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate random IV: %w", err)
	}

	// Create the cipher mode
	cipherText := make([]byte, len(plaintext))
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText, []byte(plaintext))

	// Concatenate the IV and cipher text
	result := append(iv, cipherText...)
	return base64.StdEncoding.EncodeToString(result), nil
}

func DecryptAES(key []byte, ciphertext string) (string, error) {
	encryptedData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to base64 decode ciphertext: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Split the IV and cipher text
	iv := encryptedData[:aes.BlockSize]
	cipherText := encryptedData[aes.BlockSize:]

	// Create the cipher mode
	plainText := make([]byte, len(cipherText))
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}
