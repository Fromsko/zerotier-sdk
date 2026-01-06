---
inclusion: fileMatch
fileMatchPattern: "**/client/**/*.go,**/central/**/*.go"
---

# API 客户端开发规范

## 客户端结构

### 基础客户端

```go
type Client struct {
    baseURL    string
    token      string
    httpClient *http.Client
}

func New(opts ...Option) *Client {
    c := &Client{
        baseURL:    defaultBaseURL,
        httpClient: &http.Client{Timeout: 30 * time.Second},
    }
    for _, opt := range opts {
        opt(c)
    }
    return c
}
```

### Option 模式

```go
type Option func(*Client)

func WithBaseURL(url string) Option {
    return func(c *Client) {
        c.baseURL = strings.TrimSuffix(url, "/")
    }
}

func WithToken(token string) Option {
    return func(c *Client) {
        c.token = token
    }
}
```

## API 方法规范

### 列表方法

```go
func (s *Service) List() ([]Item, error) {
    var items []Item
    err := s.client.doRequest("GET", "/path", nil, &items)
    return items, err
}
```

### 获取单个

```go
func (s *Service) Get(id string) (*Item, error) {
    var item Item
    err := s.client.doRequest("GET", "/path/"+id, nil, &item)
    return &item, err
}
```

### 创建

```go
func (s *Service) Create(config *Config) (*Item, error) {
    var item Item
    err := s.client.doRequest("POST", "/path", config, &item)
    return &item, err
}
```

### 更新

```go
func (s *Service) Update(id string, config *Config) (*Item, error) {
    var item Item
    err := s.client.doRequest("POST", "/path/"+id, config, &item)
    return &item, err
}
```

### 删除

```go
func (s *Service) Delete(id string) error {
    return s.client.doRequest("DELETE", "/path/"+id, nil, nil)
}
```

## 错误处理

```go
// API 错误类型
type APIError struct {
    StatusCode int
    Message    string
}

func (e *APIError) Error() string {
    return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
}
```

## 认证

### Client (本地)

- Header: `X-ZT1-Auth: <token>`
- Token 来源: `authtoken.secret` 文件

### Central (云端)

- Header: `Authorization: token <api_token>`
- Token 来源: 用户配置
