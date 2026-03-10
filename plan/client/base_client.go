package client

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

// BaseClient 基础HTTP客户端接口
type BaseClient interface {
	Get(ctx context.Context, url string, params map[string]string, headers map[string]string) (*Response, error)
	Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*Response, error)
	SetTimeout(timeout time.Duration)
	SetProxy(proxy string)
	SetFollowRedirect(follow bool)
}

// Response 统一响应结构
type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
}

// RestyClient 基于 resty 的基础客户端
type RestyClient struct {
	client *resty.Client
}

// NewRestyClient 创建基础客户端
func NewRestyClient(timeout time.Duration) *RestyClient {
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	client := resty.New().
		SetTimeout(timeout).
		SetRetryCount(0)

	return &RestyClient{client: client}
}

func (c *RestyClient) Get(ctx context.Context, url string, params map[string]string, headers map[string]string) (*Response, error) {
	req := c.client.R().SetContext(ctx)

	if params != nil {
		req.SetQueryParams(params)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}

	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	return c.buildResponse(resp), nil
}

func (c *RestyClient) Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*Response, error) {
	req := c.client.R().SetContext(ctx)

	if body != nil {
		req.SetBody(body)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}

	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}

	return c.buildResponse(resp), nil
}

func (c *RestyClient) SetTimeout(timeout time.Duration) {
	c.client.SetTimeout(timeout)
}

func (c *RestyClient) SetProxy(proxy string) {
	if proxy != "" {
		c.client.SetProxy(proxy)
	}
}

func (c *RestyClient) SetFollowRedirect(follow bool) {
	if !follow {
		c.client.SetRedirectPolicy(resty.NoRedirectPolicy())
	} else {
		c.client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
	}
}

func (c *RestyClient) GetRestyClient() *resty.Client {
	return c.client
}

func (c *RestyClient) buildResponse(resp *resty.Response) *Response {
	headers := make(map[string]string)
	for k, v := range resp.Header() {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	return &Response{
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
		Headers:    headers,
	}
}
