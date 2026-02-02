package feishu

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
)

// MultiTableClient 飞书多维表格客户端
type MultiTableClient struct {
	client *lark.Client
}

// NewMultiTableClient 新建客户端
func NewMultiTableClient(appID, appSecret string) *MultiTableClient {
	// 使用官方 SDK 创建客户端
	client := lark.NewClient(appID, appSecret)

	return &MultiTableClient{
		client: client,
	}
}

// GetClient 获取原始 lark.Client
func (c *MultiTableClient) GetClient() *lark.Client {
	return c.client
}
