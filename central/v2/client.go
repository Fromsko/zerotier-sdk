package v2

import (
	"context"
	"net/http"
	"time"
)

// DefaultBaseURL Central V2 API 默认地址
const DefaultBaseURL = "https://central.zerotier.com"

// NewClientWithToken 使用 Bearer Token 创建 Central V2 客户端
//
// token 应使用 ZeroTier Central Service Account API Key
func NewClientWithToken(token string, opts ...ClientOption) (*ClientWithResponses, error) {
	allOpts := append([]ClientOption{
		WithToken(token),
		WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
	}, opts...)

	return NewClientWithResponses(DefaultBaseURL, allOpts...)
}

// WithToken 返回设置 Authorization: Bearer <token> 的 ClientOption
func WithToken(token string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+token)
		return nil
	})
}
