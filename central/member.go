package central

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// MemberService 成员管理服务接口
type MemberService interface {
	// List 列出网络所有成员
	List() ([]Member, error)
	// Get 获取成员详情
	Get(memberID string) (*Member, error)
	// Update 更新成员配置
	Update(memberID string, req *UpdateMemberRequest) (*Member, error)
	// Authorize 授权成员
	Authorize(memberID string) (*Member, error)
	// Deauthorize 取消授权
	Deauthorize(memberID string) (*Member, error)
	// Delete 删除成员
	Delete(memberID string) error
	// SetName 设置成员昵称
	SetName(memberID string, name string) (*Member, error)
	// SetDescription 设置成员描述
	SetDescription(memberID string, description string) (*Member, error)
	// SetIPAssignments 设置成员 IP 地址
	SetIPAssignments(memberID string, ips []string) (*Member, error)
}

type memberService struct {
	client    *client
	networkID string
}

func (s *memberService) basePath() string {
	return "/network/" + s.networkID + "/member"
}

// List 列出网络所有成员
func (s *memberService) List() ([]Member, error) {
	data, err := s.client.do(http.MethodGet, s.basePath(), nil)
	if err != nil {
		return nil, err
	}

	var members []Member
	if err := json.Unmarshal(data, &members); err != nil {
		return nil, err
	}

	return members, nil
}

// Get 获取成员详情
func (s *memberService) Get(memberID string) (*Member, error) {
	data, err := s.client.do(http.MethodGet, s.basePath()+"/"+memberID, nil)
	if err != nil {
		return nil, err
	}

	var member Member
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, err
	}

	return &member, nil
}

// Update 更新成员配置
func (s *memberService) Update(memberID string, req *UpdateMemberRequest) (*Member, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, s.basePath()+"/"+memberID, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var member Member
	if err := json.Unmarshal(data, &member); err != nil {
		return nil, err
	}

	return &member, nil
}

// Authorize 授权成员
func (s *memberService) Authorize(memberID string) (*Member, error) {
	authorized := true
	return s.Update(memberID, &UpdateMemberRequest{
		Config: &UpdateMemberConfig{Authorized: &authorized},
	})
}

// Deauthorize 取消授权
func (s *memberService) Deauthorize(memberID string) (*Member, error) {
	authorized := false
	return s.Update(memberID, &UpdateMemberRequest{
		Config: &UpdateMemberConfig{Authorized: &authorized},
	})
}

// Delete 删除成员
func (s *memberService) Delete(memberID string) error {
	_, err := s.client.do(http.MethodDelete, s.basePath()+"/"+memberID, nil)
	return err
}

// SetName 设置成员昵称
func (s *memberService) SetName(memberID string, name string) (*Member, error) {
	return s.Update(memberID, &UpdateMemberRequest{
		Name: name,
	})
}

// SetDescription 设置成员描述
func (s *memberService) SetDescription(memberID string, description string) (*Member, error) {
	return s.Update(memberID, &UpdateMemberRequest{
		Description: description,
	})
}

// SetIPAssignments 设置成员 IP 地址
func (s *memberService) SetIPAssignments(memberID string, ips []string) (*Member, error) {
	return s.Update(memberID, &UpdateMemberRequest{
		Config: &UpdateMemberConfig{
			IPAssignments: ips,
		},
	})
}
