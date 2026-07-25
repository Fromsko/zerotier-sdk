package central

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// UserService 用户管理服务接口
type UserService interface {
	// Get 获取用户信息
	Get(userID string) (*User, error)
	// Update 更新用户信息
	Update(userID string, req *User) (*User, error)
	// Delete 删除用户
	Delete(userID string) error
	// AddToken 为用户添加 API Token
	AddToken(userID string, req *APIToken) (*APIToken, error)
	// DeleteToken 删除用户的 API Token
	DeleteToken(userID, tokenName string) error
}

type userService struct {
	client *client
}

// Get 获取用户信息
func (s *userService) Get(userID string) (*User, error) {
	data, err := s.client.do(http.MethodGet, "/user/"+userID, nil)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Update 更新用户信息
func (s *userService) Update(userID string, req *User) (*User, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/user/"+userID, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Delete 删除用户
func (s *userService) Delete(userID string) error {
	_, err := s.client.do(http.MethodDelete, "/user/"+userID, nil)
	return err
}

// AddToken 为用户添加 API Token
func (s *userService) AddToken(userID string, req *APIToken) (*APIToken, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data, err := s.client.do(http.MethodPost, "/user/"+userID+"/token", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var token APIToken
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

// DeleteToken 删除用户的 API Token
func (s *userService) DeleteToken(userID, tokenName string) error {
	_, err := s.client.do(http.MethodDelete, "/user/"+userID+"/token/"+tokenName, nil)
	return err
}
