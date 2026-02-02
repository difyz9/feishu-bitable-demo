package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"feishu_bitable_demo/feishu"

	"gopkg.in/yaml.v3"
)

// Config 配置结构
type Config struct {
	Feishu struct {
		AppID     string `yaml:"app_id"`
		AppSecret string `yaml:"app_secret"`
		AppToken  string `yaml:"app_token"`
		TableID   string `yaml:"table_id"`
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
	fmt.Println("🚀 飞书多维表格操作验证程序")
	fmt.Println("=================================================\n")

	// 1. 测试获取 Access Token（官方 SDK 会自动管理 token）
	fmt.Println("📝 步骤 1: 初始化飞书客户端")
	fmt.Printf("✅ 成功初始化客户端\n\n")

	// 2. 测试创建单个记录
	fmt.Println("📝 步骤 2: 测试创建单个记录")
	recordID, err := testCreateRecord(client, config.Feishu.AppToken, config.Feishu.TableID)
	if err != nil {
		log.Fatalf("❌ 创建记录失败: %v", err)
	}
	fmt.Printf("✅ 成功创建记录，ID: %s\n\n", recordID)

	// 等待一秒，确保记录已创建
	time.Sleep(1 * time.Second)

	// 3. 测试读取记录
	fmt.Println("📝 步骤 3: 测试读取记录")
	err = testGetRecord(client, config.Feishu.AppToken, config.Feishu.TableID, recordID)
	if err != nil {
		log.Fatalf("❌ 读取记录失败: %v", err)
	}
	fmt.Println("✅ 成功读取记录\n")

	// 4. 测试更新记录
	fmt.Println("📝 步骤 4: 测试更新记录")
	err = testUpdateRecord(client, config.Feishu.AppToken, config.Feishu.TableID, recordID)
	if err != nil {
		log.Fatalf("❌ 更新记录失败: %v", err)
	}
	fmt.Println("✅ 成功更新记录\n")

	// 5. 测试查询所有记录
	fmt.Println("📝 步骤 5: 测试查询所有记录")
	err = testListRecords(client, config.Feishu.AppToken, config.Feishu.TableID)
	if err != nil {
		log.Fatalf("❌ 查询记录失败: %v", err)
	}
	fmt.Println("✅ 成功查询记录\n")

	// 6. 测试批量创建记录
	fmt.Println("📝 步骤 6: 测试批量创建记录")
	batchRecordIDs, err := testBatchCreateRecords(client, config.Feishu.AppToken, config.Feishu.TableID)
	if err != nil {
		log.Fatalf("❌ 批量创建记录失败: %v", err)
	}
	fmt.Printf("✅ 成功批量创建 %d 条记录\n\n", len(batchRecordIDs))

	// 7. 测试批量更新记录
	fmt.Println("📝 步骤 7: 测试批量更新记录")
	err = testBatchUpdateRecords(client, config.Feishu.AppToken, config.Feishu.TableID, batchRecordIDs)
	if err != nil {
		log.Fatalf("❌ 批量更新记录失败: %v", err)
	}
	fmt.Println("✅ 成功批量更新记录\n")

	// 8. 测试删除记录
	fmt.Println("📝 步骤 8: 测试删除记录")
	// 删除第一条测试记录
	err = client.DeleteRecord(config.Feishu.AppToken, config.Feishu.TableID, recordID)
	if err != nil {
		log.Fatalf("❌ 删除记录失败: %v", err)
	}
	fmt.Printf("✅ 成功删除记录 ID: %s\n\n", recordID)

	// 删除批量创建的记录
	for _, id := range batchRecordIDs {
		err = client.DeleteRecord(config.Feishu.AppToken, config.Feishu.TableID, id)
		if err != nil {
			fmt.Printf("⚠️  删除记录 %s 失败: %v\n", id, err)
		}
	}
	fmt.Printf("✅ 成功清理批量创建的记录\n\n")

	fmt.Println("=================================================")
	fmt.Println("🎉 所有测试通过！飞书多维表格操作功能正常")
	fmt.Println("=================================================")
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

// testCreateRecord 测试创建记录
func testCreateRecord(client *feishu.MultiTableClient, appToken, tableID string) (string, error) {
	fields := map[string]interface{}{
		"名称":   feishu.CreateTextField("测试产品"),
		"数量":   feishu.CreateNumberField(100),
		"价格":   feishu.CreateNumberField(299.99),
		"描述":   feishu.CreateTextField("这是一个测试产品，用于验证飞书多维表格的创建功能"),
		"创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
		"是否上架": feishu.CreateCheckboxField(true),
	}

	recordID, err := client.CreateRecord(appToken, tableID, fields)
	if err != nil {
		return "", err
	}

	return recordID, nil
}

// testGetRecord 测试读取记录
func testGetRecord(client *feishu.MultiTableClient, appToken, tableID, recordID string) error {
	fields, err := client.GetRecord(appToken, tableID, recordID)
	if err != nil {
		return err
	}

	fmt.Printf("   记录字段: %+v\n", fields)
	return nil
}

// testUpdateRecord 测试更新记录
func testUpdateRecord(client *feishu.MultiTableClient, appToken, tableID, recordID string) error {
	fields := map[string]interface{}{
		"数量":   feishu.CreateNumberField(200),
		"价格":   feishu.CreateNumberField(399.99),
		"描述":   feishu.CreateTextField("更新后的产品描述"),
		"是否上架": feishu.CreateCheckboxField(false),
	}

	err := client.UpdateRecord(appToken, tableID, recordID, fields)
	if err != nil {
		return err
	}

	return nil
}

// testListRecords 测试查询记录
func testListRecords(client *feishu.MultiTableClient, appToken, tableID string) error {
	items, pageToken, hasMore, err := client.ListRecords(appToken, tableID, 10, "")
	if err != nil {
		return err
	}

	fmt.Printf("   查询到 %d 条记录\n", len(items))
	fmt.Printf("   是否有更多: %v\n", hasMore)
	if hasMore {
		fmt.Printf("   下一页 Token: %s\n", pageToken)
	}

	return nil
}

// testBatchCreateRecords 测试批量创建记录
func testBatchCreateRecords(client *feishu.MultiTableClient, appToken, tableID string) ([]string, error) {
	records := []feishu.CreateRecordRequest{
		{
			Fields: map[string]interface{}{
				"名称":   feishu.CreateTextField("批量产品 A"),
				"数量":   feishu.CreateNumberField(50),
				"价格":   feishu.CreateNumberField(199.99),
				"描述":   feishu.CreateTextField("批量创建的测试产品 A"),
				"创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
				"是否上架": feishu.CreateCheckboxField(true),
			},
		},
		{
			Fields: map[string]interface{}{
				"名称":   feishu.CreateTextField("批量产品 B"),
				"数量":   feishu.CreateNumberField(75),
				"价格":   feishu.CreateNumberField(249.99),
				"描述":   feishu.CreateTextField("批量创建的测试产品 B"),
				"创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
				"是否上架": feishu.CreateCheckboxField(false),
			},
		},
		{
			Fields: map[string]interface{}{
				"名称":   feishu.CreateTextField("批量产品 C"),
				"数量":   feishu.CreateNumberField(120),
				"价格":   feishu.CreateNumberField(349.99),
				"描述":   feishu.CreateTextField("批量创建的测试产品 C"),
				"创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
				"是否上架": feishu.CreateCheckboxField(true),
			},
		},
	}

	recordIDs, err := client.BatchCreateRecords(appToken, tableID, records)
	if err != nil {
		return nil, err
	}

	return recordIDs, nil
}

// testBatchUpdateRecords 测试批量更新记录
func testBatchUpdateRecords(client *feishu.MultiTableClient, appToken, tableID string, recordIDs []string) error {
	if len(recordIDs) == 0 {
		return nil
	}

	var records []struct {
		RecordID string
		Fields   map[string]interface{}
	}

	for i, recordID := range recordIDs {
		records = append(records, struct {
			RecordID string
			Fields   map[string]interface{}
		}{
			RecordID: recordID,
			Fields: map[string]interface{}{
				"数量": feishu.CreateNumberField(float64(100 + i*10)),
				"描述": feishu.CreateTextField(fmt.Sprintf("批量更新的产品 %d", i+1)),
			},
		})
	}

	err := client.BatchUpdateRecords(appToken, tableID, records)
	if err != nil {
		return err
	}

	return nil
}
