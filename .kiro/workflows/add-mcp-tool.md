# 添加 MCP 工具工作流

## 概述

为 ZeroTier MCP 服务添加新工具的标准工作流。

## 步骤

### 1. 确定工具类型

- 本地 API 工具：`zt_xxx`
- 云端 API 工具：`zt_central_xxx`

### 2. 定义工具

在 `mcp/server.go` 中添加：

```go
// 定义工具
toolXxx := mcp.NewTool("zt_xxx",
    mcp.WithDescription("工具描述"),
    mcp.WithString("param_name",
        mcp.Required(),
        mcp.Description("参数描述"),
    ),
)
```

### 3. 实现处理函数

```go
func (s *Server) handleXxx(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    // 1. 解析参数
    param, ok := req.Params.Arguments["param_name"].(string)
    if !ok || param == "" {
        return nil, fmt.Errorf("param_name is required")
    }

    // 2. 调用 SDK
    result, err := s.localClient.Xxx(param)
    if err != nil {
        return nil, fmt.Errorf("xxx failed: %w", err)
    }

    // 3. 格式化输出
    data, err := json.MarshalIndent(result, "", "  ")
    if err != nil {
        return nil, fmt.Errorf("marshal failed: %w", err)
    }

    return mcp.NewToolResultText(string(data)), nil
}
```

### 4. 注册工具

在 `registerTools()` 方法中添加：

```go
s.server.AddTool(toolXxx, s.handleXxx)
```

### 5. 更新文档

在 `mcp/README.md` 中添加工具说明：

| 工具     | 描述     | 参数       |
| -------- | -------- | ---------- |
| `zt_xxx` | 工具描述 | param_name |

### 6. 测试工具

```bash
# 构建
go build -o zerotier-mcp ./cmd/mcp

# 测试运行
./zerotier-mcp
```

## 检查清单

- [ ] 工具命名符合规范
- [ ] 参数定义完整
- [ ] 错误处理完善
- [ ] 返回格式正确
- [ ] 文档已更新
- [ ] 测试通过
