// Package central 提供 ZeroTier Central API 的 Go 客户端
package central

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client ZeroTier Central API 客户端接口
type Client interface {
	// Status 获取 Central 状态（包含当前用户信息）
	Status() (*CentralStatus, error)
	// Networks 网络管理
	Networks() NetworkService
}

// client 客户端实现
type client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Option 客户端配置选项
type Option func(*client)

// WithBaseURL 设置 API 地址
func WithBaseURL(url string) Option {
	return func(c *client) {
		c.baseURL = strings.TrimSuffix(url, "/")
	}
}

// WithToken 设置 API Token
func WithToken(token string) Option {
	return func(c *client) {
		c.token = token
	}
}

// WithTimeout 设置请求超时
func WithTimeout(timeout time.Duration) Option {
	return func(c *client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient 自定义 HTTP 客户端
func WithHTTPClient(hc *http.Client) Option {
	return func(c *client) {
		c.httpClient = hc
	}
}

// New 创建新的 Central API 客户端
func New(token string, opts ...Option) Client {
	c := &client{
		baseURL:    "https://api.zerotier.com/api/v1",
		token:      token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// do 执行 HTTP 请求
func (c *client) do(method, path string, body io.Reader) ([]byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "token "+c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("api error: %s - %s", resp.Status, string(data))
	}

	return data, nil
}

// Status 获取 Central 状态
func (c *client) Status() (*CentralStatus, error) {
	data, err := c.do(http.MethodGet, "/status", nil)
	if err != nil {
		return nil, err
	}

	var status CentralStatus
	if err := json.Unmarshal(data, &status); err != nil {
		return nil, fmt.Errorf("unmarshal status: %w", err)
	}

	return &status, nil
}

// Networks 返回网络服务
func (c *client) Networks() NetworkService {
	return &networkService{client: c}
}
