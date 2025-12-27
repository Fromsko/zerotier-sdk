package client

// NodeStatus 节点状态
type NodeStatus struct {
	Address           string `json:"address"`
	Clock             int64  `json:"clock"`
	Online            bool   `json:"online"`
	PlanetWorldID     int64  `json:"planetWorldId"`
	PublicIdentity    string `json:"publicIdentity"`
	TCPFallbackActive bool   `json:"tcpFallbackActive"`
	Version           string `json:"version"`
}

// Network 本地网络信息
type Network struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	Status            string   `json:"status"`
	Type              string   `json:"type"`
	MAC               string   `json:"mac"`
	MTU               int      `json:"mtu"`
	Bridge            bool     `json:"bridge"`
	BroadcastEnabled  bool     `json:"broadcastEnabled"`
	PortDeviceName    string   `json:"portDeviceName"`
	NetconfRevision   int      `json:"netconfRevision"`
	AssignedAddresses []string `json:"assignedAddresses"`
	AllowDNS          bool     `json:"allowDNS"`
	AllowDefault      bool     `json:"allowDefault"`
	AllowGlobal       bool     `json:"allowGlobal"`
	AllowManaged      bool     `json:"allowManaged"`
	DNS               *DNS     `json:"dns,omitempty"`
}

// DNS 配置
type DNS struct {
	Domain  string   `json:"domain"`
	Servers []string `json:"servers,omitempty"`
}

// NetworkSettings 网络设置（用于更新）
type NetworkSettings struct {
	AllowDNS     *bool `json:"allowDNS,omitempty"`
	AllowDefault *bool `json:"allowDefault,omitempty"`
	AllowGlobal  *bool `json:"allowGlobal,omitempty"`
	AllowManaged *bool `json:"allowManaged,omitempty"`
}

// Peer 节点信息
type Peer struct {
	Address string     `json:"address"`
	Version string     `json:"version"`
	Role    string     `json:"role"`
	Latency int        `json:"latency"`
	Paths   []PeerPath `json:"paths"`
}

// PeerPath 节点路径
type PeerPath struct {
	Active        bool   `json:"active"`
	Address       string `json:"address"`
	Expired       bool   `json:"expired"`
	LastReceive   int64  `json:"lastReceive"`
	LastSend      int64  `json:"lastSend"`
	Preferred     bool   `json:"preferred"`
	TrustedPathID int64  `json:"trustedPathId"`
}

// ControllerStatus 控制器状态
type ControllerStatus struct {
	Controller bool  `json:"controller"`
	APIVersion int   `json:"apiVersion"`
	Clock      int64 `json:"clock"`
}

// ControllerNetwork 控制器网络配置
type ControllerNetwork struct {
	ID                string             `json:"id"`
	Name              string             `json:"name"`
	Private           bool               `json:"private"`
	CreationTime      int64              `json:"creationTime"`
	Revision          int                `json:"revision"`
	MulticastLimit    int                `json:"multicastLimit"`
	EnableBroadcast   bool               `json:"enableBroadcast"`
	Routes            []Route            `json:"routes"`
	IPAssignmentPools []IPAssignmentPool `json:"ipAssignmentPools"`
	V4AssignMode      *AssignMode        `json:"v4AssignMode,omitempty"`
	V6AssignMode      *AssignMode        `json:"v6AssignMode,omitempty"`
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
	ZT bool `json:"zt"`
}

// ControllerMember 控制器成员
type ControllerMember struct {
	ID                   string   `json:"id"`
	Address              string   `json:"address"`
	NetworkID            string   `json:"networkId"`
	Authorized           bool     `json:"authorized"`
	ActiveBridge         bool     `json:"activeBridge"`
	IPAssignments        []string `json:"ipAssignments"`
	NoAutoAssignIPs      bool     `json:"noAutoAssignIps"`
	Revision             int      `json:"revision"`
	CreationTime         int64    `json:"creationTime"`
	LastAuthorizedTime   int64    `json:"lastAuthorizedTime"`
	LastDeauthorizedTime int64    `json:"lastDeauthorizedTime"`
}

// ControllerNetworkConfig 创建/更新网络的配置
type ControllerNetworkConfig struct {
	Name              string             `json:"name,omitempty"`
	Private           *bool              `json:"private,omitempty"`
	EnableBroadcast   *bool              `json:"enableBroadcast,omitempty"`
	MulticastLimit    *int               `json:"multicastLimit,omitempty"`
	Routes            []Route            `json:"routes,omitempty"`
	IPAssignmentPools []IPAssignmentPool `json:"ipAssignmentPools,omitempty"`
	V4AssignMode      *AssignMode        `json:"v4AssignMode,omitempty"`
	V6AssignMode      *AssignMode        `json:"v6AssignMode,omitempty"`
}

// ControllerMemberConfig 成员配置
type ControllerMemberConfig struct {
	Authorized      *bool    `json:"authorized,omitempty"`
	ActiveBridge    *bool    `json:"activeBridge,omitempty"`
	IPAssignments   []string `json:"ipAssignments,omitempty"`
	NoAutoAssignIPs *bool    `json:"noAutoAssignIps,omitempty"`
}
