// Package mcp 提供 ZeroTier SDK 的 MCP 服务集成
package mcp

import (
	"context"
	"fmt"

	"github.com/fromsko/zerotier-sdk/central"
	"github.com/fromsko/zerotier-sdk/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Server ZeroTier MCP 服务
type Server struct {
	mcpServer     *server.MCPServer
	localClient   client.Client
	centralClient central.Client
}

// Option 服务配置选项
type Option func(*Server)

// WithLocalClient 设置本地客户端
func WithLocalClient(c client.Client) Option {
	return func(s *Server) {
		s.localClient = c
	}
}

// WithCentralClient 设置云端客户端
func WithCentralClient(c central.Client) Option {
	return func(s *Server) {
		s.centralClient = c
	}
}

// WithCentralToken 使用 Token 创建云端客户端
func WithCentralToken(token string) Option {
	return func(s *Server) {
		s.centralClient = central.New(token)
	}
}

// New 创建 MCP 服务
func New(name, version string, opts ...Option) *Server {
	s := &Server{}

	for _, opt := range opts {
		opt(s)
	}

	// 默认创建本地客户端
	if s.localClient == nil {
		s.localClient = client.New()
	}

	// 创建 MCP 服务器
	s.mcpServer = server.NewMCPServer(
		name,
		version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	// 注册工具
	s.registerTools()

	return s
}

// ServeStdio 启动 stdio 服务
func (s *Server) ServeStdio() error {
	return server.ServeStdio(s.mcpServer)
}

// MCPServer 返回底层 MCP 服务器（用于自定义扩展）
func (s *Server) MCPServer() *server.MCPServer {
	return s.mcpServer
}

// registerTools 注册所有工具
func (s *Server) registerTools() {
	// 本地 API 工具
	s.registerLocalTools()

	// 云端 API 工具（如果配置了）
	if s.centralClient != nil {
		s.registerCentralTools()
	}
}

// registerLocalTools 注册本地 API 工具
func (s *Server) registerLocalTools() {
	// 获取节点状态
	s.mcpServer.AddTool(
		mcp.NewTool("zt_status",
			mcp.WithDescription("获取本地 ZeroTier 节点状态"),
		),
		s.handleStatus,
	)

	// 列出网络
	s.mcpServer.AddTool(
		mcp.NewTool("zt_networks",
			mcp.WithDescription("列出已加入的 ZeroTier 网络"),
		),
		s.handleNetworks,
	)

	// 加入网络
	s.mcpServer.AddTool(
		mcp.NewTool("zt_join",
			mcp.WithDescription("加入 ZeroTier 网络"),
			mcp.WithString("network_id",
				mcp.Required(),
				mcp.Description("网络 ID（16位十六进制）"),
			),
		),
		s.handleJoin,
	)

	// 离开网络
	s.mcpServer.AddTool(
		mcp.NewTool("zt_leave",
			mcp.WithDescription("离开 ZeroTier 网络"),
			mcp.WithString("network_id",
				mcp.Required(),
				mcp.Description("网络 ID"),
			),
		),
		s.handleLeave,
	)

	// 列出 Peers
	s.mcpServer.AddTool(
		mcp.NewTool("zt_peers",
			mcp.WithDescription("列出所有 ZeroTier Peers"),
		),
		s.handlePeers,
	)
}

// registerCentralTools 注册云端 API 工具
func (s *Server) registerCentralTools() {
	// 列出云端网络
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_networks",
			mcp.WithDescription("列出云端 ZeroTier 网络"),
		),
		s.handleCentralNetworks,
	)

	// 获取网络成员
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_members",
			mcp.WithDescription("列出网络成员"),
			mcp.WithString("network_id",
				mcp.Required(),
				mcp.Description("网络 ID"),
			),
		),
		s.handleCentralMembers,
	)

	// 授权成员
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_authorize",
			mcp.WithDescription("授权网络成员"),
			mcp.WithString("network_id",
				mcp.Required(),
				mcp.Description("网络 ID"),
			),
			mcp.WithString("member_id",
				mcp.Required(),
				mcp.Description("成员 ID"),
			),
		),
		s.handleCentralAuthorize,
	)

	// 取消授权
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_deauthorize",
			mcp.WithDescription("取消成员授权"),
			mcp.WithString("network_id",
				mcp.Required(),
				mcp.Description("网络 ID"),
			),
			mcp.WithString("member_id",
				mcp.Required(),
				mcp.Description("成员 ID"),
			),
		),
		s.handleCentralDeauthorize,
	)
}

// ============================================
// 本地 API 处理器
// ============================================

func (s *Server) handleStatus(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	status, err := s.localClient.Status()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取状态失败: %v", err)), nil
	}

	result := fmt.Sprintf(`节点状态:
- 地址: %s
- 版本: %s
- 在线: %v
- TCP回退: %v`,
		status.Address,
		status.Version,
		status.Online,
		status.TCPFallbackActive,
	)

	return mcp.NewToolResultText(result), nil
}

func (s *Server) handleNetworks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networks, err := s.localClient.Networks().List()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取网络失败: %v", err)), nil
	}

	if len(networks) == 0 {
		return mcp.NewToolResultText("暂未加入任何网络"), nil
	}

	result := "已加入的网络:\n"
	for _, n := range networks {
		result += fmt.Sprintf("\n[%s] %s\n", n.ID, n.Name)
		result += fmt.Sprintf("  状态: %s\n", n.Status)
		result += fmt.Sprintf("  IP: %v\n", n.AssignedAddresses)
	}

	return mcp.NewToolResultText(result), nil
}

func (s *Server) handleJoin(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	network, err := s.localClient.Networks().Join(networkID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("加入网络失败: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("已加入网络: %s (%s)", network.ID, network.Name)), nil
}

func (s *Server) handleLeave(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := s.localClient.Networks().Leave(networkID); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("离开网络失败: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("已离开网络: %s", networkID)), nil
}

func (s *Server) handlePeers(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	peers, err := s.localClient.Peers().List()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取 Peers 失败: %v", err)), nil
	}

	if len(peers) == 0 {
		return mcp.NewToolResultText("暂无 Peers"), nil
	}

	result := "Peers:\n"
	for _, p := range peers {
		result += fmt.Sprintf("\n[%s]\n", p.Address)
		result += fmt.Sprintf("  角色: %s\n", p.Role)
		result += fmt.Sprintf("  版本: %s\n", p.Version)
		result += fmt.Sprintf("  延迟: %dms\n", p.Latency)
	}

	return mcp.NewToolResultText(result), nil
}

// ============================================
// 云端 API 处理器
// ============================================

func (s *Server) handleCentralNetworks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networks, err := s.centralClient.Networks().List()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取网络失败: %v", err)), nil
	}

	if len(networks) == 0 {
		return mcp.NewToolResultText("暂无网络"), nil
	}

	result := "云端网络:\n"
	for _, n := range networks {
		result += fmt.Sprintf("\n[%s] %s\n", n.ID, n.Config.Name)
		result += fmt.Sprintf("  在线: %d / 授权: %d / 总计: %d\n",
			n.OnlineMemberCount, n.AuthorizedMemberCount, n.TotalMemberCount)
	}

	return mcp.NewToolResultText(result), nil
}

func (s *Server) handleCentralMembers(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	members, err := s.centralClient.Networks().Members(networkID).List()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("获取成员失败: %v", err)), nil
	}

	if len(members) == 0 {
		return mcp.NewToolResultText("暂无成员"), nil
	}

	result := fmt.Sprintf("网络 %s 的成员:\n", networkID)
	for _, m := range members {
		status := "❌"
		if m.Config.Authorized {
			status = "✅"
		}
		result += fmt.Sprintf("\n%s [%s] %s\n", status, m.NodeID, m.Name)
		result += fmt.Sprintf("  IP: %v\n", m.Config.IPAssignments)
	}

	return mcp.NewToolResultText(result), nil
}

func (s *Server) handleCentralAuthorize(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	memberID, err := req.RequireString("member_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	member, err := s.centralClient.Networks().Members(networkID).Authorize(memberID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("授权失败: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("已授权成员: %s (%s)", member.NodeID, member.Name)), nil
}

func (s *Server) handleCentralDeauthorize(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	memberID, err := req.RequireString("member_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	member, err := s.centralClient.Networks().Members(networkID).Deauthorize(memberID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("取消授权失败: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("已取消授权: %s (%s)", member.NodeID, member.Name)), nil
}
