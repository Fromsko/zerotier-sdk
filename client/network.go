package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NetworkService 网络管理服务接口
type NetworkService interface {
	// List 列出所有已加入的网络
	List() ([]Network, error)
	// Get 获取指定网络详情
	Get(networkID string) (*Network, error)
	// Join 加入网络
	Join(networkID string) (*Network, error)
	// Leave 离开网络
	Leave(networkID string) error
	// Update 更新网络设置
	Update(networkID string, settings *NetworkSettings) (*Network, error)
}

type networkService struct {
	client *client
}

// List 列出所有已加入的网络
func (s *networkService) List() ([]Network, error) {
	data, err := s.client.do(http.MethodGet, "/network", nil)
	if err != nil {
		return nil, err
	}

	var networks []Network
	if err := json.Unmarshal(data, &networks); err != nil {
		return nil, err
	}

	return networks, nil
}

// Get 获取指定网络详情
func (s *networkService) Get(networkID string) (*Network, error) {
	data, err := s.client.do(http.MethodGet, "/network/"+networkID, nil)
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Join 加入网络
func (s *networkService) Join(networkID string) (*Network, error) {
	data, err := s.client.do(http.MethodPost, "/network/"+networkID, bytes.NewReader([]byte("{}")))
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Leave 离开网络
func (s *networkService) Leave(networkID string) error {
	_, err := s.client.do(http.MethodDelete, "/network/"+networkID, nil)
	return err
}

// Update 更新网络设置
func (s *networkService) Update(networkID string, settings *NetworkSettings) (*Network, error) {
	body, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/network/"+networkID, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// NetworkSettingsBuilder 网络设置构建器
type NetworkSettingsBuilder struct {
	settings NetworkSettings
}

// NewNetworkSettings 创建网络设置构建器
func NewNetworkSettings() *NetworkSettingsBuilder {
	return &NetworkSettingsBuilder{}
}

// AllowDNS 设置是否允许 DNS
func (b *NetworkSettingsBuilder) AllowDNS(v bool) *NetworkSettingsBuilder {
	b.settings.AllowDNS = &v
	return b
}

// AllowDefault 设置是否允许默认路由
func (b *NetworkSettingsBuilder) AllowDefault(v bool) *NetworkSettingsBuilder {
	b.settings.AllowDefault = &v
	return b
}

// AllowGlobal 设置是否允许全局路由
func (b *NetworkSettingsBuilder) AllowGlobal(v bool) *NetworkSettingsBuilder {
	b.settings.AllowGlobal = &v
	return b
}

// AllowManaged 设置是否允许托管路由
func (b *NetworkSettingsBuilder) AllowManaged(v bool) *NetworkSettingsBuilder {
	b.settings.AllowManaged = &v
	return b
}

// Build 构建设置
func (b *NetworkSettingsBuilder) Build() *NetworkSettings {
	return &b.settings
}
