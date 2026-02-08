
package main

import (
	"fmt"
	"log"
	"time"

	"feishu_bitable_demo/feishu"
	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"

	"gopkg.in/yaml.v3"
	"os"
)

// DocsConfig 云文档配置结构
type DocsConfig struct {
	Feishu struct {
		AppID       string `yaml:"app_id"`
		AppSecret   string `yaml:"app_secret"`
		FolderToken string `yaml:"folder_token"` // 文件夹 Token（可选）
	} `yaml:"feishu"`
}

func main() {
	// 读取配置文件
	config, err := loadDocsConfig("config.yaml")
	if err != nil {
		log.Fatalf("❌ 读取配置文件失败: %v", err)
	}

	// 验证配置
	if config.Feishu.AppID == "your_app_id_here" || config.Feishu.AppSecret == "your_app_secret_here" {
		log.Fatalf("❌ 请先在 config.yaml 中填写正确的飞书应用配置")
	}

	// 创建飞书客户端
	client := feishu.NewMultiTableClient(config.Feishu.AppID, config.Feishu.AppSecret)

	fmt.Println("=================================================")
	fmt.Println("🚀 飞书云文档操作示例程序")
	fmt.Println("=================================================\n")


	// 1. 创建云文档
	fmt.Println("📝 步骤 1: 创建云文档")
	documentID, err := testCreateDocument(client, config.Feishu.FolderToken)
	if err != nil {
		log.Fatalf("❌ 创建云文档失败: %v", err)
	}
	fmt.Printf("✅ 成功创建云文档，Document ID: %s\n\n", documentID)

	// 等待一秒，确保文档已创建
	time.Sleep(1 * time.Second)

	// 写入自定义信息到文档
	fmt.Println("📝 步骤 1.1: 写入自定义信息到云文档")
	err = writeInfoToDocument(client, documentID, "本云文档由 feishu_golang 项目自动生成，演示云文档写入功能。\n\n可在此处记录项目说明、操作日志、或其他自定义内容。\n\n-- Powered by GitHub Copilot")
	if err != nil {
		log.Printf("⚠️ 写入信息失败: %v\n", err)
		fmt.Println("  💡 请检查应用是否有云文档编辑权限\n")
	} else {
		fmt.Println("✅ 已成功写入自定义信息到云文档\n")
	}

	// 2. 获取云文档信息
	fmt.Println("📝 步骤 2: 获取云文档信息")
	err = testGetDocument(client, documentID)
	if err != nil {
		log.Fatalf("❌ 获取云文档信息失败: %v", err)
	}
	fmt.Println("✅ 成功获取云文档信息\n")

	// 3. 获取云文档所有块
	fmt.Println("📝 步骤 3: 获取云文档所有块")
	blockIDs, err := testListDocumentBlocks(client, documentID)
	if err != nil {
		log.Fatalf("❌ 获取文档块失败: %v", err)
	}
	fmt.Printf("✅ 成功获取文档块，共 %d 个块\n\n", len(blockIDs))

	// 4. 获取云文档纯文本内容
	fmt.Println("📝 步骤 4: 获取云文档纯文本内容")
	err = testGetDocumentRawContent(client, documentID)
	if err != nil {
		log.Fatalf("❌ 获取文档内容失败: %v", err)
	}
	fmt.Println("✅ 成功获取文档内容\n")

	fmt.Println("=================================================")
	fmt.Println("🎉 所有云文档操作测试完成！")
	fmt.Printf("📄 文档链接: https://example.feishu.cn/docx/%s\n", documentID)
	fmt.Println("=================================================")
	fmt.Println("\n💡 提示：")
	fmt.Println("  - 云文档已成功创建，你可以在飞书中查看和编辑")
	fmt.Println("  - 已实现文本块写入功能，可以自动向文档添加内容")
	fmt.Println("  - 更多文档块类型（标题、列表等）请参考飞书开放平台文档")
}

// testCreateDocument 测试创建云文档
func testCreateDocument(client *feishu.MultiTableClient, folderToken string) (string, error) {
	title := fmt.Sprintf("测试云文档 - %s", time.Now().Format("2006-01-02 15:04:05"))

	resp, err := client.CreateDocument(title, folderToken)
	if err != nil {
		return "", err
	}

	if resp.Data == nil || resp.Data.Document == nil {
		return "", fmt.Errorf("响应数据为空")
	}

	documentID := *resp.Data.Document.DocumentId
	fmt.Printf("  📄 文档标题: %s\n", title)
	fmt.Printf("  📋 文档 ID: %s\n", documentID)

	return documentID, nil
}

// writeInfoToDocument 写入一段信息到云文档（插入文本块）
func writeInfoToDocument(client *feishu.MultiTableClient, documentID string, info string) error {
	// 获取页面块 ID（通常是第一个块）
	blocksResp, err := client.ListDocumentBlocks(documentID)
	if err != nil {
		return err
	}
	if blocksResp.Data == nil || len(blocksResp.Data.Items) == 0 {
		return fmt.Errorf("无法获取文档页面块")
	}
	pageBlockID := *blocksResp.Data.Items[0].BlockId

	// 构建文本块
	textBlock := feishu.CreateTextBlock(info)
	
	// 插入文本块到页面块
	resp, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, []*larkdocx.Block{textBlock})
	if err != nil {
		return err
	}
	
	if resp.Data != nil && resp.Data.Children != nil {
		fmt.Printf("  📦 成功写入 %d 个文本块\n", len(resp.Data.Children))
	}
	
	return nil
}

// testGetDocument 测试获取云文档信息
func testGetDocument(client *feishu.MultiTableClient, documentID string) error {
	resp, err := client.GetDocument(documentID)
	if err != nil {
		return err
	}

	if resp.Data == nil || resp.Data.Document == nil {
		return fmt.Errorf("响应数据为空")
	}

	doc := resp.Data.Document
	fmt.Printf("  📄 文档标题: %s\n", *doc.Title)
	fmt.Printf("  📋 文档 ID: %s\n", *doc.DocumentId)

	return nil
}

// testListDocumentBlocks 测试获取文档所有块
func testListDocumentBlocks(client *feishu.MultiTableClient, documentID string) ([]string, error) {
	resp, err := client.ListDocumentBlocks(documentID)
	if err != nil {
		return nil, err
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("响应数据为空")
	}

	var blockIDs []string
	for _, block := range resp.Data.Items {
		if block.BlockId != nil {
			blockIDs = append(blockIDs, *block.BlockId)
			fmt.Printf("  📦 块 ID: %s, 类型: %d\n", *block.BlockId, *block.BlockType)
		}
	}

	return blockIDs, nil
}

// testGetDocumentRawContent 测试获取文档纯文本内容
func testGetDocumentRawContent(client *feishu.MultiTableClient, documentID string) error {
	resp, err := client.GetDocumentRawContent(documentID)
	if err != nil {
		return err
	}

	if resp.Data == nil || resp.Data.Content == nil {
		return fmt.Errorf("响应数据为空")
	}

	content := *resp.Data.Content
	fmt.Printf("  📄 文档内容（原始）:\n")
	if len(content) > 200 {
		fmt.Printf("  %s...\n", content[:200])
	} else {
		fmt.Printf("  %s\n", content)
	}

	return nil
}

// loadDocsConfig 加载配置文件
func loadDocsConfig(filename string) (*DocsConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config DocsConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

