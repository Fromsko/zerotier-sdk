package mcp

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fromsko/zerotier-sdk"
	"github.com/mark3labs/mcp-go/mcp"
)

// registerBatchTools 注册批量操作工具
func (s *Server) registerBatchTools() {
	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_batch_authorize",
			mcp.WithDescription("批量授权网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_ids", mcp.Required(), mcp.Description("成员 ID，多个用逗号分隔")),
		),
		s.handleBatchAuthorize,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_batch_deauthorize",
			mcp.WithDescription("批量取消授权网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_ids", mcp.Required(), mcp.Description("成员 ID，多个用逗号分隔")),
		),
		s.handleBatchDeauthorize,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_batch_delete",
			mcp.WithDescription("批量删除网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_ids", mcp.Required(), mcp.Description("成员 ID，多个用逗号分隔")),
		),
		s.handleBatchDelete,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_batch_rename",
			mcp.WithDescription("批量重命名网络成员"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_ids", mcp.Required(), mcp.Description("成员 ID，多个用逗号分隔")),
			mcp.WithString("name_pattern", mcp.Required(), mcp.Description("名称模板，支持 {index} / {node_id} 占位")),
		),
		s.handleBatchRename,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_batch_set_ip",
			mcp.WithDescription("批量设置成员 IP"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("member_ips", mcp.Required(), mcp.Description("成员 IP 映射，格式：member_id:ip，多个用逗号分隔")),
		),
		s.handleBatchSetIP,
	)

	s.mcpServer.AddTool(
		mcp.NewTool("zt_central_members_rank",
			mcp.WithDescription("网络成员排行榜/排序筛选"),
			mcp.WithString("network_id", mcp.Required(), mcp.Description("网络 ID")),
			mcp.WithString("sort_by", mcp.Description("排序字段：name(默认)/online/ip/creation/last_seen")),
			mcp.WithString("order", mcp.Description("排序方向：asc(默认)/desc")),
			mcp.WithNumber("limit", mcp.Description("返回数量上限，0 表示全部")),
			mcp.WithNumber("offset", mcp.Description("偏移量")),
			mcp.WithString("filter", mcp.Description("过滤：all(默认)/authorized/unauthorized")),
		),
		s.handleMembersRank,
	)
}

// parseMemberIDs 解析逗号分隔的成员 ID 列表
func parseMemberIDs(s string) []string {
	parts := strings.Split(s, ",")
	ids := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			ids = append(ids, p)
		}
	}
	return ids
}

// handleBatchAuthorize 批量授权
func (s *Server) handleBatchAuthorize(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberIDs := parseMemberIDs(req.GetString("member_ids", ""))

	svc := s.centralClient.Networks().Members(networkID)
	var success []string
	var failed []string
	for _, id := range memberIDs {
		if _, err := svc.Authorize(id); err != nil {
			failed = append(failed, fmt.Sprintf("%s: %v", id, err))
		} else {
			success = append(success, id)
		}
	}

	return jsonResult(map[string]any{
		"success": success,
		"failed":  failed,
		"total":   len(memberIDs),
	})
}

// handleBatchDeauthorize 批量取消授权
func (s *Server) handleBatchDeauthorize(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberIDs := parseMemberIDs(req.GetString("member_ids", ""))

	svc := s.centralClient.Networks().Members(networkID)
	var success []string
	var failed []string
	for _, id := range memberIDs {
		if _, err := svc.Deauthorize(id); err != nil {
			failed = append(failed, fmt.Sprintf("%s: %v", id, err))
		} else {
			success = append(success, id)
		}
	}

	return jsonResult(map[string]any{
		"success": success,
		"failed":  failed,
		"total":   len(memberIDs),
	})
}

// handleBatchDelete 批量删除
func (s *Server) handleBatchDelete(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberIDs := parseMemberIDs(req.GetString("member_ids", ""))

	svc := s.centralClient.Networks().Members(networkID)
	var success []string
	var failed []string
	for _, id := range memberIDs {
		if err := svc.Delete(id); err != nil {
			failed = append(failed, fmt.Sprintf("%s: %v", id, err))
		} else {
			success = append(success, id)
		}
	}

	return jsonResult(map[string]any{
		"success": success,
		"failed":  failed,
		"total":   len(memberIDs),
	})
}

// handleBatchRename 批量重命名
func (s *Server) handleBatchRename(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberIDs := parseMemberIDs(req.GetString("member_ids", ""))
	pattern, err := req.RequireString("name_pattern")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	svc := s.centralClient.Networks().Members(networkID)
	var success []string
	var failed []string
	for i, id := range memberIDs {
		name := strings.ReplaceAll(pattern, "{index}", strconv.Itoa(i+1))
		name = strings.ReplaceAll(name, "{node_id}", id)
		if _, err := svc.SetName(id, name); err != nil {
			failed = append(failed, fmt.Sprintf("%s: %v", id, err))
		} else {
			success = append(success, fmt.Sprintf("%s -> %s", id, name))
		}
	}

	return jsonResult(map[string]any{
		"success": success,
		"failed":  failed,
		"total":   len(memberIDs),
	})
}

// handleBatchSetIP 批量设置 IP
func (s *Server) handleBatchSetIP(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}
	memberIPs, err := req.RequireString("member_ips")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	svc := s.centralClient.Networks().Members(networkID)
	type result struct {
		MemberID string `json:"member_id"`
		IP       string `json:"ip"`
		Error    string `json:"error,omitempty"`
	}
	var results []result

	for _, part := range strings.Split(memberIPs, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		sep := ":"
		if strings.Contains(part, "=") {
			sep = "="
		}
		kv := strings.SplitN(part, sep, 2)
		if len(kv) != 2 {
			results = append(results, result{MemberID: part, Error: "映射格式错误"})
			continue
		}
		memberID := strings.TrimSpace(kv[0])
		ip := strings.TrimSpace(kv[1])
		if _, err := svc.SetIPAssignments(memberID, []string{ip}); err != nil {
			results = append(results, result{MemberID: memberID, IP: ip, Error: err.Error()})
		} else {
			results = append(results, result{MemberID: memberID, IP: ip})
		}
	}

	return jsonResult(results)
}

// isMemberOnline 根据 lastSeen 判断成员是否在线（6 分钟内有心跳视为在线）
func isMemberOnline(m *zerotier.Member) bool {
	if m == nil || m.LastSeen == 0 {
		return false
	}
	return time.Since(time.Unix(m.LastSeen/1000, 0)) < 6*time.Minute
}

// handleMembersRank 成员排行榜
func (s *Server) handleMembersRank(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	networkID, err := req.RequireString("network_id")
	if err != nil {
		return errorResult("%s", err.Error()), nil
	}

	members, err := s.centralClient.Networks().Members(networkID).List()
	if err != nil {
		return errorResult("获取成员失败: %v", err), nil
	}

	sortBy := req.GetString("sort_by", "name")
	order := req.GetString("order", "asc")
	filter := req.GetString("filter", "all")
	limit := req.GetInt("limit", 0)
	offset := req.GetInt("offset", 0)

	// 过滤
	filtered := make([]zerotier.Member, 0, len(members))
	for _, m := range members {
		switch filter {
		case "authorized":
			if m.Config == nil || !m.Config.Authorized {
				continue
			}
		case "unauthorized":
			if m.Config != nil && m.Config.Authorized {
				continue
			}
		}
		filtered = append(filtered, m)
	}

	// 排序
	sort.SliceStable(filtered, func(i, j int) bool {
		var less bool
		switch sortBy {
		case "online":
			less = isMemberOnline(&filtered[i]) && !isMemberOnline(&filtered[j])
		case "ip":
			less = firstIP(&filtered[i]) < firstIP(&filtered[j])
		case "creation":
			ci, cj := int64(0), int64(0)
			if filtered[i].Config != nil {
				ci = filtered[i].Config.CreationTime
			}
			if filtered[j].Config != nil {
				cj = filtered[j].Config.CreationTime
			}
			less = ci < cj
		case "last_seen":
			less = filtered[i].LastSeen < filtered[j].LastSeen
		case "name":
			fallthrough
		default:
			less = strings.ToLower(filtered[i].Name) < strings.ToLower(filtered[j].Name)
		}
		if order == "desc" {
			return !less
		}
		return less
	})

	// 分页
	if offset > len(filtered) {
		offset = len(filtered)
	}
	end := len(filtered)
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}
	page := filtered[offset:end]

	// 构造输出
	type item struct {
		Index      int    `json:"index"`
		NodeID     string `json:"node_id"`
		Name       string `json:"name"`
		Online     bool   `json:"online"`
		Authorized bool   `json:"authorized"`
		IP         string `json:"ip,omitempty"`
		LastSeen   int64  `json:"last_seen"`
		Version    string `json:"version,omitempty"`
		PhysicalIP string `json:"physical_ip,omitempty"`
	}
	items := make([]item, 0, len(page))
	for i, m := range page {
		authorized := false
		if m.Config != nil {
			authorized = m.Config.Authorized
		}
		items = append(items, item{
			Index:      i + 1,
			NodeID:     m.NodeID,
			Name:       m.Name,
			Online:     isMemberOnline(&m),
			Authorized: authorized,
			IP:         firstIP(&m),
			LastSeen:   m.LastSeen,
			Version:    m.ClientVersion,
			PhysicalIP: m.PhysicalAddress,
		})
	}

	return jsonResult(map[string]any{
		"total":       len(filtered),
		"returned":    len(items),
		"sort_by":     sortBy,
		"order":       order,
		"filter":      filter,
		"leaderboard": items,
	})
}

// firstIP 返回成员首个 IP 地址
func firstIP(m *zerotier.Member) string {
	if m == nil || m.Config == nil {
		return ""
	}
	if len(m.Config.IPAssignments) > 0 {
		return m.Config.IPAssignments[0]
	}
	return ""
}
