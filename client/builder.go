package client

// ControllerNetworkBuilder 控制器网络配置构建器
type ControllerNetworkBuilder struct {
	config ControllerNetworkConfig
}

// NewControllerNetworkConfig 创建网络配置构建器
func NewControllerNetworkConfig() *ControllerNetworkBuilder {
	return &ControllerNetworkBuilder{}
}

// Name 设置网络名称
func (b *ControllerNetworkBuilder) Name(name string) *ControllerNetworkBuilder {
	b.config.Name = name
	return b
}

// Private 设置是否私有网络
func (b *ControllerNetworkBuilder) Private(v bool) *ControllerNetworkBuilder {
	b.config.Private = &v
	return b
}

// EnableBroadcast 设置是否启用广播
func (b *ControllerNetworkBuilder) EnableBroadcast(v bool) *ControllerNetworkBuilder {
	b.config.EnableBroadcast = &v
	return b
}

// MulticastLimit 设置多播限制
func (b *ControllerNetworkBuilder) MulticastLimit(limit int) *ControllerNetworkBuilder {
	b.config.MulticastLimit = &limit
	return b
}

// AddRoute 添加路由
func (b *ControllerNetworkBuilder) AddRoute(target string, via *string) *ControllerNetworkBuilder {
	b.config.Routes = append(b.config.Routes, Route{Target: target, Via: via})
	return b
}

// AddIPPool 添加 IP 分配池
func (b *ControllerNetworkBuilder) AddIPPool(start, end string) *ControllerNetworkBuilder {
	b.config.IPAssignmentPools = append(b.config.IPAssignmentPools, IPAssignmentPool{
		IPRangeStart: start,
		IPRangeEnd:   end,
	})
	return b
}

// V4AssignMode 设置 IPv4 分配模式
func (b *ControllerNetworkBuilder) V4AssignMode(zt bool) *ControllerNetworkBuilder {
	b.config.V4AssignMode = &AssignMode{ZT: zt}
	return b
}

// V6AssignMode 设置 IPv6 分配模式
func (b *ControllerNetworkBuilder) V6AssignMode(zt bool) *ControllerNetworkBuilder {
	b.config.V6AssignMode = &AssignMode{ZT: zt}
	return b
}

// Build 构建配置
func (b *ControllerNetworkBuilder) Build() *ControllerNetworkConfig {
	return &b.config
}

// ControllerMemberBuilder 成员配置构建器
type ControllerMemberBuilder struct {
	config ControllerMemberConfig
}

// NewControllerMemberConfig 创建成员配置构建器
func NewControllerMemberConfig() *ControllerMemberBuilder {
	return &ControllerMemberBuilder{}
}

// Authorized 设置是否授权
func (b *ControllerMemberBuilder) Authorized(v bool) *ControllerMemberBuilder {
	b.config.Authorized = &v
	return b
}

// ActiveBridge 设置是否为活动桥接
func (b *ControllerMemberBuilder) ActiveBridge(v bool) *ControllerMemberBuilder {
	b.config.ActiveBridge = &v
	return b
}

// IPAssignments 设置 IP 分配
func (b *ControllerMemberBuilder) IPAssignments(ips ...string) *ControllerMemberBuilder {
	b.config.IPAssignments = ips
	return b
}

// NoAutoAssignIPs 设置是否禁用自动 IP 分配
func (b *ControllerMemberBuilder) NoAutoAssignIPs(v bool) *ControllerMemberBuilder {
	b.config.NoAutoAssignIPs = &v
	return b
}

// Build 构建配置
func (b *ControllerMemberBuilder) Build() *ControllerMemberConfig {
	return &b.config
}
