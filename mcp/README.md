# MCP - Model Context Protocol 集成

为 ZeroTier SDK 提供 MCP 服务支持，可与 Claude、Cursor 等 AI 工具集成。

## 安装

### 方式一：使用安装脚本（推荐）

**Linux/macOS:**

```bash
curl -fsSL https://raw.githubusercontent.com/fromsko/zerotier-sdk/main/scripts/install-mcp.sh | bash
# 或指定版本
curl -fsSL https://raw.githubusercontent.com/fromsko/zerotier-sdk/main/scripts/install-mcp.sh | bash -s v1.0.0
```

**Windows (PowerShell):**

```powershell
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/fromsko/zerotier-sdk/main/scripts/install-mcp.ps1" -OutFile install-mcp.ps1
.\install-mcp.ps1
```

### 方式二：手动下载

从 [Releases](https://github.com/fromsko/zerotier-sdk/releases) 页面下载对应平台的二进制文件。

### 方式三：从源码构建

```bash
git clone https://github.com/fromsko/zerotier-sdk.git
cd zerotier-sdk
make build-mcp
# 二进制文件在 dist/ 目录
```

## 快速开始

```go
package main

import (
    "log"
    "github.com/fromsko/zerotier-sdk/mcp"
)

func main() {
    // 创建 MCP 服务（仅本地 API）
    s := mcp.New("zerotier", "1.0.0")

    // 启动服务
    if err := s.ServeStdio(); err != nil {
        log.Fatal(err)
    }
}
```

## 配置选项

```go
// 同时启用本地和云端 API
s := mcp.New("zerotier", "1.0.0",
    mcp.WithCentralToken("your_api_token"),
)

// 自定义客户端
s := mcp.New("zerotier", "1.0.0",
    mcp.WithLocalClient(myLocalClient),
    mcp.WithCentralClient(myCentralClient),
)
```

## 可用工具

### 本地 API

| 工具          | 描述             |
| ------------- | ---------------- |
| `zt_status`   | 获取节点状态     |
| `zt_networks` | 列出已加入的网络 |
| `zt_join`     | 加入网络         |
| `zt_leave`    | 离开网络         |
| `zt_peers`    | 列出 Peers       |

### 云端 API（需配置 Token）

| 工具                     | 描述         |
| ------------------------ | ------------ |
| `zt_central_networks`    | 列出云端网络 |
| `zt_central_members`     | 列出网络成员 |
| `zt_central_authorize`   | 授权成员     |
| `zt_central_deauthorize` | 取消授权     |

## Claude Desktop 配置

编辑 `claude_desktop_config.json`：

```json
{
  "mcpServers": {
    "zerotier": {
      "command": "/path/to/zerotier-mcp",
      "env": {
        "ZT_CENTRAL_TOKEN": "your_api_token"
      }
    }
  }
}
```

## 构建 MCP 服务

```bash
cd cmd/mcp
go build -o zerotier-mcp
```

## 本地构建

使用 Makefile 快速构建：

```bash
# 构建所有平台
make build-mcp

# 构建特定平台
make build-mcp-linux
make build-mcp-darwin
make build-mcp-windows
```

## 环境变量

| 变量               | 说明              | 必需 |
| ------------------ | ----------------- | ---- |
| `ZT_CENTRAL_TOKEN` | Central API Token | 否   |

## 故障排除

### 权限错误

```bash
chmod +x zerotier-mcp-linux-amd64
```

### 验证完整性

```bash
sha256sum -c SHA256SUMS
```

### 测试连接

```bash
export ZT_CENTRAL_TOKEN="your_token"
./zerotier-mcp-linux-amd64
# 按 Ctrl+C 退出
```
