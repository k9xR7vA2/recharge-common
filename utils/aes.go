package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// EncryptAES 使用 AES 进行加密
func EncryptAES(plaintext string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 使用 CBC 模式
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	// 生成随机 IV
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 使用 CBC 模式加密
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// 返回加密后的数据（base64 编码）
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES 使用 AES 进行解密
func DecryptAES(cipherText string, key string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// 使用 CBC 模式解密
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

// GenerateRandomKey 生成32字节（256位）的随机密钥
func GenerateRandomKey() (string, error) {
	key := make([]byte, 32) // 32字节 = 256位

	// 生成随机字节
	_, err := rand.Read(key)
	if err != nil {
		return "", fmt.Errorf("error generating random key: %v", err)
	}

	// 将字节转换为十六进制字符串表示
	return hex.EncodeToString(key), nil
}

// AesGCMEncrypt Encrypt AES GCM加密
func AesGCMEncrypt(plainText, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, plainText, nil)
	return hex.EncodeToString(cipherText), nil
}

// AesGCMDecrypt Decrypt AES GCM解密
func AesGCMDecrypt(cipherTextHex string, key []byte) ([]byte, error) {
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
