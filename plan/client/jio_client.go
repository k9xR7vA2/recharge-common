package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	jioBaseURL = "https://www.jio.com"
)

// JioClient Jio专用客户端（带 Session/Cookie 管理）
type JioClient struct {
	client      *resty.Client
	cookieJar   []*http.Cookie
	cookieLock  sync.RWMutex
	sessionInit bool
}

// NewJioClient 创建 Jio 客户端
func NewJioClient(timeout time.Duration) *JioClient {
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	client := resty.New().
		SetTimeout(timeout).
		SetHeader("Host", "www.jio.com").
		SetHeader("Pragma", "no-cache").
		SetHeader("Cache-Control", "no-cache").
		SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 18_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.5 Mobile/15E148 Safari/604.1").
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Sec-Fetch-Site", "same-origin").
		SetHeader("Sec-Fetch-Mode", "cors").
		SetHeader("Sec-Fetch-Dest", "empty").
		SetHeader("Referer", "https://www.jio.com/").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9")

	return &JioClient{
		client:    client,
		cookieJar: make([]*http.Cookie, 0),
	}
}

// InitSession 初始化会话（验证手机号并获取 cookie）
func (c *JioClient) InitSession(ctx context.Context, phoneNumber string) error {
	url := fmt.Sprintf("%s/api/jio-recharge-service/recharge/mobility/number/%s", jioBaseURL, phoneNumber)
	resp, err := c.client.R().
		SetContext(ctx).
		Get(url)

	if err != nil {
		return errors.New("init session failed")
	}

	if resp.StatusCode() != 200 {
		return errors.New("init session status failed")
	}

	// 保存 cookies
	c.cookieLock.Lock()
	c.cookieJar = resp.Cookies()
	c.sessionInit = true
	c.cookieLock.Unlock()

	return nil
}

// GetWithSession 带 session 的 GET 请求
func (c *JioClient) GetWithSession(ctx context.Context, url string, referer string) ([]byte, error) {
	c.cookieLock.RLock()
	if !c.sessionInit {
		c.cookieLock.RUnlock()
		return nil, fmt.Errorf("session not initialized, call InitSession first")
	}
	cookies := c.cookieJar
	c.cookieLock.RUnlock()

	req := c.client.R().
		SetContext(ctx).
		SetCookies(cookies)

	if referer != "" {
		req.SetHeader("Referer", referer)
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	// 更新 cookies（如果有新的）
	if len(resp.Cookies()) > 0 {
		c.updateCookies(resp.Cookies())
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("request failed, status: %d", resp.StatusCode())
	}

	return resp.Body(), nil
}

// updateCookies 更新 cookie（合并新旧）
func (c *JioClient) updateCookies(newCookies []*http.Cookie) {
	c.cookieLock.Lock()
	defer c.cookieLock.Unlock()

	cookieMap := make(map[string]*http.Cookie)

	// 先放旧的
	for _, cookie := range c.cookieJar {
		cookieMap[cookie.Name] = cookie
	}

	// 新的覆盖旧的
	for _, cookie := range newCookies {
		cookieMap[cookie.Name] = cookie
	}

	// 重建数组
	c.cookieJar = make([]*http.Cookie, 0, len(cookieMap))
	for _, cookie := range cookieMap {
		c.cookieJar = append(c.cookieJar, cookie)
	}
}

// IsSessionInit 检查 session 是否已初始化
func (c *JioClient) IsSessionInit() bool {
	c.cookieLock.RLock()
	defer c.cookieLock.RUnlock()
	return c.sessionInit
}

// ClearSession 清除 session
func (c *JioClient) ClearSession() {
	c.cookieLock.Lock()
	defer c.cookieLock.Unlock()
	c.cookieJar = make([]*http.Cookie, 0)
	c.sessionInit = false
}
