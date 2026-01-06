---
inclusion: fileMatch
fileMatchPattern: "**/*.go"
---

# Go 开发规范

## 代码结构

### 文件组织

```
package xxx

// 常量定义
const (...)

// 类型定义
type Xxx struct {...}

// 构造函数
func New() *Xxx {...}

// 公开方法
func (x *Xxx) PublicMethod() {...}

// 私有方法
func (x *Xxx) privateMethod() {...}
```

### 错误处理

```go
// 推荐：返回 error
func DoSomething() error {
    if err := operation(); err != nil {
        return fmt.Errorf("operation failed: %w", err)
    }
    return nil
}

// 避免：使用 panic
func DoSomething() {
    if err := operation(); err != nil {
        panic(err) // ❌ 不推荐
    }
}
```

### HTTP 客户端

```go
// 使用项目统一的 HTTP 请求模式
func (c *Client) doRequest(method, path string, body, result interface{}) error {
    // 构建请求
    // 设置 Header
    // 发送请求
    // 解析响应
}
```

## 依赖管理

- 使用 Go Modules
- 最小化外部依赖
- 核心依赖：`github.com/mark3labs/mcp-go`

## 性能考虑

- 复用 HTTP Client
- 避免不必要的内存分配
- 使用 `sync.Pool` 处理频繁创建的对象
