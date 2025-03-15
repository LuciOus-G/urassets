package Utilities

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"io"
	"strings"
)

const (
	timeCost  = 1
	memory    = 64 * 1024
	threads   = 4
	keyLength = 32
	saltSize  = 16
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, timeCost, memory, threads, keyLength)
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s..%s", hash, salt))), nil
}

func CheckPassword(plain string, password string) (bool, error) {
	decoded, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		return false, err
	}

	passwordHash := strings.Split(string(decoded), "..")[0]
	salt := []byte(strings.Split(string(decoded), "..")[1])

	compare := argon2.IDKey([]byte(plain), salt, timeCost, memory, threads, keyLength)
	return string(compare) == passwordHash, nil
}

// PKCS7Padding pads the input to be a multiple of the block size
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding removes the padding after decryption
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	padding := int(data[length-1])

	if (length - padding) < 0 {
		return nil, errors.New("Data is too short, or key and iv doesnt match")
	}

	return data[:(length - padding)], nil
}

func EncryptAES256Rest(plainText string, key string) (string, error) {
	keyHash := sha256.Sum256([]byte(key)) // Ensure the key is 32 bytes long (256 bits)
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return "", err
	}

	paddedData := PKCS7Padding([]byte(plainText), aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)

	return hex.EncodeToString(ciphertext), nil
}

func DecryptAES256Rest(cipherTextHex, key string) (string, error) {
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return "", err
	}

	keyHash := sha256.Sum256([]byte(key)) // Ensure the key is 32 bytes long (256 bits)
	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	unPaddingData, err := PKCS7UnPadding(cipherText)
	if err != nil {
		return "", err
	}

	return string(unPaddingData), nil
}
