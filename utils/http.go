package utils

import (
	"github.com/go-resty/resty/v2"
	"net"
	"time"
)

func GetHttpClient() *resty.Client {
	// HTTP 客户端配置
	client := resty.New().
		SetTimeout(5*time.Second). // 单次 HTTP 请求超时较短
		SetHeader("Content-Type", "application/json").
		EnableTrace()
	// 只对网络类临时错误进行 HTTP 重试
	client.AddRetryCondition(func(r *resty.Response, err error) bool {
		// 网络超时
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				return true
			}
		}
		// 服务端临时错误
		return r != nil && r.StatusCode() >= 500
	})
	// HTTP 重试配置（针对网络抖动的快速重试）
	client.SetRetryCount(2). // HTTP 重试次数较少
					SetRetryWaitTime(100 * time.Millisecond).   // 重试等待时间短
					SetRetryMaxWaitTime(500 * time.Millisecond) // 最大等待时间短
	return client
}
