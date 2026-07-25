// Package zerotier 提供 ZeroTier API 的 Go SDK
//
// 本 SDK 包含三个子模块：
//   - client: 本地 Service API（localhost:9993）
//   - central: 云端 Central API（api.zerotier.com / api/v1）
//   - central/v2: 新版 Central API（central.zerotier.com / api/v2）
//   - mcp: MCP 服务集成（stdio）
//
// 快速开始：
//
//	// 本地节点管理
//	local := zerotier.NewClient()
//	status, _ := local.Status()
//
//	// 云端管理（v1）
//	cloud := zerotier.NewCentral("your_api_token")
//	networks, _ := cloud.Networks().List()
//
//	// 云端管理（v2）
//	cloudV2, _ := zerotier.NewCentralV2("your_service_account_key")
//	resp, _ := cloudV2.ListNetworksWithResponse(context.Background(), nil)
package zerotier

import (
	"net/http"

	"github.com/fromsko/zerotier-sdk/central"
	centralv2 "github.com/fromsko/zerotier-sdk/central/v2"
	"github.com/fromsko/zerotier-sdk/client"
)

// ============================================
// 类型别名导出（方便外部使用）
// ============================================

// Client 本地 Service API 客户端
type Client = client.Client

// Central 云端 Central API 客户端
type Central = central.Client

// ============================================
// Client 相关类型
// ============================================

type (
	// ClientOption 本地客户端配置选项
	ClientOption = client.Option

	// NodeStatus 本地节点状态
	NodeStatus = client.NodeStatus

	// LocalNetwork 本地网络信息
	LocalNetwork = client.Network

	// NetworkSettings 本地网络设置
	NetworkSettings = client.NetworkSettings

	// Peer 节点信息
	Peer = client.Peer

	// PeerPath 节点路径
	PeerPath = client.PeerPath

	// ControllerStatus 控制器状态
	ControllerStatus = client.ControllerStatus

	// ControllerNetwork 控制器网络
	ControllerNetwork = client.ControllerNetwork

	// ControllerMember 控制器成员
	ControllerMember = client.ControllerMember

	// ControllerNetworkConfig 控制器网络配置
	ControllerNetworkConfig = client.ControllerNetworkConfig

	// ControllerMemberConfig 控制器成员配置
	ControllerMemberConfig = client.ControllerMemberConfig
)

// ============================================
// Central 相关类型
// ============================================

type (
	// CentralOption 云端客户端配置选项
	CentralOption = central.Option

	// CentralStatus 云端状态
	CentralStatus = central.CentralStatus

	// StatusUser 状态中的用户信息
	StatusUser = central.StatusUser

	// User 云端用户信息
	User = central.User

	// Network 云端网络信息
	Network = central.Network

	// NetworkConfig 网络配置
	NetworkConfig = central.NetworkConfig

	// Member 网络成员
	Member = central.Member

	// MemberConfig 成员配置
	MemberConfig = central.MemberConfig

	// CreateNetworkConfig 创建网络配置
	CreateNetworkConfig = central.CreateNetworkConfig

	// UpdateMemberRequest 更新成员请求
	UpdateMemberRequest = central.UpdateMemberRequest

	// Route 路由配置
	Route = central.Route

	// IPAssignmentPool IP 分配池
	IPAssignmentPool = central.IPAssignmentPool

	// AssignMode IP 分配模式
	AssignMode = central.AssignMode

	// DNS 配置
	DNS = central.DNS

	// Organization 组织信息
	Organization = central.Organization

	// OrganizationMember 组织成员
	OrganizationMember = central.OrganizationMember

	// OrganizationInvitation 组织邀请
	OrganizationInvitation = central.OrganizationInvitation

	// InviteStatus 邀请状态
	InviteStatus = central.InviteStatus

	// APIToken API Token
	APIToken = central.APIToken

	// RandomToken 随机 Token
	RandomToken = central.RandomToken

	// NetworkUserPermissions 网络用户权限
	NetworkUserPermissions = central.NetworkUserPermissions

	// Permissions 通用权限
	Permissions = central.Permissions

	// PermissionsMap 权限映射
	PermissionsMap = central.PermissionsMap

	// AuthMethods 认证方式
	AuthMethods = central.AuthMethods

	// OrgSSOConfig 组织 SSO 配置
	OrgSSOConfig = central.OrgSSOConfig

	// NetworkSSOConfig 网络 SSO 配置
	NetworkSSOConfig = central.NetworkSSOConfig

	// SsoIssuer SSO Issuer 配置
	SsoIssuer = central.SsoIssuer

	// OrganizationService 组织服务接口
	OrganizationService = central.OrganizationService

	// InvitationService 邀请服务接口
	InvitationService = central.InvitationService

	// UserService 用户服务接口
	UserService = central.UserService

	// NetworkPermissionService 网络权限服务接口
	NetworkPermissionService = central.NetworkPermissionService
)

// ============================================
// 构造函数
// ============================================

// NewClient 创建本地 Service API 客户端
//
// 默认连接 localhost:9993，自动读取系统 authtoken.secret
//
//	client := zerotier.NewClient()
//	client := zerotier.NewClient(zerotier.WithClientToken("token"))
func NewClient(opts ...ClientOption) Client {
	return client.New(opts...)
}

// NewCentral 创建云端 Central API 客户端
//
// 需要提供 API Token（从 my.zerotier.com 获取）
//
//	central := zerotier.NewCentral("your_api_token")
func NewCentral(token string, opts ...CentralOption) Central {
	return central.New(token, opts...)
}

// ============================================
// Client 选项函数
// ============================================

// WithClientBaseURL 设置本地 API 地址
func WithClientBaseURL(url string) ClientOption {
	return client.WithBaseURL(url)
}

// WithClientToken 设置本地认证 Token
func WithClientToken(token string) ClientOption {
	return client.WithToken(token)
}

// WithClientTokenFile 从文件读取本地 Token
func WithClientTokenFile(path string) ClientOption {
	return client.WithTokenFile(path)
}

// ============================================
// Central 选项函数
// ============================================

// WithCentralBaseURL 设置云端 API 地址
func WithCentralBaseURL(url string) CentralOption {
	return central.WithBaseURL(url)
}

// ============================================
// Central V2 相关类型
// ============================================

type (
	// CentralV2 Central API V2 客户端
	CentralV2 = centralv2.ClientWithResponses

	// CentralV2Option Central V2 客户端配置选项
	CentralV2Option = centralv2.ClientOption
)

// NewCentralV2 创建 Central API V2 客户端
//
// token 为 Central Service Account API Key，使用 Bearer 方式认证
//
//	centralV2, _ := zerotier.NewCentralV2("your_service_account_key")
func NewCentralV2(token string, opts ...CentralV2Option) (*CentralV2, error) {
	return centralv2.NewClientWithToken(token, opts...)
}

// WithCentralV2BaseURL 设置 Central V2 API 地址
func WithCentralV2BaseURL(url string) CentralV2Option {
	return centralv2.WithBaseURL(url)
}

// WithCentralV2HTTPClient 设置 Central V2 HTTP 客户端
func WithCentralV2HTTPClient(hc *http.Client) CentralV2Option {
	return centralv2.WithHTTPClient(hc)
}

// ============================================
// Builder 函数
// ============================================

// NewNetworkSettings 创建本地网络设置构建器
func NewNetworkSettings() *client.NetworkSettingsBuilder {
	return client.NewNetworkSettings()
}

// NewControllerNetworkConfig 创建控制器网络配置构建器
func NewControllerNetworkConfig() *client.ControllerNetworkBuilder {
	return client.NewControllerNetworkConfig()
}

// NewControllerMemberConfig 创建控制器成员配置构建器
func NewControllerMemberConfig() *client.ControllerMemberBuilder {
	return client.NewControllerMemberConfig()
}

// NewCentralNetworkConfig 创建云端网络配置构建器
func NewCentralNetworkConfig() *central.NetworkConfigBuilder {
	return central.NewNetworkConfig()
}

// NewCentralMemberConfig 创建云端成员配置构建器
func NewCentralMemberConfig() *central.MemberConfigBuilder {
	return central.NewMemberConfig()
}
