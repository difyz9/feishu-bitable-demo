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

// MarkdownDocsConfig 云文档配置结构
type MarkdownDocsConfig struct {
	Feishu struct {
		AppID       string `yaml:"app_id"`
		AppSecret   string `yaml:"app_secret"`
		FolderToken string `yaml:"folder_token"`
	} `yaml:"feishu"`
}

func main() {
	// 读取配置文件
	config, err := loadMarkdownDocsConfig("config.yaml")
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
	fmt.Println("📝 飞书云文档 Markdown 写入示例")
	fmt.Println("=================================================\n")

	// 1. 创建云文档
	fmt.Println("📝 步骤 1: 创建云文档")
	documentID, err := createMarkdownDocument(client, config.Feishu.FolderToken)
	if err != nil {
		log.Fatalf("❌ 创建云文档失败: %v", err)
	}
	fmt.Printf("✅ 成功创建云文档，Document ID: %s\n\n", documentID)

	time.Sleep(1 * time.Second)

	// 2. 写入 Markdown 格式的内容
	fmt.Println("📝 步骤 2: 写入 Markdown 格式内容")
	err = writeMarkdownContent(client, documentID)
	if err != nil {
		log.Fatalf("❌ 写入内容失败: %v", err)
	}
	fmt.Println("✅ 成功写入 Markdown 格式内容\n")

	time.Sleep(1 * time.Second)

	// 3. 验证写入的内容
	fmt.Println("📝 步骤 3: 验证文档内容")
	err = verifyDocumentContent(client, documentID)
	if err != nil {
		log.Fatalf("❌ 验证内容失败: %v", err)
	}
	fmt.Println("✅ 内容验证完成\n")

	fmt.Println("=================================================")
	fmt.Println("🎉 Markdown 内容写入完成！")
	fmt.Printf("📄 文档链接: https://example.feishu.cn/docx/%s\n", documentID)
	fmt.Println("=================================================")
	fmt.Println("\n💡 提示：")
	fmt.Println("  - 文档包含多种 Markdown 格式的内容")
	fmt.Println("  - 包括：普通文本、粗体、斜体、代码、链接等")
	fmt.Println("  - 你可以在飞书中查看渲染效果")
}

// createMarkdownDocument 创建用于 Markdown 演示的云文档
func createMarkdownDocument(client *feishu.MultiTableClient, folderToken string) (string, error) {
	title := fmt.Sprintf("Markdown 示例文档 - %s", time.Now().Format("2006-01-02 15:04:05"))

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

// writeMarkdownContent 写入各种 Markdown 格式的内容
func writeMarkdownContent(client *feishu.MultiTableClient, documentID string) error {
	// 获取页面块 ID
	blocksResp, err := client.ListDocumentBlocks(documentID)
	if err != nil {
		return err
	}
	if blocksResp.Data == nil || len(blocksResp.Data.Items) == 0 {
		return fmt.Errorf("无法获取文档页面块")
	}
	pageBlockID := *blocksResp.Data.Items[0].BlockId

	// 创建多个文本块，展示不同的 Markdown 格式
	blocks := []*larkdocx.Block{
		// 1. 标题和简介
		createTextBlock("# Markdown 格式示例\n\n这是一个展示如何在飞书云文档中写入 Markdown 格式内容的示例。"),
		
		// 2. 文本格式
		createTextBlock("\n## 文本格式\n\n**粗体文本** 和 *斜体文本* 以及 `代码文本`。\n\n你也可以使用 ~~删除线~~ 和 __下划线__。"),
		
		// 3. 列表示例
		createTextBlock("\n## 列表示例\n\n无序列表：\n- 第一项\n- 第二项\n  - 子项 2.1\n  - 子项 2.2\n- 第三项"),
		
		createTextBlock("\n有序列表：\n1. 第一步\n2. 第二步\n3. 第三步"),
		
		// 4. 链接和引用
		createTextBlock("\n## 链接和引用\n\n访问 [飞书开放平台](https://open.feishu.cn) 了解更多。\n\n> 这是一个引用块\n> 可以包含多行内容"),
		
		// 5. 代码块
		createTextBlock("\n## 代码示例\n\n内联代码：`fmt.Println(\"Hello World\")`\n\n代码块：\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, Feishu!\")\n}\n```"),
		
		// 6. 表格
		createTextBlock("\n## 表格\n\n| 功能 | 状态 | 说明 |\n|------|------|------|\n| 创建文档 | ✅ | 已实现 |\n| 写入内容 | ✅ | 已实现 |\n| 读取内容 | ✅ | 已实现 |"),
		
		// 7. 任务列表
		createTextBlock("\n## 任务清单\n\n- [x] 创建云文档\n- [x] 写入 Markdown 内容\n- [x] 验证文档内容\n- [ ] 添加更多功能"),
		
		// 8. 分隔线和结语
		createTextBlock("\n---\n\n📅 生成时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n\n🚀 Powered by feishu_golang 项目"),
	}

	// 批量写入所有块
	resp, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, blocks)
	if err != nil {
		return err
	}

	if resp.Data != nil && resp.Data.Children != nil {
		fmt.Printf("  📦 成功写入 %d 个内容块\n", len(resp.Data.Children))
	}

	return nil
}

// createTextBlock 创建文本块（复用 feishu 包的函数）
func createTextBlock(text string) *larkdocx.Block {
	return feishu.CreateTextBlock(text)
}

// verifyDocumentContent 验证文档内容
func verifyDocumentContent(client *feishu.MultiTableClient, documentID string) error {
	// 获取文档所有块
	blocksResp, err := client.ListDocumentBlocks(documentID)
	if err != nil {
		return err
	}

	if blocksResp.Data == nil {
		return fmt.Errorf("响应数据为空")
	}

	fmt.Printf("  📊 文档共有 %d 个块\n", len(blocksResp.Data.Items))

	// 获取文档纯文本内容
	contentResp, err := client.GetDocumentRawContent(documentID)
	if err != nil {
		return err
	}

	if contentResp.Data != nil && contentResp.Data.Content != nil {
		content := *contentResp.Data.Content
		fmt.Printf("  📝 文档内容长度: %d 字符\n", len(content))
		
		// 显示内容预览
		if len(content) > 300 {
			fmt.Printf("  📄 内容预览（前 300 字符）:\n%s...\n", content[:300])
		} else {
			fmt.Printf("  📄 完整内容:\n%s\n", content)
		}
	}

	return nil
}

// loadMarkdownDocsConfig 加载配置文件
func loadMarkdownDocsConfig(filename string) (*MarkdownDocsConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config MarkdownDocsConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
