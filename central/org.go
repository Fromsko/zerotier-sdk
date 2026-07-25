package central

import (
	"encoding/json"
	"net/http"
)

// OrganizationService 组织管理服务接口
type OrganizationService interface {
	// Get 获取当前用户所属组织
	Get() (*Organization, error)
	// GetByID 根据 ID 获取组织信息
	GetByID(orgID string) (*Organization, error)
	// Members 获取组织成员列表
	Members(orgID string) ([]OrganizationMember, error)
}

type organizationService struct {
	client *client
}

// Get 获取当前用户所属组织
func (s *organizationService) Get() (*Organization, error) {
	data, err := s.client.do(http.MethodGet, "/org", nil)
	if err != nil {
		return nil, err
	}

	var org Organization
	if err := json.Unmarshal(data, &org); err != nil {
		return nil, err
	}

	return &org, nil
}

// GetByID 根据 ID 获取组织信息
func (s *organizationService) GetByID(orgID string) (*Organization, error) {
	data, err := s.client.do(http.MethodGet, "/org/"+orgID, nil)
	if err != nil {
		return nil, err
	}

	var org Organization
	if err := json.Unmarshal(data, &org); err != nil {
		return nil, err
	}

	return &org, nil
}

// Members 获取组织成员列表
func (s *organizationService) Members(orgID string) ([]OrganizationMember, error) {
	data, err := s.client.do(http.MethodGet, "/org/"+orgID+"/user", nil)
	if err != nil {
		return nil, err
	}

	var members []OrganizationMember
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, err
	}

	return members, nil
}
