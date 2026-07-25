package central

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NetworkPermissionService 网络用户权限服务接口
type NetworkPermissionService interface {
	// SetUserPermissions 设置特定用户对网络的权限
	SetUserPermissions(req *NetworkUserPermissions) (*NetworkUserPermissions, error)
}

type networkPermissionService struct {
	client    *client
	networkID string
}

// SetUserPermissions 设置特定用户对网络的权限
func (s *networkPermissionService) SetUserPermissions(req *NetworkUserPermissions) (*NetworkUserPermissions, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/network/"+s.networkID+"/users", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var permissions NetworkUserPermissions
	if err := json.Unmarshal(data, &permissions); err != nil {
		return nil, err
	}

	return &permissions, nil
}
