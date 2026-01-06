# 添加 API 功能工作流

## 概述

为 ZeroTier SDK 添加新的 API 功能的标准工作流。

## 步骤

### 1. 需求分析

- [ ] 确定功能属于 Client 还是 Central 模块
- [ ] 查阅 ZeroTier API 文档
- [ ] 确定 API 端点和参数

### 2. 类型定义

在 `types.go` 中添加：

```go
// 请求类型
type XxxRequest struct {
    Field string `json:"field"`
}

// 响应类型
type XxxResponse struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

### 3. 实现 API 方法

在对应模块中添加：

```go
func (s *Service) Xxx(req *XxxRequest) (*XxxResponse, error) {
    var resp XxxResponse
    err := s.client.doRequest("POST", "/xxx", req, &resp)
    return &resp, err
}
```

### 4. 添加 Builder（如需要）

在 `builder.go` 中添加：

```go
type XxxBuilder struct {
    config *XxxConfig
}

func NewXxxConfig() *XxxBuilder {
    return &XxxBuilder{config: &XxxConfig{}}
}

func (b *XxxBuilder) Field(value string) *XxxBuilder {
    b.config.Field = value
    return b
}

func (b *XxxBuilder) Build() *XxxConfig {
    return b.config
}
```

### 5. 导出到主包

在 `zerotier.go` 中添加类型别名和构造函数。

### 6. 编写测试

创建 `xxx_test.go`：

```go
func TestXxx(t *testing.T) {
    // 测试代码
}
```

### 7. 更新文档

- 更新模块 README.md
- 更新主 README.md
- 添加使用示例

### 8. 添加 MCP 工具（可选）

如果需要 MCP 支持，在 `mcp/server.go` 中添加工具。

## 检查清单

- [ ] 类型定义完整
- [ ] API 方法实现
- [ ] Builder 模式（复杂配置）
- [ ] 主包导出
- [ ] 单元测试
- [ ] 文档更新
- [ ] MCP 工具（可选）
