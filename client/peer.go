package client

import (
	"encoding/json"
	"net/http"
)

// PeerService 节点管理服务接口
type PeerService interface {
	// List 列出所有 Peers
	List() ([]Peer, error)
	// Get 获取指定 Peer 信息
	Get(peerID string) (*Peer, error)
}

type peerService struct {
	client *client
}

// List 列出所有 Peers
func (s *peerService) List() ([]Peer, error) {
	data, err := s.client.do(http.MethodGet, "/peer", nil)
	if err != nil {
		return nil, err
	}

	var peers []Peer
	if err := json.Unmarshal(data, &peers); err != nil {
		return nil, err
	}

	return peers, nil
}

// Get 获取指定 Peer 信息
func (s *peerService) Get(peerID string) (*Peer, error) {
	data, err := s.client.do(http.MethodGet, "/peer/"+peerID, nil)
	if err != nil {
		return nil, err
	}

	var peer Peer
	if err := json.Unmarshal(data, &peer); err != nil {
		return nil, err
	}

	return &peer, nil
}
