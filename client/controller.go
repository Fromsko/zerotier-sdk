package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ControllerService 控制器管理服务接口（自托管时可用）
type ControllerService interface {
	// Status 获取控制器状态
	Status() (*ControllerStatus, error)
	// ListNetworks 列出控制器管理的所有网络
	ListNetworks() ([]string, error)
	// GetNetwork 获取网络配置
	GetNetwork(networkID string) (*ControllerNetwork, error)
	// CreateNetwork 创建新网络
	CreateNetwork(nodeID string, config *ControllerNetworkConfig) (*ControllerNetwork, error)
	// UpdateNetwork 更新网络配置
	UpdateNetwork(networkID string, config *ControllerNetworkConfig) (*ControllerNetwork, error)
	// DeleteNetwork 删除网络
	DeleteNetwork(networkID string) error
	// ListMembers 列出网络成员
	ListMembers(networkID string) ([]string, error)
	// GetMember 获取成员信息
	GetMember(networkID, memberID string) (*ControllerMember, error)
	// UpdateMember 更新成员配置
	UpdateMember(networkID, memberID string, config *ControllerMemberConfig) (*ControllerMember, error)
	// DeleteMember 删除成员
	DeleteMember(networkID, memberID string) error
}

type controllerService struct {
	client *client
}

// Status 获取控制器状态
func (s *controllerService) Status() (*ControllerStatus, error) {
	data, err := s.client.do(http.MethodGet, "/controller", nil)
	if err != nil {
		return nil, err
	}

	var status ControllerStatus
	if err := json.Unmarshal(data, &status); err != nil {
		return nil, err
	}

	return &status, nil
}

// ListNetworks 列出控制器管理的所有网络
func (s *controllerService) ListNetworks() ([]string, error) {
	data, err := s.client.do(http.MethodGet, "/controller/network", nil)
	if err != nil {
		return nil, err
	}

	var networks []string
	if err := json.Unmarshal(data, &networks); err != nil {
		return nil, err
	}

	return networks, nil
}

// GetNetwork 获取网络配置
func (s *controllerService) GetNetwork(networkID string) (*ControllerNetwork, error) {
	data, err := s.client.do(http.MethodGet, "/controller/network/"+networkID, nil)
	if err != nil {
		return nil, err
	}

	var network ControllerNetwork
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// CreateNetwork 创建新网络（networkID 格式：nodeID + 6个下划线，如 "abc123______"）
func (s *controllerService) CreateNetwork(nodeID string, config *ControllerNetworkConfig) (*ControllerNetwork, error) {
	networkID := nodeID + "______"

	var body []byte
	var err error
	if config != nil {
		body, err = json.Marshal(config)
		if err != nil {
			return nil, err
		}
	} else {
		body = []byte("{}")
	}

	data, err := s.client.do(http.MethodPost, "/controller/network/"+networkID, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var network ControllerNetwork
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// UpdateNetwork 更新网络配置
func (s *controllerService) UpdateNetwork(networkID string, config *ControllerNetworkConfig) (*ControllerNetwork, error) {
	body, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/controller/network/"+networkID, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var network ControllerNetwork
	if err := json.Unmarshal(data, &network); err != nil {
		return nil, err
	}

	return &network, nil
}

// DeleteNetwork 删除网络
func (s *controllerService) DeleteNetwork(networkID string) error {
	_, err := s.client.do(http.MethodDelete, "/controller/network/"+networkID, nil)
	return err
}

// ListMembers 列出网络成员
func (s *controllerService) ListMembers(networkID string) ([]string, error) {
	data, err := s.client.do(http.MethodGet, "/controller/network/"+networkID+"/member", nil)
	if err != nil {
		return nil, err
	}

	var members []string
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, err
	}

	return members, nil
}

// GetMember 获取成员信息
func (s *controllerService) GetMember(networkID, memberID string) (*ControllerMember, error) {
	data, err := s.client.do(http.MethodGet, "/controller/network/"+networkID+"/member/"+memberID, nil)
	if err != nil {
		return nil, err
	}

	var member ControllerMember
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, err
	}

	return &member, nil
}

// UpdateMember 更新成员配置
func (s *controllerService) UpdateMember(networkID, memberID string, config *ControllerMemberConfig) (*ControllerMember, error) {
	body, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/controller/network/"+networkID+"/member/"+memberID, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var member ControllerMember
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, err
	}

	return &member, nil
}

// DeleteMember 删除成员
func (s *controllerService) DeleteMember(networkID, memberID string) error {
	_, err := s.client.do(http.MethodDelete, "/controller/network/"+networkID+"/member/"+memberID, nil)
	return err
}
