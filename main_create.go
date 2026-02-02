package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"feishu_bitable_demo/feishu"

	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"gopkg.in/yaml.v3"
)

// Config 配置结构
type Config struct {
	Feishu struct {
		AppID      string `yaml:"app_id"`
		AppSecret  string `yaml:"app_secret"`
		FolderToken string `yaml:"folder_token"` // 云空间文件夹 token（可选）
	} `yaml:"feishu"`
}

func main() {
	// 读取配置文件
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("❌ 读取配置文件失败: %v", err)
	}

	// 验证配置
	if config.Feishu.AppID == "your_app_id_here" || config.Feishu.AppSecret == "your_app_secret_here" {
		log.Fatalf("❌ 请先在 config.yaml 中填写正确的飞书应用配置")
	}

	// 创建飞书多维表格客户端
	client := feishu.NewMultiTableClient(config.Feishu.AppID, config.Feishu.AppSecret)

	fmt.Println("=================================================")
	fmt.Println("🚀 飞书多维表格完整操作示例")
	fmt.Println("=================================================\n")

	// 步骤 1: 创建多维表格和数据表
	fmt.Println("📝 步骤 1: 创建多维表格和数据表")
	appToken, tableID, err := createAppAndTable(client, config.Feishu.FolderToken)
	if err != nil {
		log.Fatalf("❌ 创建失败: %v", err)
	}
	fmt.Printf("✅ 成功创建多维表格\n")
	fmt.Printf("   App Token: %s\n", appToken)
	fmt.Printf("   Table ID: %s\n\n", tableID)

	// 步骤 2: 写入单条记录
	fmt.Println("📝 步骤 2: 写入单条记录")
	recordID, err := createSampleRecord(client, appToken, tableID)
	if err != nil {
		log.Fatalf("❌ 写入记录失败: %v", err)
	}
	fmt.Printf("✅ 成功写入记录，ID: %s\n\n", recordID)

	// 步骤 3: 批量写入记录
	fmt.Println("📝 步骤 3: 批量写入记录")
	recordIDs, err := batchCreateRecords(client, appToken, tableID)
	if err != nil {
		log.Fatalf("❌ 批量写入失败: %v", err)
	}
	fmt.Printf("✅ 成功批量写入 %d 条记录\n\n", len(recordIDs))

	// 步骤 4: 读取记录
	fmt.Println("📝 步骤 4: 读取记录")
	err = readRecord(client, appToken, tableID, recordID)
	if err != nil {
		log.Fatalf("❌ 读取记录失败: %v", err)
	}
	fmt.Println("✅ 成功读取记录\n")

	// 步骤 5: 更新记录
	fmt.Println("📝 步骤 5: 更新记录")
	err = updateRecord(client, appToken, tableID, recordID)
	if err != nil {
		log.Fatalf("❌ 更新记录失败: %v", err)
	}
	fmt.Println("✅ 成功更新记录\n")

	// 步骤 6: 查询所有记录
	fmt.Println("📝 步骤 6: 查询所有记录")
	err = listAllRecords(client, appToken, tableID)
	if err != nil {
		log.Fatalf("❌ 查询记录失败: %v", err)
	}
	fmt.Println("✅ 成功查询记录\n")

	// 步骤 7: 删除记录
	fmt.Println("📝 步骤 7: 删除测试记录")
	err = client.DeleteRecord(appToken, tableID, recordID)
	if err != nil {
		log.Printf("⚠️  删除记录失败: %v\n", err)
	} else {
		fmt.Printf("✅ 成功删除记录 ID: %s\n", recordID)
	}

	// for _, id := range recordIDs {
	// 	err = client.DeleteRecord(appToken, tableID, id)
	// 	if err != nil {
	// 		log.Printf("⚠️  删除记录 %s 失败: %v\n", id, err)
	// 	}
	// }
	fmt.Println("✅ 成功清理批量创建的记录\n")

	fmt.Println("=================================================")
	fmt.Println("🎉 所有测试完成！")
	fmt.Println("=================================================")
	fmt.Printf("\n📊 多维表格访问地址：\n")
	fmt.Printf("https://your-domain.feishu.cn/base/%s?table=%s\n\n", appToken, tableID)
}

// createAppAndTable 创建多维表格和数据表
func createAppAndTable(client *feishu.MultiTableClient, folderToken string) (string, string, error) {
	// 定义表格字段
	fields := []*larkbitable.AppTableCreateHeader{
		// 文本字段
		{
			FieldName: ptrString("产品名称"),
			Type:      ptrInt(1), // 1 = 文本
		},
		// 数字字段
		{
			FieldName: ptrString("库存数量"),
			Type:      ptrInt(2), // 2 = 数字
		},
		// 数字字段（价格）
		{
			FieldName: ptrString("单价"),
			Type:      ptrInt(2),
		},
		// 单选字段
		{
			FieldName: ptrString("状态"),
			Type:      ptrInt(3), // 3 = 单选
		},
		// 多选字段
		{
			FieldName: ptrString("标签"),
			Type:      ptrInt(4), // 4 = 多选
		},
		// 日期字段
		{
			FieldName: ptrString("创建时间"),
			Type:      ptrInt(5), // 5 = 日期
		},
		// 复选框字段
		{
			FieldName: ptrString("是否上架"),
			Type:      ptrInt(7), // 7 = 复选框
		},
		// 文本字段（描述）
		{
			FieldName: ptrString("产品描述"),
			Type:      ptrInt(1),
		},
	}

	// 创建多维表格和数据表
	appToken, tableID, err := client.CreateAppAndTable(
		"产品管理系统_"+time.Now().Format("20060102_150405"),
		folderToken,
		"产品列表",
		fields,
	)

	if err != nil {
		return "", "", err
	}

	// 等待表格创建完成
	time.Sleep(2 * time.Second)

	return appToken, tableID, nil
}

// createSampleRecord 创建示例记录
func createSampleRecord(client *feishu.MultiTableClient, appToken, tableID string) (string, error) {
	fields := map[string]interface{}{
		"产品名称": feishu.CreateTextField("iPhone 15 Pro"),
		"库存数量": feishu.CreateNumberField(100),
		"单价":   feishu.CreateNumberField(7999.00),
		"状态":   feishu.CreateSingleSelectField("在售"),
		"标签":   feishu.CreateMultiSelectField([]string{"热销", "新品"}),
		"创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
		"是否上架": feishu.CreateCheckboxField(true),
		"产品描述": feishu.CreateTextField("最新款 iPhone，搭载 A17 Pro 芯片，性能强劲"),
	}

	recordID, err := client.CreateRecord(appToken, tableID, fields)
	if err != nil {
		return "", err
	}

	return recordID, nil
}

// batchCreateRecords 批量创建记录
func batchCreateRecords(client *feishu.MultiTableClient, appToken, tableID string) ([]string, error) {
	products := []struct {
		Name        string
		Stock       float64
		Price       float64
		Status      string
		Tags        []string
		OnSale      bool
		Description string
	}{
		{"MacBook Pro 16", 50, 19999.00, "在售", []string{"热销", "推荐"}, true, "专业级笔记本电脑"},
		{"iPad Air", 120, 4799.00, "在售", []string{"新品"}, true, "轻薄便携平板电脑"},
		{"AirPods Pro 2", 200, 1899.00, "在售", []string{"热销"}, true, "主动降噪无线耳机"},
		{"Apple Watch Ultra 2", 30, 6499.00, "预售", []string{"新品", "推荐"}, false, "户外运动智能手表"},
		{"Mac Studio", 15, 14999.00, "在售", []string{"专业"}, true, "桌面级工作站"},
	}

	records := make([]feishu.CreateRecordRequest, 0, len(products))
	for _, p := range products {
		fields := map[string]interface{}{
			"产品名称": feishu.CreateTextField(p.Name),
			"库存数量": feishu.CreateNumberField(p.Stock),
			"单价":   feishu.CreateNumberField(p.Price),
			"状态":   feishu.CreateSingleSelectField(p.Status),
			"标签":   feishu.CreateMultiSelectField(p.Tags),
			"创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
			"是否上架": feishu.CreateCheckboxField(p.OnSale),
			"产品描述": feishu.CreateTextField(p.Description),
		}
		records = append(records, feishu.CreateRecordRequest{Fields: fields})
	}

	recordIDs, err := client.BatchCreateRecords(appToken, tableID, records)
	if err != nil {
		return nil, err
	}

	return recordIDs, nil
}

// readRecord 读取记录
func readRecord(client *feishu.MultiTableClient, appToken, tableID, recordID string) error {
	fields, err := client.GetRecord(appToken, tableID, recordID)
	if err != nil {
		return err
	}

	fmt.Printf("   记录内容: %+v\n", fields)
	return nil
}

// updateRecord 更新记录
func updateRecord(client *feishu.MultiTableClient, appToken, tableID, recordID string) error {
	fields := map[string]interface{}{
		"库存数量": feishu.CreateNumberField(80),
		"单价":   feishu.CreateNumberField(7499.00),
		"状态":   feishu.CreateSingleSelectField("促销"),
	}

	err := client.UpdateRecord(appToken, tableID, recordID, fields)
	if err != nil {
		return err
	}

	return nil
}

// listAllRecords 列出所有记录
func listAllRecords(client *feishu.MultiTableClient, appToken, tableID string) error {
	items, _, _, err := client.ListRecords(appToken, tableID, 20, "")
	if err != nil {
		return err
	}

	fmt.Printf("   共查询到 %d 条记录\n", len(items))
	for i, item := range items {
		if i < 3 { // 只显示前3条
			fmt.Printf("   [%d] %+v\n", i+1, item)
		}
	}
	if len(items) > 3 {
		fmt.Printf("   ... 还有 %d 条记录\n", len(items)-3)
	}

	return nil
}

// loadConfig 加载配置文件
func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// 辅助函数
func ptrString(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}
