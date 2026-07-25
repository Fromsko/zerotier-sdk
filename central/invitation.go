package central

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// InvitationService 组织邀请管理服务接口
type InvitationService interface {
	// List 列出当前用户的组织邀请
	List() ([]OrganizationInvitation, error)
	// Invite 发送组织邀请
	Invite(req *OrganizationInvitation) (*OrganizationInvitation, error)
	// Get 获取单个邀请详情
	Get(inviteID string) (*OrganizationInvitation, error)
	// Accept 接受邀请
	Accept(inviteID string) (*OrganizationInvitation, error)
	// Decline 拒绝/取消邀请
	Decline(inviteID string) error
}

type invitationService struct {
	client *client
}

// List 列出当前用户的组织邀请
func (s *invitationService) List() ([]OrganizationInvitation, error) {
	data, err := s.client.do(http.MethodGet, "/org-invitation", nil)
	if err != nil {
		return nil, err
	}

	var invites []OrganizationInvitation
	if err := json.Unmarshal(data, &invites); err != nil {
		return nil, err
	}

	return invites, nil
}

// Invite 发送组织邀请
func (s *invitationService) Invite(req *OrganizationInvitation) (*OrganizationInvitation, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/org-invitation", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var invite OrganizationInvitation
	if err := json.Unmarshal(data, &invite); err != nil {
		return nil, err
	}

	return &invite, nil
}

// Get 获取单个邀请详情
func (s *invitationService) Get(inviteID string) (*OrganizationInvitation, error) {
	data, err := s.client.do(http.MethodGet, "/org-invitation/"+inviteID, nil)
	if err != nil {
		return nil, err
	}

	var invite OrganizationInvitation
	if err := json.Unmarshal(data, &invite); err != nil {
		return nil, err
	}

	return &invite, nil
}

// Accept 接受邀请
func (s *invitationService) Accept(inviteID string) (*OrganizationInvitation, error) {
	data, err := s.client.do(http.MethodPost, "/org-invitation/"+inviteID, nil)
	if err != nil {
		return nil, err
	}

	var invite OrganizationInvitation
	if err := json.Unmarshal(data, &invite); err != nil {
		return nil, err
	}

	return &invite, nil
}

// Decline 拒绝/取消邀请
func (s *invitationService) Decline(inviteID string) error {
	_, err := s.client.do(http.MethodDelete, "/org-invitation/"+inviteID, nil)
	return err
}
