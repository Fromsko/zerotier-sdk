package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fromsko/zerotier-sdk"
	"github.com/mark3labs/mcp-go/mcp"
)

// registerLocalExtraTools 注册额外的本地 API 工具
func (s *Server) registerLocalExtraTools() {
	s.mcpServer.AddTool(
		mcp.NewTool("zt_network_info",
			mcp.WithDescription("获取已加入网络详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleNetworkInfo,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_update_network",
			mcp.WithDescription("更新本地网络设置"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithBoolean("allow_dns", mcp.Description("允许 DNS")),
			mcp.WithBoolean("allow_default", mcp.Description("允许默认路由")),
			mcp.WithBoolean("allow_global", mcp.Description("允许全局路由")),
			mcp.WithBoolean("allow_managed", mcp.Description("允许托管路由")),
		),
		s.handleUpdateNetwork,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_peer_info",
			mcp.WithDescription("获取指定 Peer 详情"),
			mcp.WithString("peer_id", mcp.Required(), mcp.Description("Peer ID")),
		),
		s.handlePeerInfo,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_controller_networks",
			mcp.WithDescription("列出控制器管理的所有网络")),
		s.handleControllerNetworks,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_controller_network_info",
			mcp.WithDescription("获取控制器网络详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleControllerNetworkInfo,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_controller_members",
			mcp.WithDescription("列出控制器网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleControllerMembers,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_controller_member_info",
			mcp.WithDescription("获取控制器网络成员详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员 ID")),
		),
		s.handleControllerMemberInfo,
	)
}

// registerCentralExtraTools 注册额外的云端 API 工具
func (s *Server) registerCentralExtraTools() {
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_network_info",
			mcp.WithDescription("获取云端网络详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleCentralNetworkInfo,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_create_network",
			mcp.WithDescription("创建云端网络"),
			mcp.WithString("name", mcp.Required(), mcp.Description("网络名称")),
			mcp.WithBoolean("private", mcp.Description("是否私有网络")),
			mcp.WithString("start_ip", mcp.Description("IPv4 分配池起始地址")),
			mcp.WithString("end_ip", mcp.Description("IPv4 分配池结束地址")),
		),
		s.handleCentralCreateNetwork,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_update_network",
			mcp.WithDescription("更新云端网络配置"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("name", mcp.Description("网络名称")),
			mcp.WithBoolean("private", mcp.Description("是否私有网络")),
			mcp.WithBoolean("enable_broadcast", mcp.Description("是否启用广播")),
		),
		s.handleCentralUpdateNetwork,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_delete_network",
			mcp.WithDescription("删除云端网络"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleCentralDeleteNetwork,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_member_info",
			mcp.WithDescription("获取云端网络成员详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员 ID")),
		),
		s.handleCentralMemberInfo,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_delete_member",
			mcp.WithDescription("删除云端网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员 ID")),
		),
		s.handleCentralDeleteMember,
	)

	// 特殊业务工具：授权并分配 IP（兼容旧版参数 ip_address + name）
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_authorize_with_ip",
			mcp.WithDescription("授权成员并分配自定义 IP 地址"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员 ID")),
			mcp.WithString("ip_address", mcp.Required(), mcp.Description("IP 地址（单个或多个用逗号分隔）")),
			mcp.WithString("name", mcp.Description("成员名称（可选）")),
		),
		s.handleCentralAuthorizeWithIP,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_invitations",
			mcp.WithDescription("列出组织邀请")),
		s.handleCentralInvitations,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_create_invitation",
			mcp.WithDescription("创建组织邀请"),
			mcp.WithString("email", mcp.Required(), mcp.Description("被邀请人邮箱")),
			mcp.WithString("org_id", mcp.Description("组织 ID（可选）")),
		),
		s.handleCentralCreateInvitation,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_accept_invitation",
			mcp.WithDescription("接受组织邀请"),
			mcp.WithString("invite_id", mcp.Required(), mcp.Description("邀请 ID")),
		),
		s.handleCentralAcceptInvitation,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_decline_invitation",
			mcp.WithDescription("拒绝/取消组织邀请"),
			mcp.WithString("invite_id", mcp.Required(), mcp.Description("邀请 ID")),
		),
		s.handleCentralDeclineInvitation,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_organization",
			mcp.WithDescription("获取组织信息"),
			mcp.WithString("org_id", mcp.Description("组织 ID（为空表示当前用户所属组织）")),
		),
		s.handleCentralOrganization,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_organization_members",
			mcp.WithDescription("获取组织成员"),
			mcp.WithString("org_id", mcp.Required(), mcp.Description("组织 ID")),
		),
		s.handleCentralOrganizationMembers,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_user",
			mcp.WithDescription("获取用户信息"),
			mcp.WithString("user_id", mcp.Required(), mcp.Description("用户 ID")),
		),
		s.handleCentralUser,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_random_token",
			mcp.WithDescription("获取 Central 随机 Token")),
		s.handleCentralRandomToken,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_create_token",
			mcp.WithDescription("为用户创建 API Token"),
			mcp.WithString("user_id", mcp.Required(), mcp.Description("用户 ID")),
			mcp.WithString("token_name", mcp.Required(), mcp.Description("Token 名称")),
			mcp.WithString("token_value", mcp.Required(), mcp.Description("Token 值（至少 32 位）")),
		),
		s.handleCentralCreateToken,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_delete_token",
			mcp.WithDescription("删除用户的 API Token"),
			mcp.WithString("user_id", mcp.Required(), mcp.Description("用户 ID")),
			mcp.WithString("token_name", mcp.Required(), mcp.Description("Token 名称")),
		),
		s.handleCentralDeleteToken,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_set_network_user_permissions",
			mcp.WithDescription("设置用户对网络的权限"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("user_id", mcp.Required(), mcp.Description("用户 ID")),
			mcp.WithBoolean("read", mcp.Description("读取权限")),
			mcp.WithBoolean("authorize", mcp.Description("授权权限")),
			mcp.WithBoolean("modify", mcp.Description("修改权限")),
			mcp.WithBoolean("delete", mcp.Description("删除权限")),
		),
		s.handleCentralSetNetworkUserPermissions,
	)
}

// jsonResult 将对象序列化为格式化的 JSON 返回
func jsonResult(v interface{}) (*mcp.CallToolResult, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal result: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// textResult 返回纯文本结果
func textResult(format string, args ...interface{}) *mcp.CallToolResult {
	return mcp.NewToolResultText(fmt.Sprintf(format, args...))
}

// errorResult 返回工具错误但不返回 Go error
func errorResult(format string, args ...interface{}) *mcp.CallToolResult {
	return mcp.NewToolResultError(fmt.Sprintf(format, args...))
}

// parseBool 解析可选布尔参数，未指定返回 nil
func parseBool(req mcp.CallToolRequest, key string) *bool {
	val, err := req.RequireBool(key)
	if err != nil {
		return nil
	}
	return &val
}

// parseIPList 解析逗号分隔的 IP 列表
func parseIPList(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

// ============================================
// 本地 API 扩展处理器
// ============================================

func (s *Server) handleNetworkInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	network, err := s.localClient.Networks().Get(networkID)
	if err != nil {
		return errorResult("获取网络失败: %v", err), nil
	}
	return jsonResult(network)
}

func (s *Server) handleUpdateNetwork(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	dns := req.GetBool("allow_dns", false)
	def := req.GetBool("allow_default", false)
	global := req.GetBool("allow_global", false)
	managed := req.GetBool("allow_managed", false)

	settings := zerotier.NewNetworkSettings().
		AllowDNS(dns).
		AllowDefault(def).
		AllowGlobal(global).
		AllowManaged(managed).
		Build()

	network, err := s.localClient.Networks().Update(networkID, settings)
	if err != nil {
		return errorResult("更新网络失败: %v", err), nil
	}
	return jsonResult(network)
}

func (s *Server) handlePeerInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	peerID, err := req.RequireString("peer_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	peer, err := s.localClient.Peers().Get(peerID)
	if err != nil {
		return errorResult("获取 Peer 失败: %v", err), nil
	}
	return jsonResult(peer)
}

func (s *Server) handleControllerNetworks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networks, err := s.localClient.Controller().ListNetworks()
	if err != nil {
		return errorResult("获取控制器网络失败: %v", err), nil
	}
	return jsonResult(networks)
}

func (s *Server) handleControllerNetworkInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	network, err := s.localClient.Controller().GetNetwork(networkID)
	if err != nil {
		return errorResult("获取控制器网络详情失败: %v", err), nil
	}
	return jsonResult(network)
}

func (s *Server) handleControllerMembers(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	members, err := s.localClient.Controller().ListMembers(networkID)
	if err != nil {
		return errorResult("获取控制器网络成员失败: %v", err), nil
	}
	return jsonResult(members)
}

func (s *Server) handleControllerMemberInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	member, err := s.localClient.Controller().GetMember(networkID, memberID)
	if err != nil {
		return errorResult("获取控制器成员详情失败: %v", err), nil
	}
	return jsonResult(member)
}

// ============================================
// 云端 API 扩展处理器
// ============================================

func (s *Server) handleCentralNetworkInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	network, err := s.centralClient.Networks().Get(networkID)
	if err != nil {
		return errorResult("获取网络失败: %v", err), nil
	}
	return jsonResult(network)
}

func (s *Server) handleCentralCreateNetwork(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := req.RequireString("name")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	cfg := zerotier.NewCentralNetworkConfig().Name(name)

	if v := parseBool(req, "private"); v != nil {
		cfg.Private(*v)
	}

	startIP := req.GetString("start_ip", "")
	endIP := req.GetString("end_ip", "")
	if startIP != "" && endIP != "" {
		cfg.AddIPPool(startIP, endIP)
		cfg.V4AssignMode(true)
	}

	network, err := s.centralClient.Networks().Create(cfg.Build())
	if err != nil {
		return errorResult("创建网络失败: %v", err), nil
	}
	return jsonResult(network)
}

func (s *Server) handleCentralUpdateNetwork(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	cfg := zerotier.NewCentralNetworkConfig()
	if name := req.GetString("name", ""); name != "" {
		cfg.Name(name)
	}
	if v := parseBool(req, "private"); v != nil {
		cfg.Private(*v)
	}
	if v := parseBool(req, "enable_broadcast"); v != nil {
		cfg.EnableBroadcast(*v)
	}

	network, err := s.centralClient.Networks().Update(networkID, cfg.Build())
	if err != nil {
		return errorResult("更新网络失败: %v", err), nil
	}
	return jsonResult(network)
}

func (s *Server) handleCentralDeleteNetwork(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	if err := s.centralClient.Networks().Delete(networkID); err != nil {
		return errorResult("删除网络失败: %v", err), nil
	}
	return textResult("已删除网络: %s", networkID), nil
}

func (s *Server) handleCentralMemberInfo(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	member, err := s.centralClient.Networks().Members(networkID).Get(memberID)
	if err != nil {
		return errorResult("获取成员失败: %v", err), nil
	}
	return jsonResult(member)
}

func (s *Server) handleCentralDeleteMember(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	if err := s.centralClient.Networks().Members(networkID).Delete(memberID); err != nil {
		return errorResult("删除成员失败: %v", err), nil
	}
	return textResult("已删除成员: %s", memberID), nil
}

// handleCentralAuthorizeWithIP 是业务组合工具：先授权，再分配 IP，并可设置名称
func (s *Server) handleCentralAuthorizeWithIP(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	ipStr, err := req.RequireString("ip_address")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	authorized := true
	updateReq := &zerotier.UpdateMemberRequest{
		Config: &zerotier.UpdateMemberConfig{
			Authorized:    &authorized,
			IPAssignments: parseIPList(ipStr),
		},
	}
	if name := req.GetString("name", ""); name != "" {
		updateReq.Name = name
	}

	member, err := s.centralClient.Networks().Members(networkID).Update(memberID, updateReq)
	if err != nil {
		return errorResult("授权并分配 IP 失败: %v", err), nil
	}
	return jsonResult(member)
}

func (s *Server) handleCentralInvitations(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	invites, err := s.centralClient.Invitations().List()
	if err != nil {
		return errorResult("获取邀请列表失败: %v", err), nil
	}
	return jsonResult(invites)
}

func (s *Server) handleCentralCreateInvitation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	email, err := req.RequireString("email")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	inviteReq := &zerotier.OrganizationInvitation{Email: email}
	if orgID := req.GetString("org_id", ""); orgID != "" {
		inviteReq.OrgID = orgID
	}

	invite, err := s.centralClient.Invitations().Invite(inviteReq)
	if err != nil {
		return errorResult("创建邀请失败: %v", err), nil
	}
	return jsonResult(invite)
}

func (s *Server) handleCentralAcceptInvitation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	inviteID, err := req.RequireString("invite_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	invite, err := s.centralClient.Invitations().Accept(inviteID)
	if err != nil {
		return errorResult("接受邀请失败: %v", err), nil
	}
	return jsonResult(invite)
}

func (s *Server) handleCentralDeclineInvitation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	inviteID, err := req.RequireString("invite_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	if err := s.centralClient.Invitations().Decline(inviteID); err != nil {
		return errorResult("拒绝邀请失败: %v", err), nil
	}
	return textResult("已拒绝邀请: %s", inviteID), nil
}

func (s *Server) handleCentralOrganization(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgID := req.GetString("org_id", "")
	var org *zerotier.Organization
	var err error
	if orgID == "" {
		org, err = s.centralClient.Organization().Get()
	} else {
		org, err = s.centralClient.Organization().GetByID(orgID)
	}
	if err != nil {
		return errorResult("获取组织失败: %v", err), nil
	}
	return jsonResult(org)
}

func (s *Server) handleCentralOrganizationMembers(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgID, err := req.RequireString("org_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	members, err := s.centralClient.Organization().Members(orgID)
	if err != nil {
		return errorResult("获取组织成员失败: %v", err), nil
	}
	return jsonResult(members)
}

func (s *Server) handleCentralUser(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	userID, err := req.RequireString("user_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	user, err := s.centralClient.Users().Get(userID)
	if err != nil {
		return errorResult("获取用户失败: %v", err), nil
	}
	return jsonResult(user)
}

func (s *Server) handleCentralRandomToken(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	token, err := s.centralClient.RandomToken()
	if err != nil {
		return errorResult("获取随机 Token 失败: %v", err), nil
	}
	return jsonResult(token)
}

func (s *Server) handleCentralCreateToken(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	userID, err := req.RequireString("user_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	tokenName, err := req.RequireString("token_name")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	tokenValue, err := req.RequireString("token_value")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	token, err := s.centralClient.Users().AddToken(userID, &zerotier.APIToken{
		TokenName: tokenName,
		Token:     tokenValue,
	})
	if err != nil {
		return errorResult("创建 Token 失败: %v", err), nil
	}
	return jsonResult(token)
}

func (s *Server) handleCentralDeleteToken(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	userID, err := req.RequireString("user_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	tokenName, err := req.RequireString("token_name")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	if err := s.centralClient.Users().DeleteToken(userID, tokenName); err != nil {
		return errorResult("删除 Token 失败: %v", err), nil
	}
	return textResult("已删除 Token: %s", tokenName), nil
}

func (s *Server) handleCentralSetNetworkUserPermissions(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	userID, err := req.RequireString("user_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	permission := &zerotier.NetworkUserPermissions{
		ID:        userID,
		Read:      req.GetBool("read", false),
		Authorize: req.GetBool("authorize", false),
		Modify:    req.GetBool("modify", false),
		Delete:    req.GetBool("delete", false),
	}

	result, err := s.centralClient.Networks().Permissions(networkID).SetUserPermissions(permission)
	if err != nil {
		return errorResult("设置权限失败: %v", err), nil
	}
	return jsonResult(result)
}
