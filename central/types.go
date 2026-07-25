package central

// CentralStatus Central API 状态（包含当前用户信息）
type CentralStatus struct {
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	Clock        int64           `json:"clock"`
	Version      string          `json:"version"`
	APIVersion   string          `json:"apiVersion"`
	Uptime       int64           `json:"uptime"`
	User         *StatusUser     `json:"user,omitempty"`
	ReadOnlyMode bool            `json:"readOnlyMode"`
	LoginMethods map[string]bool `json:"loginMethods"`
}

// StatusUser 状态中的用户信息
type StatusUser struct {
	ID          string `json:"id"`
	OrgID       string `json:"orgId"`
	DisplayName string `json:"displayName"`
	SMSNumber   string `json:"smsNumber,omitempty"`
}

// User 用户信息（完整）
type User struct {
	ID                string       `json:"id"`
	OrgID             string       `json:"orgId"`
	GlobalPermissions *Permissions `json:"globalPermissions,omitempty"`
	DisplayName       string       `json:"displayName"`
	Email             string       `json:"email"`
	Auth              *AuthMethods `json:"auth,omitempty"`
	SMSNumber         string       `json:"smsNumber,omitempty"`
	Tokens            []string     `json:"tokens,omitempty"`
}

// Network 网络信息
type Network struct {
	ID                    string         `json:"id"`
	Clock                 int64          `json:"clock"`
	Config                *NetworkConfig `json:"config"`
	Description           string         `json:"description"`
	RulesSource           string         `json:"rulesSource"`
	OwnerID               string         `json:"ownerId"`
	OnlineMemberCount     int            `json:"onlineMemberCount"`
	AuthorizedMemberCount int            `json:"authorizedMemberCount"`
	TotalMemberCount      int            `json:"totalMemberCount"`
	CapabilitiesByName    map[string]int `json:"capabilitiesByName"`
	TagsByName            map[string]int `json:"tagsByName"`
}

// NetworkConfig 网络配置
type NetworkConfig struct {
	ID                string             `json:"id"`
	Name              string             `json:"name"`
	Private           bool               `json:"private"`
	CreationTime      int64              `json:"creationTime"`
	LastModified      int64              `json:"lastModified"`
	EnableBroadcast   bool               `json:"enableBroadcast"`
	MTU               int                `json:"mtu"`
	MulticastLimit    int                `json:"multicastLimit"`
	Routes            []Route            `json:"routes"`
	IPAssignmentPools []IPAssignmentPool `json:"ipAssignmentPools"`
	V4AssignMode      *AssignMode        `json:"v4AssignMode"`
	V6AssignMode      *AssignMode        `json:"v6AssignMode"`
	DNS               *DNS               `json:"dns,omitempty"`
}

// Route 路由配置
type Route struct {
	Target string  `json:"target"`
	Via    *string `json:"via,omitempty"`
}

// IPAssignmentPool IP 分配池
type IPAssignmentPool struct {
	IPRangeStart string `json:"ipRangeStart"`
	IPRangeEnd   string `json:"ipRangeEnd"`
}

// AssignMode IP 分配模式
type AssignMode struct {
	ZT      bool `json:"zt"`
	RFC4193 bool `json:"rfc4193,omitempty"`
	N6Plane bool `json:"6plane,omitempty"`
}

// DNS 配置
type DNS struct {
	Domain  string   `json:"domain"`
	Servers []string `json:"servers"`
}

// Permissions 通用权限对象
type Permissions struct {
	Authorize bool `json:"a"`
	Delete    bool `json:"d"`
	Modify    bool `json:"m"`
	Read      bool `json:"r"`
}

// PermissionsMap 用户 ID 到权限的映射
type PermissionsMap map[string]*Permissions

// AuthMethods 认证方式信息
type AuthMethods struct {
	Local  string `json:"local,omitempty"`
	OIDC   string `json:"oidc,omitempty"`
	Google string `json:"google,omitempty"`
}

// APIToken API Token 信息
type APIToken struct {
	TokenName string `json:"tokenName"`
	Token     string `json:"token,omitempty"`
}

// RandomToken 随机生成的 Token
type RandomToken struct {
	Clock int64  `json:"clock"`
	Hex   string `json:"hex"`
	Token string `json:"token"`
}

// NetworkUserPermissions 网络用户权限（用于设置特定用户的权限）
type NetworkUserPermissions struct {
	ID        string `json:"id"`
	Read      bool   `json:"r"`
	Authorize bool   `json:"a"`
	Modify    bool   `json:"m"`
	Delete    bool   `json:"d"`
}

// InviteStatus 邀请状态
type InviteStatus string

// 邀请状态常量
const (
	InvitePending  InviteStatus = "pending"
	InviteAccepted InviteStatus = "accepted"
	InviteCanceled InviteStatus = "canceled"
)

// OrganizationInvitation 组织邀请
type OrganizationInvitation struct {
	OrgID        string       `json:"orgId,omitempty"`
	Email        string       `json:"email"`
	ID           string       `json:"id,omitempty"`
	CreationTime int64        `json:"creation_time,omitempty"`
	Status       InviteStatus `json:"status,omitempty"`
	UpdateTime   int64        `json:"update_time,omitempty"`
	OwnerEmail   string       `json:"ownerEmail,omitempty"`
}

// OrganizationMember 组织成员
type OrganizationMember struct {
	OrgID  string `json:"orgId,omitempty"`
	UserID string `json:"userId"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
}

// Organization 组织信息
type Organization struct {
	ID         string                `json:"id"`
	OwnerID    string                `json:"ownerId"`
	OwnerEmail string                `json:"ownerEmail"`
	Members    []*OrganizationMember `json:"members,omitempty"`
	SSOConfig  *OrgSSOConfig         `json:"ssoConfig,omitempty"`
}

// SsoIssuer SSO Issuer 配置
type SsoIssuer struct {
	Provider              string `json:"provider"`
	ClientID              string `json:"clientId"`
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty"`
}

// OrgSSOConfig 组织 SSO 配置
type OrgSSOConfig struct {
	Enabled bool         `json:"enabled"`
	Issuers []*SsoIssuer `json:"issuers,omitempty"`
}

// NetworkSSOConfig 网络 SSO 配置
type NetworkSSOConfig struct {
	Enabled               bool     `json:"enabled"`
	Mode                  string   `json:"mode"`
	ClientID              string   `json:"clientId"`
	Issuer                string   `json:"issuer,omitempty"`
	Provider              string   `json:"provider,omitempty"`
	AuthorizationEndpoint string   `json:"authorizationEndpoint,omitempty"`
	AllowList             []string `json:"allowList,omitempty"`
}

// Member 网络成员
type Member struct {
	ID                  string        `json:"id"`
	NetworkID           string        `json:"networkId"`
	NodeID              string        `json:"nodeId"`
	Name                string        `json:"name"`
	Description         string        `json:"description"`
	Config              *MemberConfig `json:"config"`
	LastOnline          int64         `json:"lastOnline"`
	LastSeen            int64         `json:"lastSeen"`
	PhysicalAddress     string        `json:"physicalAddress"`
	ClientVersion       string        `json:"clientVersion"`
	ProtocolVersion     int           `json:"protocolVersion"`
	SupportsRulesEngine bool          `json:"supportsRulesEngine"`
}

// MemberConfig 成员配置
type MemberConfig struct {
	Authorized      bool     `json:"authorized"`
	ActiveBridge    bool     `json:"activeBridge"`
	NoAutoAssignIPs bool     `json:"noAutoAssignIps"`
	CreationTime    int64    `json:"creationTime"`
	IPAssignments   []string `json:"ipAssignments"`
	SSOExempt       bool     `json:"ssoExempt"`
}

// CreateNetworkRequest 创建网络请求
type CreateNetworkRequest struct {
	Config *CreateNetworkConfig `json:"config"`
}

// CreateNetworkConfig 创建网络配置
type CreateNetworkConfig struct {
	Name              string             `json:"name,omitempty"`
	Private           *bool              `json:"private,omitempty"`
	EnableBroadcast   *bool              `json:"enableBroadcast,omitempty"`
	MTU               *int               `json:"mtu,omitempty"`
	MulticastLimit    *int               `json:"multicastLimit,omitempty"`
	Routes            []Route            `json:"routes,omitempty"`
	IPAssignmentPools []IPAssignmentPool `json:"ipAssignmentPools,omitempty"`
	V4AssignMode      *AssignMode        `json:"v4AssignMode,omitempty"`
	V6AssignMode      *AssignMode        `json:"v6AssignMode,omitempty"`
	DNS               *DNS               `json:"dns,omitempty"`
}

// UpdateMemberRequest 更新成员请求
type UpdateMemberRequest struct {
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Config      *UpdateMemberConfig `json:"config,omitempty"`
}

// UpdateMemberConfig 更新成员配置
type UpdateMemberConfig struct {
	Authorized      *bool    `json:"authorized,omitempty"`
	ActiveBridge    *bool    `json:"activeBridge,omitempty"`
	NoAutoAssignIPs *bool    `json:"noAutoAssignIps,omitempty"`
	IPAssignments   []string `json:"ipAssignments,omitempty"`
}
