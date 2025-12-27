// Package client 提供 ZeroTier Service API（本地节点）的 Go 客户端
package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// Client ZeroTier Service API 客户端接口
type Client interface {
	// Status 获取节点状态
	Status() (*NodeStatus, error)
	// Networks 网络管理
	Networks() NetworkService
	// Peers 节点管理
	Peers() PeerService
	// Controller 控制器管理（自托管时可用）
	Controller() ControllerService
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

// WithToken 设置认证 Token
func WithToken(token string) Option {
	return func(c *client) {
		c.token = token
	}
}

// WithTokenFile 从文件读取 Token
func WithTokenFile(path string) Option {
	return func(c *client) {
		data, err := os.ReadFile(path)
		if err == nil {
			c.token = strings.TrimSpace(string(data))
		}
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

// New 创建新的 ZeroTier Service API 客户端
func New(opts ...Option) Client {
	c := &client{
		baseURL:    "http://localhost:9993",
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}

	for _, opt := range opts {
		opt(c)
	}

	// 如果没有设置 token，尝试从默认位置读取
	if c.token == "" {
		c.token = readDefaultToken()
	}

	return c
}

// readDefaultToken 从系统默认位置读取 token
func readDefaultToken() string {
	var path string
	switch runtime.GOOS {
	case "windows":
		path = `C:\ProgramData\ZeroTier\One\authtoken.secret`
	case "darwin":
		home, _ := os.UserHomeDir()
		path = home + "/Library/Application Support/ZeroTier/authtoken.secret"
	default: // linux
		path = "/var/lib/zerotier-one/authtoken.secret"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

// do 执行 HTTP 请求
func (c *client) do(method, path string, body io.Reader) ([]byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("X-ZT1-AUTH", c.token)
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

// Status 获取节点状态
func (c *client) Status() (*NodeStatus, error) {
	data, err := c.do(http.MethodGet, "/status", nil)
	if err != nil {
		return nil, err
	}

	var status NodeStatus
	if err := json.Unmarshal(data, &status); err != nil {
		return nil, fmt.Errorf("unmarshal status: %w", err)
	}

	return &status, nil
}

// Networks 返回网络服务
func (c *client) Networks() NetworkService {
	return &networkService{client: c}
}

// Peers 返回节点服务
func (c *client) Peers() PeerService {
	return &peerService{client: c}
}

// Controller 返回控制器服务
func (c *client) Controller() ControllerService {
	return &controllerService{client: c}
}
