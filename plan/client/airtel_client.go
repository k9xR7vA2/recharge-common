package client

import (
	"context"
	"crypto/des"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// AirtelClient Airtel专用客户端（带加解密）
type AirtelClient struct {
	*RestyClient
}

// NewAirtelClient 创建 Airtel 客户端
func NewAirtelClient(timeout time.Duration) *AirtelClient {
	base := NewRestyClient(timeout)

	// 设置 Airtel 默认请求头
	base.client.
		SetHeader("Host", "digi-api.airtel.in").
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("x-at-application", "recast").
		SetHeader("x-at-client", "WEB").
		SetHeader("googleCookie", "airtel.com").
		SetHeader("Origin", "https://www.airtel.in").
		SetHeader("Referer", "https://www.airtel.in/").
		SetHeader("X-Consumer-Name", "AirtelIn").
		SetHeader("requesterId", "WEB").
		SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 16_7_10 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1").
		SetHeader("Sec-Fetch-Site", "same-site").
		SetHeader("Sec-Fetch-Mode", "cors").
		SetHeader("Sec-Fetch-Dest", "empty").
		SetHeader("Accept-Language", "zh-CN,zh-Hans;q=0.9")

	return &AirtelClient{RestyClient: base}
}

// GetDecrypted 发起GET请求并自动解密响应
func (c *AirtelClient) GetDecrypted(ctx context.Context, url string, params map[string]string) ([]byte, error) {
	// 生成动态请求头
	headers := map[string]string{
		"adsHeader":  GenerateAdsHeader(),
		"x-at-reqid": uuid.New().String(),
	}

	resp, err := c.Get(ctx, url, params, headers)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http status: %d", resp.StatusCode)
	}

	// 从响应头获取解密密钥
	googleCookie := resp.Headers["Googlecookie"]
	if googleCookie == "" {
		googleCookie = resp.Headers["googleCookie"]
	}
	if googleCookie == "" {
		googleCookie = resp.Headers["googlecookie"]
	}
	if googleCookie == "" || len(googleCookie) < 8 {
		return nil, fmt.Errorf("missing googleCookie in response header")
	}

	// 解密响应
	encryptedData := strings.Trim(string(resp.Body), "\"")
	decrypted, err := DecryptDES(encryptedData, googleCookie[:8])
	if err != nil {
		return nil, fmt.Errorf("decrypt failed: %w", err)
	}

	return []byte(decrypted), nil
}

// GenerateAdsHeader 生成 adsHeader (SHA1)
func GenerateAdsHeader() string {
	now := time.Now()
	str := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(),
		int(now.Month())-1,
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond()/1e6,
		rand.Int63(),
	)
	hash := sha1.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// DecryptDES 使用 DES ECB 模式解密
func DecryptDES(encrypted string, key string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %w", err)
	}

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("create cipher failed: %w", err)
	}

	blockSize := block.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return "", fmt.Errorf("ciphertext length invalid")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += blockSize {
		block.Decrypt(plaintext[i:i+blockSize], ciphertext[i:i+blockSize])
	}

	plaintext = pkcs7Unpad(plaintext)
	return string(plaintext), nil
}

func pkcs7Unpad(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding > len(data) || padding > 8 {
		return data
	}
	return data[:len(data)-padding]
}
