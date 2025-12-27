package central

// NetworkConfigBuilder 网络配置构建器
type NetworkConfigBuilder struct {
	config CreateNetworkConfig
}

// NewNetworkConfig 创建网络配置构建器
func NewNetworkConfig() *NetworkConfigBuilder {
	return &NetworkConfigBuilder{}
}

// Name 设置网络名称
func (b *NetworkConfigBuilder) Name(name string) *NetworkConfigBuilder {
	b.config.Name = name
	return b
}

// Private 设置是否私有网络
func (b *NetworkConfigBuilder) Private(v bool) *NetworkConfigBuilder {
	b.config.Private = &v
	return b
}

// EnableBroadcast 设置是否启用广播
func (b *NetworkConfigBuilder) EnableBroadcast(v bool) *NetworkConfigBuilder {
	b.config.EnableBroadcast = &v
	return b
}

// MTU 设置 MTU
func (b *NetworkConfigBuilder) MTU(mtu int) *NetworkConfigBuilder {
	b.config.MTU = &mtu
	return b
}

// MulticastLimit 设置多播限制
func (b *NetworkConfigBuilder) MulticastLimit(limit int) *NetworkConfigBuilder {
	b.config.MulticastLimit = &limit
	return b
}

// AddRoute 添加路由
func (b *NetworkConfigBuilder) AddRoute(target string, via *string) *NetworkConfigBuilder {
	b.config.Routes = append(b.config.Routes, Route{Target: target, Via: via})
	return b
}

// AddIPPool 添加 IP 分配池
func (b *NetworkConfigBuilder) AddIPPool(start, end string) *NetworkConfigBuilder {
	b.config.IPAssignmentPools = append(b.config.IPAssignmentPools, IPAssignmentPool{
		IPRangeStart: start,
		IPRangeEnd:   end,
	})
	return b
}

// V4AssignMode 设置 IPv4 分配模式
func (b *NetworkConfigBuilder) V4AssignMode(zt bool) *NetworkConfigBuilder {
	b.config.V4AssignMode = &AssignMode{ZT: zt}
	return b
}

// V6AssignMode 设置 IPv6 分配模式
func (b *NetworkConfigBuilder) V6AssignMode(zt, rfc4193, n6plane bool) *NetworkConfigBuilder {
	b.config.V6AssignMode = &AssignMode{ZT: zt, RFC4193: rfc4193, N6Plane: n6plane}
	return b
}

// DNS 设置 DNS
func (b *NetworkConfigBuilder) DNS(domain string, servers ...string) *NetworkConfigBuilder {
	b.config.DNS = &DNS{Domain: domain, Servers: servers}
	return b
}

// Build 构建配置
func (b *NetworkConfigBuilder) Build() *CreateNetworkConfig {
	return &b.config
}

// MemberConfigBuilder 成员配置构建器
type MemberConfigBuilder struct {
	req UpdateMemberRequest
}

// NewMemberConfig 创建成员配置构建器
func NewMemberConfig() *MemberConfigBuilder {
	return &MemberConfigBuilder{
		req: UpdateMemberRequest{Config: &UpdateMemberConfig{}},
	}
}

// Name 设置成员名称
func (b *MemberConfigBuilder) Name(name string) *MemberConfigBuilder {
	b.req.Name = name
	return b
}

// Description 设置成员描述
func (b *MemberConfigBuilder) Description(desc string) *MemberConfigBuilder {
	b.req.Description = desc
	return b
}

// Authorized 设置是否授权
func (b *MemberConfigBuilder) Authorized(v bool) *MemberConfigBuilder {
	b.req.Config.Authorized = &v
	return b
}

// ActiveBridge 设置是否为活动桥接
func (b *MemberConfigBuilder) ActiveBridge(v bool) *MemberConfigBuilder {
	b.req.Config.ActiveBridge = &v
	return b
}

// NoAutoAssignIPs 设置是否禁用自动 IP 分配
func (b *MemberConfigBuilder) NoAutoAssignIPs(v bool) *MemberConfigBuilder {
	b.req.Config.NoAutoAssignIPs = &v
	return b
}

// IPAssignments 设置 IP 分配
func (b *MemberConfigBuilder) IPAssignments(ips ...string) *MemberConfigBuilder {
	b.req.Config.IPAssignments = ips
	return b
}

// Build 构建配置
func (b *MemberConfigBuilder) Build() *UpdateMemberRequest {
	return &b.req
}
