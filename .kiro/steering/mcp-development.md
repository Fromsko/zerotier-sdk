---
inclusion: fileMatch
fileMatchPattern: "**/mcp/**/*.go"
---

# MCP 开发规范

## MCP 工具定义

### 工具命名

- 本地 API 工具：`zt_xxx`（如 `zt_status`, `zt_networks`）
- 云端 API 工具：`zt_central_xxx`（如 `zt_central_networks`）

### 工具结构

```go
// 定义工具
tool := mcp.NewTool("zt_xxx",
    mcp.WithDescription("工具描述"),
    mcp.WithString("param_name",
        mcp.Required(),
        mcp.Description("参数描述"),
    ),
)

// 注册处理器
server.AddTool(tool, handleXxx)

// 处理函数
func handleXxx(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    // 解析参数
    // 调用 SDK
    // 返回结果
}
```

### 返回格式

```go
// 成功返回
return mcp.NewToolResultText(jsonString), nil

// 错误返回
return nil, fmt.Errorf("error message: %w", err)
```

## 环境变量

| 变量               | 说明              | 默认值   |
| ------------------ | ----------------- | -------- |
| `ZT_CENTRAL_TOKEN` | Central API Token | -        |
| `ZT_LOCAL_TOKEN`   | 本地 API Token    | 自动读取 |

## 测试 MCP 工具

```bash
# 构建
go build -o zerotier-mcp ./cmd/mcp

# 测试运行
./zerotier-mcp
```
