# ZeroTier Go SDK

ZeroTier API 的 Go SDK。

## 安装

```bash
go get github.com/fromsko/zerotier-sdk
```

## 模块

| 模块      | 用途         | 地址             |
| --------- | ------------ | ---------------- |
| `client`  | 本地节点管理 | localhost:9993   |
| `central` | 云端控制面板 | api.zerotier.com |

## 快速开始

```go
import zerotier "github.com/fromsko/zerotier-sdk"

// 本地节点
local := zerotier.NewClient()
status, _ := local.Status()
networks, _ := local.Networks().List()

// 云端管理
cloud := zerotier.NewCentral("your_api_token")
networks, _ := cloud.Networks().List()
cloud.Networks().Members("nwid").Authorize("member_id")
```

## Client（本地 API）

```go
c := zerotier.NewClient()

// 节点状态
status, _ := c.Status()

// 网络
c.Networks().List()
c.Networks().Join("network_id")
c.Networks().Leave("network_id")

// Peers
c.Peers().List()

// 控制器（自托管）
c.Controller().CreateNetwork(nodeID, config)
```

详见 [client/README.md](client/README.md)

## Central（云端 API）

```go
c := zerotier.NewCentral("token")

// 网络
c.Networks().List()
c.Networks().Create(config)
c.Networks().Delete("network_id")

// 成员
c.Networks().Members("nwid").List()
c.Networks().Members("nwid").Authorize("member_id")
```

详见 [central/README.md](central/README.md)

## Builder 模式

```go
// 本地网络设置
zerotier.NewNetworkSettings().
    AllowDNS(true).
    AllowManaged(true).
    Build()

// 云端网络配置
zerotier.NewCentralNetworkConfig().
    Name("My Network").
    Private(true).
    AddRoute("10.0.0.0/24", nil).
    AddIPPool("10.0.0.1", "10.0.0.254").
    V4AssignMode(true).
    Build()

// 云端成员配置
zerotier.NewCentralMemberConfig().
    Name("my-device").
    Authorized(true).
    Build()
```

## 项目结构

```
zerotier-sdk/
├── zerotier.go      # 统一接口
├── client/          # 本地 Service API
├── central/         # 云端 Central API
├── mcp/             # MCP 服务集成
├── cmd/mcp/         # MCP 服务入口
└── example/         # 使用示例
```

## MCP 集成

支持与 Claude、Cursor 等 AI 工具集成：
详见 [mcp/README.md](mcp/README.md)

```go
package main

import (
	"log"
	"os"

	"github.com/fromsko/zerotier-sdk"
	ztmcp "github.com/fromsko/zerotier-sdk/mcp"
)

func main() {
	// 从环境变量获取 Central API Token
	token := os.Getenv("ZT_CENTRAL_TOKEN")
	// 云端连接
	localClientToken := os.Getenv("ZT_LOCAL_TOKEN")
	// 如果需要采用密钥形式连接
	localClient := zerotier.NewClient(
		zerotier.WithClientBaseURL("http://localhost:9993"),
		zerotier.WithClientToken(localClientToken),
	)

	var opts []ztmcp.Option
	if token != "" {
		opts = append(opts, ztmcp.WithCentralToken(token), ztmcp.WithLocalClient(localClient))
	}

	// 创建 MCP 服务
	s := ztmcp.New("zerotier", "1.0.0", opts...)

	// 启动 stdio 服务
	if err := s.ServeStdio(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
```

![Mcp-tool-client](./docs/mcp-client-info.png)

## License

MIT
