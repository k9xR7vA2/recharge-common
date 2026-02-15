package utils

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

// PaymentParams 表示支付参数结构
type PaymentParams struct {
	//BackendHost string `json:"backend_host"` // 后端服务器地址(含端口)
	OrderSN   string `json:"order_sn"`   // 订单编号
	TimeStamp int64  `json:"time_stamp"` // 时间戳
	Amount    string `json:"amount"`     //金额
	IsVerify  uint   `json:"is_verify"`  //是否需要手机号验证
}

// DecryptSign 解密加密的 sign 字符串，返回 PaymentParams 结构体
func DecryptSign(sign string, key string) (*PaymentParams, error) {
	// 解码 Base64 URL 编码的密文
	ciphertext, err := base64.URLEncoding.DecodeString(sign)
	if err != nil {
		return nil, fmt.Errorf("Base64解码失败: %w", err)
	}
	// 生成SHA256哈希作为AES密钥
	hash := sha256.Sum256([]byte(key))

	// 创建AES cipher
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return nil, fmt.Errorf("创建AES cipher失败: %w", err)
	}

	// 创建GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("创建GCM失败: %w", err)
	}

	// 分离nonce和真正的密文
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("密文长度不足")
	}
	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 解密
	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("解密失败: %w", err)
	}

	// 反序列化为结构体
	var params PaymentParams
	if err := json.Unmarshal(plaintext, &params); err != nil {
		return nil, fmt.Errorf("JSON反序列化失败: %w", err)
	}

	return &params, nil
}

func EncryptDataURLSafe(data interface{}, key string) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("JSON序列化失败: %w", err)
	}

	hash := sha256.Sum256([]byte(key))

	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return "", fmt.Errorf("创建AES cipher失败: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("创建GCM失败: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := crand.Read(nonce); err != nil {
		return "", fmt.Errorf("生成随机数失败: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

	// URL安全的Base64编码
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// GeneratePaymentURL 生成包含加密参数的收银台URL
func GeneratePaymentURL(baseURL, orderSN, amount, secretKey string, isVerify uint) (string, error) {
	// 创建支付参数
	params := PaymentParams{
		OrderSN:   orderSN,
		TimeStamp: time.Now().Add(time.Second * 40).Unix(),
		Amount:    amount,
		IsVerify:  isVerify,
	}
	// 加密生成sign
	sign, err := EncryptDataURLSafe(params, secretKey)
	if err != nil {
		return "", err
	}

	// 构建完整URL
	return fmt.Sprintf("%s?sign=%s", baseURL, sign), nil
}
