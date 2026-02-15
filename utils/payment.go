package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"strings"
)

type TokenInfo struct {
	OrderSn  string
	TenantId uint
}

const (
	// AES-128, key必须是16字节
	secretKey = "p@ssw0rd#2024$Key"
)

// EncryptToken 加密生成token
func EncryptToken(orderSn string, tenantId uint) (string, error) {
	// 1. 将数据拼接成字符串 (用特殊字符分隔,确保不会出现在原始数据中)
	plainText := fmt.Sprintf("%s|%d", orderSn, tenantId)

	// 2. 加密
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// 3. 填充数据(PKCS7)
	blockSize := block.BlockSize()
	padding := blockSize - len(plainText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plainBytes := append([]byte(plainText), padText...)

	// 4. 加密(CBC模式)
	cipherText := make([]byte, len(plainBytes))
	iv := []byte(secretKey)[:blockSize] // 使用密钥前16位作为IV
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainBytes)

	// 5. 转换为Hex字符串
	return hex.EncodeToString(cipherText), nil
}

// DecryptToken 解密token
func DecryptToken(token string) (*TokenInfo, error) {
	// 1. Hex解码
	cipherText, err := hex.DecodeString(token)
	if err != nil {
		return nil, err
	}

	// 2. 解密
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	// 3. CBC模式解密
	plainText := make([]byte, len(cipherText))
	iv := []byte(secretKey)[:block.BlockSize()]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	// 4. 去除填充
	padding := int(plainText[len(plainText)-1])
	plainText = plainText[:len(plainText)-padding]

	// 5. 解析数据
	parts := strings.Split(string(plainText), "|")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid token format")
	}

	// 6. 转换数据类型
	tenantId := 0
	_, err = fmt.Sscanf(parts[1], "%d", &tenantId)
	if err != nil {
		return nil, err
	}

	return &TokenInfo{
		OrderSn:  parts[0],
		TenantId: uint(tenantId),
	}, nil
}
