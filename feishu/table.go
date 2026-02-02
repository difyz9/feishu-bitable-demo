package feishu

import (
	"context"
	"fmt"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// CreateApp 创建多维表格
func (c *MultiTableClient) CreateApp(name, folderToken string) (string, error) {
	req := larkbitable.NewCreateAppReqBuilder().
		ReqApp(larkbitable.NewReqAppBuilder().
			Name(name).
			FolderToken(folderToken).
			Build()).
		Build()

	resp, err := c.client.Bitable.App.Create(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("创建多维表格失败: %v", err)
	}

	if !resp.Success() {
		return "", fmt.Errorf("创建多维表格失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return *resp.Data.App.AppToken, nil
}

// CreateTable 创建数据表
func (c *MultiTableClient) CreateTable(appToken, tableName string, fields []*larkbitable.AppTableCreateHeader) (string, error) {
	table := larkbitable.NewReqTableBuilder().
		Name(tableName).
		Fields(fields).
		Build()

	req := larkbitable.NewCreateAppTableReqBuilder().
		AppToken(appToken).
		Body(larkbitable.NewCreateAppTableReqBodyBuilder().
			Table(table).
			Build()).
		Build()

	resp, err := c.client.Bitable.AppTable.Create(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("创建数据表失败: %v", err)
	}

	if !resp.Success() {
		return "", fmt.Errorf("创建数据表失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return *resp.Data.TableId, nil
}

// CreateAppAndTable 创建多维表格并添加数据表（组合操作）
func (c *MultiTableClient) CreateAppAndTable(appName, folderToken, tableName string, fields []*larkbitable.AppTableCreateHeader) (appToken, tableID string, err error) {
	// 创建多维表格
	appToken, err = c.CreateApp(appName, folderToken)
	if err != nil {
		return "", "", err
	}

	// 创建数据表
	tableID, err = c.CreateTable(appToken, tableName, fields)
	if err != nil {
		return appToken, "", err
	}

	return appToken, tableID, nil
}

// ListTables 列出多维表格中的所有数据表
func (c *MultiTableClient) ListTables(appToken string) ([]*larkbitable.AppTable, error) {
	req := larkbitable.NewListAppTableReqBuilder().
		AppToken(appToken).
		Build()

	resp, err := c.client.Bitable.AppTable.List(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("列出数据表失败: %v", err)
	}

	if !resp.Success() {
		return nil, fmt.Errorf("列出数据表失败 [code=%d]: %s", resp.Code, resp.Msg)
	}

	return resp.Data.Items, nil
}
