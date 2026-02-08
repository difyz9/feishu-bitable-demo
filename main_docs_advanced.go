package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"feishu_bitable_demo/feishu"
	"gopkg.in/yaml.v3"
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
	fmt.Println("=================================================")
	fmt.Println("🚀 飞书云文档高级 Markdown 样式示例")
	fmt.Println("=================================================\n")

	// 读取配置文件
	config, err := loadDocsConfig("config.yaml")
	if err != nil {
		log.Fatalf("❌ 读取配置文件失败: %v", err)
	}

	// 验证配置
	if config.Feishu.AppID == "your_app_id_here" || config.Feishu.AppSecret == "your_app_secret_here" {
		log.Fatalf("❌ 请先在 config.yaml 中填写正确的飞书应用配置")
	}

	// 初始化客户端
	client := feishu.NewMultiTableClient(config.Feishu.AppID, config.Feishu.AppSecret)

	// 步骤 1: 创建文档
	fmt.Println("📝 步骤 1: 创建云文档")
	documentTitle := fmt.Sprintf("Markdown 样式示例 - %s", time.Now().Format("2006-01-02 15:04:05"))
	resp, err := client.CreateDocument(documentTitle, "")
	if err != nil {
		log.Fatalf("❌ 创建文档失败: %v", err)
	}
	if resp.Data == nil || resp.Data.Document == nil {
		log.Fatalf("❌ 创建文档响应数据为空")
	}
	documentID := *resp.Data.Document.DocumentId
	fmt.Printf("  📄 文档标题: %s\n", documentTitle)
	fmt.Printf("  📋 文档 ID: %s\n", documentID)
	fmt.Println("✅ 成功创建云文档\n")

	// 获取页面块 ID
	blocksResp, err := client.ListDocumentBlocks(documentID)
	if err != nil {
		log.Fatalf("❌ 获取文档块失败: %v", err)
	}
	if blocksResp.Data == nil || len(blocksResp.Data.Items) == 0 {
		log.Fatalf("❌ 无法获取文档页面块")
	}
	pageBlockID := *blocksResp.Data.Items[0].BlockId

	// 步骤 2: 写入各种块类型
	fmt.Println("📝 步骤 2: 写入各种 Markdown 块到文档")
	
	// 2.1 写入一级标题
	fmt.Println("  ✏️  2.1: 写入一级标题")
	err = writeHeading1(client, documentID, pageBlockID, "📚 Markdown 样式完整指南")
	if err != nil {
		log.Printf("⚠️ 写入标题失败: %v\n", err)
	}

	// 2.2 写入普通文本（带样式）
	fmt.Println("  ✏️  2.2: 写入带样式的文本")
	err = writeStyledText(client, documentID, pageBlockID)
	if err != nil {
		log.Printf("⚠️ 写入样式文本失败: %v\n", err)
	}

	// 2.3 写入二级标题
	fmt.Println("  ✏️  2.3: 写入二级标题")
	err = writeHeading2(client, documentID, pageBlockID, "🎨 文本样式示例")
	if err != nil {
		log.Printf("⚠️ 写入二级标题失败: %v\n", err)
	}

	// 2.4 写入无序列表
	fmt.Println("  ✏️  2.4: 写入无序列表")
	err = writeBulletList(client, documentID, pageBlockID)
	if err != nil {
		log.Printf("⚠️ 写入无序列表失败: %v\n", err)
	}

	// 2.5 写入有序列表
	fmt.Println("  ✏️  2.5: 写入有序列表")
	err = writeOrderedList(client, documentID, pageBlockID)
	if err != nil {
		log.Printf("⚠️ 写入有序列表失败: %v\n", err)
	}

	// 2.6 写入代码块
	fmt.Println("  ✏️  2.6: 写入代码块")
	err = writeCodeBlock(client, documentID, pageBlockID)
	if err != nil {
		log.Printf("⚠️ 写入代码块失败: %v\n", err)
	}

	// 2.7 写入引用块
	fmt.Println("  ✏️  2.7: 写入引用块")
	err = writeQuote(client, documentID, pageBlockID)
	if err != nil {
		log.Printf("⚠️ 写入引用块失败: %v\n", err)
	}

	// 2.8 写入待办事项
	fmt.Println("  ✏️  2.8: 写入待办事项")
	err = writeTodoList(client, documentID, pageBlockID)
	if err != nil {
		log.Printf("⚠️ 写入待办事项失败: %v\n", err)
	}

	fmt.Println("✅ 所有 Markdown 块已成功写入\n")

	// 步骤 3: 获取并验证文档内容
	fmt.Println("📝 步骤 3: 验证文档内容")
	contentResp, err := client.GetDocumentRawContent(documentID)
	if err != nil {
		log.Printf("⚠️ 获取文档内容失败: %v\n", err)
	} else if contentResp.Data != nil && contentResp.Data.Content != nil {
		rawText := *contentResp.Data.Content
		fmt.Printf("  📄 文档包含 %d 个字符\n", len(rawText))
		fmt.Println("✅ 文档内容验证成功\n")
	}

	// 输出文档链接
	fmt.Println("=================================================")
	fmt.Println("🎉 Markdown 样式示例完成！")
	fmt.Printf("📄 文档链接: https://example.feishu.cn/docx/%s\n", documentID)
	fmt.Println("=================================================\n")

	fmt.Println("💡 本示例展示了以下 Markdown 样式：")
	fmt.Println("  ✅ 标题（H1, H2）")
	fmt.Println("  ✅ 文本样式（加粗、斜体、下划线、删除线、行内代码）")
	fmt.Println("  ✅ 文本颜色和背景色")
	fmt.Println("  ✅ 超链接")
	fmt.Println("  ✅ 无序列表")
	fmt.Println("  ✅ 有序列表")
	fmt.Println("  ✅ 代码块（支持语法高亮）")
	fmt.Println("  ✅ 引用块")
	fmt.Println("  ✅ 待办事项")
}

// writeHeading1 写入一级标题
func writeHeading1(client *feishu.MultiTableClient, documentID, pageBlockID, title string) error {
	heading1Block := feishu.CreateHeading1Block(title)
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, []*feishu.Block{heading1Block})
	return err
}

// writeHeading2 写入二级标题
func writeHeading2(client *feishu.MultiTableClient, documentID, pageBlockID, title string) error {
	heading2Block := feishu.CreateHeading2Block(title)
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, []*feishu.Block{heading2Block})
	return err
}

// writeStyledText 写入带样式的文本
func writeStyledText(client *feishu.MultiTableClient, documentID, pageBlockID string) error {
	// 创建一个包含多种样式的文本块
	styledBlock := feishu.CreateStyledTextBlock(
		"这是一段演示文本，包含 ",
		feishu.BoldText("加粗"),
		feishu.PlainText("、"),
		feishu.ItalicText("斜体"),
		feishu.PlainText("、"),
		feishu.UnderlineText("下划线"),
		feishu.PlainText("、"),
		feishu.StrikethroughText("删除线"),
		feishu.PlainText("、"),
		feishu.InlineCodeText("代码"),
		feishu.PlainText("、"),
		feishu.ColoredText("红色文本", 1),
		feishu.PlainText(" 和 "),
		feishu.LinkText("超链接", "https://open.feishu.cn"),
		feishu.PlainText("。"),
	)
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, []*feishu.Block{styledBlock})
	return err
}

// writeBulletList 写入无序列表
func writeBulletList(client *feishu.MultiTableClient, documentID, pageBlockID string) error {
	bullets := []*feishu.Block{
		feishu.CreateBulletBlock("第一个列表项"),
		feishu.CreateBulletBlock("第二个列表项"),
		feishu.CreateBulletBlock("第三个列表项"),
	}
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, bullets)
	return err
}

// writeOrderedList 写入有序列表
func writeOrderedList(client *feishu.MultiTableClient, documentID, pageBlockID string) error {
	orderedItems := []*feishu.Block{
		feishu.CreateOrderedBlock("第一步：初始化客户端", 1),
		feishu.CreateOrderedBlock("第二步：创建文档", 2),
		feishu.CreateOrderedBlock("第三步：写入内容", 3),
	}
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, orderedItems)
	return err
}

// writeCodeBlock 写入代码块
func writeCodeBlock(client *feishu.MultiTableClient, documentID, pageBlockID string) error {
	code := `package main

import "fmt"

func main() {
    fmt.Println("Hello, Feishu!")
}`
	codeBlock := feishu.CreateCodeBlock(code, 22) // 22 = Go
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, []*feishu.Block{codeBlock})
	return err
}

// writeQuote 写入引用块
func writeQuote(client *feishu.MultiTableClient, documentID, pageBlockID string) error {
	quoteBlock := feishu.CreateQuoteBlock("这是一段重要的引用内容，用于强调或引述。")
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, []*feishu.Block{quoteBlock})
	return err
}

// writeTodoList 写入待办事项
func writeTodoList(client *feishu.MultiTableClient, documentID, pageBlockID string) error {
	todos := []*feishu.Block{
		feishu.CreateTodoBlock("完成文档编写", false),
		feishu.CreateTodoBlock("代码审查", false),
		feishu.CreateTodoBlock("部署上线", false),
	}
	_, err := client.CreateDocumentBlock(documentID, pageBlockID, -1, todos)
	return err
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
