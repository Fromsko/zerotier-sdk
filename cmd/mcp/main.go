// ZeroTier MCP 服务入口
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
	var opts []ztmcp.Option

	// 本地连接
	if localClientToken != "" {
		// 如果需要采用密钥形式连接
		localClient := zerotier.NewClient(
			zerotier.WithClientBaseURL("http://localhost:9993"),
			zerotier.WithClientToken(localClientToken),
		)
		opts = append(opts, ztmcp.WithLocalClient(localClient))
	}

	if token != "" {
		opts = append(opts, ztmcp.WithCentralToken(token))
	}

	// Central V2 连接
	if v2Token := os.Getenv("ZT_CENTRAL_V2_TOKEN"); v2Token != "" {
		opts = append(opts, ztmcp.WithCentralV2Token(v2Token))
	}

	// 创建 MCP 服务
	s := ztmcp.New("zerotier", "1.0.0", opts...)

	// 启动 stdio 服务
	if err := s.ServeStdio(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
