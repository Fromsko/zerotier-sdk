# ZeroTier Go SDK 项目指南

## 项目概述

这是一个 ZeroTier API 的 Go SDK，提供本地节点管理和云端控制面板的统一接口。

## 核心模块

| 模块          | 用途                  | 地址                      |
| ------------- | --------------------- | ------------------------- |
| `client/`     | 本地 Service API      | localhost:9993            |
| `central/`    | 云端 Central API v1   | api.zerotier.com          |
| `central/v2/` | 云端 Central API v2   | central.zerotier.com      |
| `mcp/`        | MCP 服务集成          | stdio                     |
| `zerotier.go` | 统一入口              | -                         |

## 代码规范

### Go 代码风格

- 使用 `gofmt` 格式化代码
- 遵循 Go 官方代码规范
- 所有导出的函数和类型必须有中文注释
- 错误处理使用 `error` 返回，不使用 panic

### 命名规范

- 包名：小写单词，如 `client`, `central`
- 类型名：驼峰命名，如 `NetworkConfig`
- 函数名：驼峰命名，如 `NewClient`
- 常量：全大写下划线，如 `DEFAULT_BASE_URL`

### Builder 模式

项目大量使用 Builder 模式，新增配置类型时应遵循：

```go
type XxxBuilder struct {
    config *XxxConfig
}

func NewXxx() *XxxBuilder {
    return &XxxBuilder{config: &XxxConfig{}}
}

func (b *XxxBuilder) FieldName(value Type) *XxxBuilder {
    b.config.FieldName = value
    return b
}

func (b *XxxBuilder) Build() *XxxConfig {
    return b.config
}
```

## 构建命令

```bash
# 构建所有包
make build

# 构建 MCP 二进制
make build-mcp

# 运行测试
make test

# 代码检查
make lint

# 清理
make clean
```

## 测试规范

- 测试文件命名：`xxx_test.go`
- 使用 `testing` 标准库
- 表驱动测试优先
- Mock 外部 API 调用

## 文档要求

- README.md 使用中文
- 代码注释使用中文
- API 文档包含使用示例
