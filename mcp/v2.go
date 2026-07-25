package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"

	centralv2 "github.com/fromsko/zerotier-sdk/central/v2"
	"github.com/mark3labs/mcp-go/mcp"
)

// v2Response 响应基础接口（所有生成的 V2 Response 均实现该接口）
type v2Response interface {
	StatusCode() int
	Status() string
	GetBody() []byte
}

// hasArg 判断请求参数中是否包含指定 key
func hasArg(req mcp.CallToolRequest, key string) bool {
	_, ok := req.GetArguments()[key]
	return ok
}

// v2Result 统一处理 Central V2 响应：2xx 优先返回 JSON200，否则返回格式化后的 Body
func v2Result(resp v2Response) (*mcp.CallToolResult, error) {
	statusCode := resp.StatusCode()
	if statusCode < 200 || statusCode >= 300 {
		return errorResult("HTTP %d: %s", statusCode, string(resp.GetBody())), nil
	}

	// 尝试取 JSON200 字段
	rv := reflect.ValueOf(resp)
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}
	if jsonField := rv.FieldByName("JSON200"); jsonField.IsValid() && !jsonField.IsNil() {
		return jsonResult(jsonField.Interface())
	}

	// 退化为格式化 JSON body
	body := resp.GetBody()
	var out bytes.Buffer
	if err := json.Indent(&out, body, "", "  "); err == nil {
		return mcp.NewToolResultText(out.String()), nil
	}
	return mcp.NewToolResultText(string(body)), nil
}

// registerCentralV2Tools 注册 Central V2 API 工具
func (s *Server) registerCentralV2Tools() {
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_orgs",
			mcp.WithDescription("列出 Central V2 组织")),
		s.handleV2Orgs,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_network_groups",
			mcp.WithDescription("列出 Central V2 网络组"),
			mcp.WithString("org_id", mcp.Description("按组织 ID 过滤")),
			mcp.WithBoolean("stats", mcp.Description("包含统计信息")),
		),
		s.handleV2NetworkGroups,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_network_group",
			mcp.WithDescription("获取 Central V2 网络组详情"),
			mcp.WithString("network_group_id", mcp.Required(), mcp.Description("网络组 ID")),
		),
		s.handleV2NetworkGroup,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_networks",
			mcp.WithDescription("列出 Central V2 网络"),
			mcp.WithString("org_id", mcp.Description("按组织 ID 过滤")),
			mcp.WithBoolean("stats", mcp.Description("包含统计信息")),
		),
		s.handleV2Networks,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_network",
			mcp.WithDescription("获取 Central V2 网络详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleV2Network,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_network_members",
			mcp.WithDescription("列出 Central V2 网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleV2NetworkMembers,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_network_member",
			mcp.WithDescription("获取 Central V2 网络成员详情"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员/设备 ID")),
		),
		s.handleV2NetworkMember,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_authorize_member",
			mcp.WithDescription("授权 Central V2 网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员/设备 ID")),
		),
		s.handleV2AuthorizeMember,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_deauthorize_member",
			mcp.WithDescription("取消授权 Central V2 网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_id", mcp.Required(), mcp.Description("成员/设备 ID")),
		),
		s.handleV2DeauthorizeMember,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_create_network",
			mcp.WithDescription("在指定网络组中创建 Central V2 网络"),
			mcp.WithString("network_group_id", mcp.Required(), mcp.Description("网络组 ID")),
			mcp.WithString("name", mcp.Required(), mcp.Description("网络名称")),
			mcp.WithString("description", mcp.Description("网络描述")),
			mcp.WithString("v4_subnet", mcp.Description("IPv4 子网 CIDR，例如 192.168.192.0/24")),
			mcp.WithString("start_ip", mcp.Description("IPv4 分配池起始地址")),
			mcp.WithString("end_ip", mcp.Description("IPv4 分配池结束地址")),
			mcp.WithBoolean("private", mcp.Description("是否为私有网络")),
			mcp.WithBoolean("enable_broadcast", mcp.Description("是否启用广播")),
			mcp.WithNumber("mtu", mcp.Description("MTU")),
			mcp.WithNumber("multicast_limit", mcp.Description("组播限制")),
		),
		s.handleV2CreateNetwork,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_delete_network",
			mcp.WithDescription("删除 Central V2 网络"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
		),
		s.handleV2DeleteNetwork,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_v2_create_network_group",
			mcp.WithDescription("在指定组织中创建 Central V2 网络组"),
			mcp.WithString("org_id", mcp.Required(), mcp.Description("组织 ID")),
			mcp.WithString("name", mcp.Required(), mcp.Description("网络组名称")),
			mcp.WithString("description", mcp.Description("网络组描述")),
		),
		s.handleV2CreateNetworkGroup,
	)
}

// handleV2Orgs 列出组织
func (s *Server) handleV2Orgs(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := s.centralV2Client.ListOrgsWithResponse(ctx, nil)
	if err != nil {
		return errorResult("列出组织失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2NetworkGroups 列出网络组
func (s *Server) handleV2NetworkGroups(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := &centralv2.ListNetworkGroupsParams{}
	if orgID := req.GetString("org_id", ""); orgID != "" {
		params.OrgId = &orgID
	}
	if hasArg(req, "stats") {
		stats := req.GetBool("stats", false)
		params.Stats = &stats
	}

	resp, err := s.centralV2Client.ListNetworkGroupsWithResponse(ctx, params)
	if err != nil {
		return errorResult("列出网络组失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2NetworkGroup 获取网络组
func (s *Server) handleV2NetworkGroup(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("network_group_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.GetNetworkGroupWithResponse(ctx, id)
	if err != nil {
		return errorResult("获取网络组失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2Networks 列出网络
func (s *Server) handleV2Networks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	params := &centralv2.ListNetworksParams{}
	if orgID := req.GetString("org_id", ""); orgID != "" {
		params.OrgId = &orgID
	}
	if req.GetBool("stats", false) {
		stats := true
		params.Stats = &stats
	}

	resp, err := s.centralV2Client.ListNetworksWithResponse(ctx, params)
	if err != nil {
		return errorResult("列出网络失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2Network 获取网络
func (s *Server) handleV2Network(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.GetNetworkWithResponse(ctx, id)
	if err != nil {
		return errorResult("获取网络失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2NetworkMembers 列出网络成员
func (s *Server) handleV2NetworkMembers(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.ListNetworkMembersWithResponse(ctx, id)
	if err != nil {
		return errorResult("列出成员失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2NetworkMember 获取网络成员
func (s *Server) handleV2NetworkMember(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.GetNetworkMemberWithResponse(ctx, networkID, memberID)
	if err != nil {
		return errorResult("获取成员失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2AuthorizeMember 授权成员
func (s *Server) handleV2AuthorizeMember(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.AuthorizeNetworkMemberWithResponse(ctx, networkID, memberID)
	if err != nil {
		return errorResult("授权成员失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2DeauthorizeMember 取消授权成员
func (s *Server) handleV2DeauthorizeMember(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberID, err := req.RequireString("member_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.DeauthorizeNetworkMemberWithResponse(ctx, networkID, memberID)
	if err != nil {
		return errorResult("取消授权失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2CreateNetwork 创建网络
func (s *Server) handleV2CreateNetwork(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkGroupID, err := req.RequireString("network_group_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	name, err := req.RequireString("name")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	body := centralv2.CreateNetworkRequest{Name: name}
	if desc := req.GetString("description", ""); desc != "" {
		body.Description = &desc
	}

	cfg := &centralv2.NetworkConfigRequest{}
	if v := req.GetString("v4_subnet", ""); v != "" {
		cfg.V4Subnet = &v
	}
	if startIP := req.GetString("start_ip", ""); startIP != "" {
		if endIP := req.GetString("end_ip", ""); endIP != "" {
			cfg.V4IpAssignmentPools = &[]centralv2.IPRange{
				{IpRangeStart: startIP, IpRangeEnd: endIP},
			}
			cfg.V4AssignmentMode = &centralv2.IPV4AssignMode{Zt: true}
		}
	}
	if v := parseBool(req, "private"); v != nil {
		cfg.Private = v
	}
	if v := parseBool(req, "enable_broadcast"); v != nil {
		cfg.EnableBroadcast = v
	}
	if hasArg(req, "mtu") {
		mtu := req.GetInt("mtu", 2800)
		cfg.Mtu = &mtu
	}
	if hasArg(req, "multicast_limit") {
		limit := req.GetInt("multicast_limit", 32)
		cfg.MulticastLimit = &limit
	}

	if cfg.V4Subnet != nil || cfg.V4IpAssignmentPools != nil || cfg.Private != nil || cfg.EnableBroadcast != nil || cfg.Mtu != nil || cfg.MulticastLimit != nil {
		body.Config = cfg
	}

	resp, err := s.centralV2Client.CreateNetworkWithResponse(ctx, networkGroupID, body)
	if err != nil {
		return errorResult("创建网络失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2DeleteNetwork 删除网络
func (s *Server) handleV2DeleteNetwork(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	id, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	resp, err := s.centralV2Client.DeleteNetworkWithResponse(ctx, id)
	if err != nil {
		return errorResult("删除网络失败: %v", err), nil
	}
	return v2Result(resp)
}

// handleV2CreateNetworkGroup 创建网络组
func (s *Server) handleV2CreateNetworkGroup(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	orgID, err := req.RequireString("org_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	name, err := req.RequireString("name")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	body := centralv2.CreateNetworkGroupRequest{Name: name}
	if desc := req.GetString("description", ""); desc != "" {
		body.Description = &desc
	}

	resp, err := s.centralV2Client.CreateNetworkGroupWithResponse(ctx, orgID, body)
	if err != nil {
		return errorResult("创建网络组失败: %v", err), nil
	}
	return v2Result(resp)
}
