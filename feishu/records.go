package feishu

import (
	"context"
	"fmt"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// CreateRecord 创建单个记录
func (c *MultiTableClient) CreateRecord(appToken, tableID string, fields map[string]interface{}) (string, error) {
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(fields).
			Build()).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.Create(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("创建记录失败: %v", err)
	}

	if !resp.Success() {
		return "", fmt.Errorf("创建记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return *resp.Data.Record.RecordId, nil
}

// BatchCreateRecords 批量创建记录
func (c *MultiTableClient) BatchCreateRecords(appToken, tableID string, records []CreateRecordRequest) ([]string, error) {
	// 转换为官方 SDK 格式
	larkRecords := make([]*larkbitable.AppTableRecord, 0, len(records))
	for _, record := range records {
		larkRecords = append(larkRecords, larkbitable.NewAppTableRecordBuilder().
			Fields(record.Fields).
			Build())
	}

	req := larkbitable.NewBatchCreateAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		Body(larkbitable.NewBatchCreateAppTableRecordReqBodyBuilder().
			Records(larkRecords).
			Build()).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.BatchCreate(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("批量创建记录失败: %v", err)
	}

	if !resp.Success() {
		return nil, fmt.Errorf("批量创建记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	recordIDs := make([]string, 0, len(resp.Data.Records))
	for _, record := range resp.Data.Records {
		recordIDs = append(recordIDs, *record.RecordId)
	}

	return recordIDs, nil
}

// UpdateRecord 更新记录
func (c *MultiTableClient) UpdateRecord(appToken, tableID, recordID string, fields map[string]interface{}) error {
	req := larkbitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		RecordId(recordID).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(fields).
			Build()).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.Update(context.Background(), req)
	if err != nil {
		return fmt.Errorf("更新记录失败: %v", err)
	}

	if !resp.Success() {
		return fmt.Errorf("更新记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return nil
}

// BatchUpdateRecords 批量更新记录
func (c *MultiTableClient) BatchUpdateRecords(appToken, tableID string, records []struct {
	RecordID string
	Fields   map[string]interface{}
}) error {
	// 转换为官方 SDK 格式
	larkRecords := make([]*larkbitable.AppTableRecord, 0, len(records))
	for _, record := range records {
		larkRecords = append(larkRecords, larkbitable.NewAppTableRecordBuilder().
			RecordId(record.RecordID).
			Fields(record.Fields).
			Build())
	}

	req := larkbitable.NewBatchUpdateAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		Body(larkbitable.NewBatchUpdateAppTableRecordReqBodyBuilder().
			Records(larkRecords).
			Build()).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.BatchUpdate(context.Background(), req)
	if err != nil {
		return fmt.Errorf("批量更新记录失败: %v", err)
	}

	if !resp.Success() {
		return fmt.Errorf("批量更新记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return nil
}

// DeleteRecord 删除记录
func (c *MultiTableClient) DeleteRecord(appToken, tableID, recordID string) error {
	req := larkbitable.NewDeleteAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		RecordId(recordID).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.Delete(context.Background(), req)
	if err != nil {
		return fmt.Errorf("删除记录失败: %v", err)
	}

	if !resp.Success() {
		return fmt.Errorf("删除记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return nil
}

// ListRecords 查询记录
func (c *MultiTableClient) ListRecords(appToken, tableID string, pageSize int, pageToken string) ([]map[string]interface{}, string, bool, error) {
	reqBuilder := larkbitable.NewListAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		PageSize(pageSize)

	if pageToken != "" {
		reqBuilder = reqBuilder.PageToken(pageToken)
	}

	req := reqBuilder.Build()

	resp, err := c.client.Bitable.AppTableRecord.List(context.Background(), req)
	if err != nil {
		return nil, "", false, fmt.Errorf("查询记录失败: %v", err)
	}

	if !resp.Success() {
		return nil, "", false, fmt.Errorf("查询记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	// 转换记录格式
	items := make([]map[string]interface{}, 0, len(resp.Data.Items))
	for _, item := range resp.Data.Items {
		items = append(items, item.Fields)
	}

	nextPageToken := ""
	if resp.Data.PageToken != nil {
		nextPageToken = *resp.Data.PageToken
	}

	hasMore := resp.Data.HasMore != nil && *resp.Data.HasMore

	return items, nextPageToken, hasMore, nil
}

// GetRecord 获取单个记录
func (c *MultiTableClient) GetRecord(appToken, tableID, recordID string) (map[string]interface{}, error) {
	req := larkbitable.NewGetAppTableRecordReqBuilder().
		AppToken(appToken).
		TableId(tableID).
		RecordId(recordID).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.Get(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("获取记录失败: %v", err)
	}

	if !resp.Success() {
		return nil, fmt.Errorf("获取记录失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Record.Fields, nil
}
