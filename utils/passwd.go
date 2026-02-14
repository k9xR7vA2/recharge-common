package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	lowerChars     = "abcdefghijklmnopqrstuvwxyz"
	upperChars     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars    = "0123456789"
	specialChars   = "!@#$%^&*()-_=+[]{}|;:,.<>?/~"
	allChars       = lowerChars + upperChars + numberChars + specialChars
	passwordLength = 12
)

// GenerateRandomPassword  生成一个包含大小写字母、数字和特殊字符的随机密码
func GenerateRandomPassword(length int, passwdType []int) (string, error) {
	ptMap := map[int]string{1: lowerChars, 2: upperChars, 3: numberChars, 4: specialChars}
	pt := ""
	for _, v := range passwdType {
		if data, ok := ptMap[v]; ok {
			pt += data
		}
	}
	if pt == "" {
		return "", errors.New("生成密码类型必须")
	}
	password := make([]byte, length)
	for i := range password {
		char, err := randCharFrom(pt)
		if err != nil {
			return "", err
		}
		password[i] = char
	}
	return string(password), nil
}

// randCharFrom 从给定的字符集中随机选择一个字符
func randCharFrom(charset string) (byte, error) {
	maxInt := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}
