// Package v2 提供 ZeroTier Central API v2 的客户端与类型
//
// 本包代码由 oapi-codegen 根据 ZeroTier 官方 OpenAPI v2 规范生成，
// 覆盖 Organizations、Network Groups、Networks、Members、Service Accounts、
// API Keys、Users、Webhooks、Flow Rules 等全部服务端点。
//
// 认证方式：使用 Central Service Account API Key，请求头携带
//
//	Authorization: Bearer <token>
//
// 快速开始：
//
//	client, err := v2.NewClientWithToken("your_service_account_key")
//	if err != nil { ... }
//	resp, err := client.ListNetworksWithResponse(context.Background(), nil)
package v2
