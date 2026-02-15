package utils

import (
	"crypto/rand"
	"math/big"
	"sync"
)

// SafeRandom 并发安全的随机数生成器
type SafeRandom struct {
	mu sync.Mutex
}

func NewSafeRandom() *SafeRandom {
	return &SafeRandom{}
}

// Float64 生成 [0.0, 1.0) 范围的随机浮点数
func (r *SafeRandom) Float64() float64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 使用 crypto/rand 生成随机数
	// 生成 [0, 100000) 范围的随机整数然后转换为浮点数
	n, err := rand.Int(rand.Reader, big.NewInt(100000))
	if err != nil {
		// 如果生成随机数失败，返回一个默认值
		return 0.5
	}

	return float64(n.Int64()) / 100000.0
}
