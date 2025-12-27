package central

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NetworkService 网络管理服务接口
type NetworkService interface {
	// List 列出所有网络
	List() ([]Network, error)
	// Get 获取网络详情
	Get(networkID string) (*Network, error)
	// Create 创建新网络
	Create(config *CreateNetworkConfig) (*Network, error)
	// Update 更新网络配置
	Update(networkID string, config *CreateNetworkConfig) (*Network, error)
	// Delete 删除网络
	Delete(networkID string) error
	// Members 获取成员服务
	Members(networkID string) MemberService
}

type networkService struct {
	client *client
}

// List 列出所有网络
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

// Get 获取网络详情
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

// Create 创建新网络
func (s *networkService) Create(config *CreateNetworkConfig) (*Network, error) {
	req := &CreateNetworkRequest{Config: config}
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/network", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var network Network
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// Update 更新网络配置
func (s *networkService) Update(networkID string, config *CreateNetworkConfig) (*Network, error) {
	req := &CreateNetworkRequest{Config: config}
	body, err := json.Marshal(req)
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

// Delete 删除网络
func (s *networkService) Delete(networkID string) error {
	_, err := s.client.do(http.MethodDelete, "/network/"+networkID, nil)
	return err
}

// Members 获取成员服务
func (s *networkService) Members(networkID string) MemberService {
	return &memberService{client: s.client, networkID: networkID}
}
