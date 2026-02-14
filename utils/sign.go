package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

// GenerateAppID 生成Appid
func GenerateAppID() (string, error) {
	const appIDLength = 16
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	appID := make([]byte, appIDLength)
	_, err := rand.Read(appID)
	if err != nil {
		return "", err
	}
	for i, b := range appID {
		appID[i] = charset[b%byte(len(charset))]
	}
	return string(appID), nil
}

// GenerateSecret 生成secret
func GenerateSecret(plainText, appID string) (string, error) {
	return AesGCMEncrypt([]byte(plainText), []byte(appID))
}

// GenerateSignature 签名算法
func GenerateSignature(params map[string]string, secret string) (string, error) {
	var keys []string
	for key := range params {
		if key != "signature" { // Exclude the signature itself from the sorted keys
			keys = append(keys, key)
		}
	}
	sort.Strings(keys) // Sort keys alphabetically
	// Concatenate sorted parameters into key=value format
	var sortedParams strings.Builder
	for _, key := range keys {
		sortedParams.WriteString(key + "=" + url.QueryEscape(params[key]) + "&")
	}
	concatenatedParams := strings.TrimRight(sortedParams.String(), "&") // Remove trailing &
	// Append secret key and create HMAC-SHA256 hash
	dataToSign := concatenatedParams + secret
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(dataToSign))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature, nil
}

// VerifySignature 验证签名
func VerifySignature(params map[string]string, providedSignature, secret string) (bool, error) {
	// Generate a new signature from the received parameters (excluding the provided signature itself)
	calculatedSignature, err := GenerateSignature(params, secret)
	if err != nil {
		return false, err
	}
	// Compare the calculated signature with the provided one
	return calculatedSignature == providedSignature, nil
}
