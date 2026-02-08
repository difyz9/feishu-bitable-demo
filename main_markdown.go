package main

import (
	"fmt"
	"log"
	"time"
	
	"feishu_golang/feishu"
	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
)

func main() {
	// 初始化客户端
	client, err := feishu.NewMultiTableClient("config.yaml")
	if err != nil {
		log.Fatalf("❌ 初始化客户端失败: %v", err)
	}
	
	fmt.Println("=== 飞书云文档 Markdown 富文本写入示例 ===\n")
	
	// 1. 创建新文档
	docID, err := createMarkdownDocument(client)
	if err != nil {
		log.Fatalf("❌ 创建文档失败: %v", err)
	}
	
	// 等待文档初始化
	time.Sleep(2 * time.Second)
	
	// 2. 写入各种 Markdown 样式内容
	if err := writeMarkdownContent(client, docID); err != nil {
		log.Fatalf("❌ 写入内容失败: %v", err)
	}
	
	fmt.Println("\n✅ 所有操作完成！")
	fmt.Printf("📄 文档 ID: %s\n", docID)
}

// createMarkdownDocument 创建一个用于演示 Markdown 的文档
func createMarkdownDocument(client *feishu.MultiTableClient) (string, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	title := fmt.Sprintf("Markdown 富文本示例 - %s", timestamp)
	
	fmt.Printf("📝 创建文档: %s\n", title)
	
	resp, err := client.CreateDocument(title, "")
	if err != nil {
		return "", err
	}
	
	docID := *resp.Data.Document.DocumentId
	fmt.Printf("✅ 文档创建成功，ID: %s\n\n", docID)
	
	return docID, nil
}

// writeMarkdownContent 写入各种 Markdown 样式内容
func writeMarkdownContent(client *feishu.MultiTableClient, docID string) error {
	// 获取文档信息以获取 page block ID
	docInfo, err := client.GetDocument(docID)
	if err != nil {
		return fmt.Errorf("获取文档信息失败: %w", err)
	}
	pageID := *docInfo.Data.Document.DocumentId
	
	fmt.Println("📋 开始写入各种 Markdown 样式内容...\n")
	
	// 1. 插入一级标题
	if err := insertHeading(client, docID, pageID, 1, "一、文本样式演示"); err != nil {
		return err
	}
	
	// 2. 插入混合样式的段落
	if err := insertMixedStyleParagraph(client, docID, pageID); err != nil {
		return err
	}
	
	// 3. 插入二级标题
	if err := insertHeading(client, docID, pageID, 2, "二、代码和链接"); err != nil {
		return err
	}
	
	// 4. 插入带代码和链接的段落
	if err := insertCodeAndLinkParagraph(client, docID, pageID); err != nil {
		return err
	}
	
	// 5. 插入三级标题
	if err := insertHeading(client, docID, pageID, 3, "三、彩色文本"); err != nil {
		return err
	}
	
	// 6. 插入彩色文本段落
	if err := insertColoredParagraph(client, docID, pageID); err != nil {
		return err
	}
	
	// 7. 插入代码块
	if err := insertCodeBlock(client, docID, pageID); err != nil {
		return err
	}
	
	// 8. 插入引用块
	if err := insertQuoteBlock(client, docID, pageID); err != nil {
		return err
	}
	
	// 9. 插入无序列表
	if err := insertBulletList(client, docID, pageID); err != nil {
		return err
	}
	
	// 10. 插入有序列表
	if err := insertOrderedList(client, docID, pageID); err != nil {
		return err
	}
	
	return nil
}

// insertHeading 插入标题块
func insertHeading(client *feishu.MultiTableClient, docID, parentID string, level int, content string) error {
	fmt.Printf("  📌 插入 H%d 标题: %s\n", level, content)
	
	headingBlock := feishu.CreateHeadingBlock(level, content)
	_, err := client.CreateDocumentBlock(docID, parentID, headingBlock)
	
	if err != nil {
		return fmt.Errorf("插入标题失败: %w", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	return nil
}

// insertMixedStyleParagraph 插入混合样式段落
func insertMixedStyleParagraph(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  📝 插入混合样式段落（加粗、斜体、删除线、下划线）")
	
	// 先插入一个空文本块
	textBlock := feishu.CreateTextBlock("")
	resp, err := client.CreateDocumentBlock(docID, parentID, textBlock)
	if err != nil {
		return fmt.Errorf("插入文本块失败: %w", err)
	}
	
	// 获取新创建的块 ID
	if len(resp.Data.Children) == 0 {
		return fmt.Errorf("未能获取新创建的块 ID")
	}
	blockID := *resp.Data.Children[0].BlockId
	
	time.Sleep(500 * time.Millisecond)
	
	// 构建富文本内容：这是普通文本，加粗文本，斜体文本，删除线，下划线。
	segments := []feishu.TextSegment{
		feishu.NewTextSegment("这是普通文本，"),
		feishu.NewBoldTextSegment("加粗文本"),
		feishu.NewTextSegment("，"),
		feishu.NewItalicTextSegment("斜体文本"),
		feishu.NewTextSegment("，"),
		feishu.NewStrikethroughTextSegment("删除线"),
		feishu.NewTextSegment("，"),
		feishu.NewUnderlineTextSegment("下划线"),
		feishu.NewTextSegment("。"),
	}
	
	elements := feishu.BuildRichTextElements(segments)
	
	// 更新块内容
	updateFields := map[string]interface{}{
		"elements": elements,
	}
	
	_, err = client.UpdateDocumentBlock(docID, blockID, updateFields)
	if err != nil {
		return fmt.Errorf("更新块内容失败: %w", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	return nil
}

// insertCodeAndLinkParagraph 插入带代码和链接的段落
func insertCodeAndLinkParagraph(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  📝 插入代码和链接段落")
	
	// 先插入一个空文本块
	textBlock := feishu.CreateTextBlock("")
	resp, err := client.CreateDocumentBlock(docID, parentID, textBlock)
	if err != nil {
		return fmt.Errorf("插入文本块失败: %w", err)
	}
	
	blockID := *resp.Data.Children[0].BlockId
	time.Sleep(500 * time.Millisecond)
	
	// 构建富文本内容：使用 fmt.Println() 函数打印，详见官方文档。
	segments := []feishu.TextSegment{
		feishu.NewTextSegment("使用 "),
		feishu.NewInlineCodeSegment("fmt.Println()"),
		feishu.NewTextSegment(" 函数打印，详见"),
		feishu.NewLinkSegment("官方文档", "https://pkg.go.dev/fmt"),
		feishu.NewTextSegment("。"),
	}
	
	elements := feishu.BuildRichTextElements(segments)
	
	updateFields := map[string]interface{}{
		"elements": elements,
	}
	
	_, err = client.UpdateDocumentBlock(docID, blockID, updateFields)
	if err != nil {
		return fmt.Errorf("更新块内容失败: %w", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	return nil
}

// insertColoredParagraph 插入彩色文本段落
func insertColoredParagraph(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  🎨 插入彩色文本段落")
	
	// 先插入一个空文本块
	textBlock := feishu.CreateTextBlock("")
	resp, err := client.CreateDocumentBlock(docID, parentID, textBlock)
	if err != nil {
		return fmt.Errorf("插入文本块失败: %w", err)
	}
	
	blockID := *resp.Data.Children[0].BlockId
	time.Sleep(500 * time.Millisecond)
	
	// 构建富文本内容：红色文字、蓝色背景、绿色高亮
	// 颜色值参考：1-红 2-橙 3-黄 4-绿 5-蓝 6-紫 7-粉 8-灰
	segments := []feishu.TextSegment{
		feishu.NewColoredTextSegment("红色文字", 1, 0),
		feishu.NewTextSegment("，"),
		feishu.NewColoredTextSegment("蓝色背景", 0, 5),
		feishu.NewTextSegment("，"),
		feishu.NewColoredTextSegment("绿色高亮", 4, 4),
		feishu.NewTextSegment("。"),
	}
	
	elements := feishu.BuildRichTextElements(segments)
	
	updateFields := map[string]interface{}{
		"elements": elements,
	}
	
	_, err = client.UpdateDocumentBlock(docID, blockID, updateFields)
	if err != nil {
		return fmt.Errorf("更新块内容失败: %w", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	return nil
}

// insertCodeBlock 插入代码块
func insertCodeBlock(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  💻 插入代码块")
	
	code := `func main() {
    fmt.Println("Hello, 飞书!")
}`
	
	codeBlock := feishu.CreateCodeBlock("go", code)
	_, err := client.CreateDocumentBlock(docID, parentID, codeBlock)
	
	if err != nil {
		return fmt.Errorf("插入代码块失败: %w", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	return nil
}

// insertQuoteBlock 插入引用块
func insertQuoteBlock(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  📖 插入引用块")
	
	// 引用块需要使用特定的 block 类型
	children := []*larkdocx.Block{
		larkdocx.NewBlockBuilder().
			BlockType(21). // 21 = Quote block
			Quote(larkdocx.NewQuoteBuilder().
				Children([]*larkdocx.Block{
					larkdocx.NewBlockBuilder().
						BlockType(2). // 2 = Text block
						Text(larkdocx.NewTextBuilder().
							Elements([]*larkdocx.TextElement{
								larkdocx.NewTextElementBuilder().
									TextRun(larkdocx.NewTextRunBuilder().
										Content("这是一段引用文字，用于强调重要信息。").
										Build()).
									Build(),
							}).
							Build()).
						Build(),
				}).
				Build()).
			Build(),
	}
	
	_, err := client.CreateDocumentBlock(docID, parentID, children)
	if err != nil {
		return fmt.Errorf("插入引用块失败: %w", err)
	}
	
	time.Sleep(500 * time.Millisecond)
	return nil
}

// insertBulletList 插入无序列表
func insertBulletList(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  🔹 插入无序列表")
	
	items := []string{
		"无序列表项 1",
		"无序列表项 2",
		"无序列表项 3",
	}
	
	for _, item := range items {
		// Bullet list block type = 3
		bulletBlock := []*larkdocx.Block{
			larkdocx.NewBlockBuilder().
				BlockType(3).
				Bullet(larkdocx.NewBulletBuilder().
					Children([]*larkdocx.Block{
						larkdocx.NewBlockBuilder().
							BlockType(2). // Text block
							Text(larkdocx.NewTextBuilder().
								Elements([]*larkdocx.TextElement{
									larkdocx.NewTextElementBuilder().
										TextRun(larkdocx.NewTextRunBuilder().
											Content(item).
											Build()).
										Build(),
								}).
								Build()).
							Build(),
					}).
					Build()).
				Build(),
		}
		
		_, err := client.CreateDocumentBlock(docID, parentID, bulletBlock)
		if err != nil {
			return fmt.Errorf("插入无序列表失败: %w", err)
		}
		
		time.Sleep(300 * time.Millisecond)
	}
	
	return nil
}

// insertOrderedList 插入有序列表
func insertOrderedList(client *feishu.MultiTableClient, docID, parentID string) error {
	fmt.Println("  🔢 插入有序列表")
	
	items := []string{
		"有序列表项 1",
		"有序列表项 2",
		"有序列表项 3",
	}
	
	for _, item := range items {
		// Ordered list block type = 4
		orderedBlock := []*larkdocx.Block{
			larkdocx.NewBlockBuilder().
				BlockType(4).
				Ordered(larkdocx.NewOrderedBuilder().
					Children([]*larkdocx.Block{
						larkdocx.NewBlockBuilder().
							BlockType(2). // Text block
							Text(larkdocx.NewTextBuilder().
								Elements([]*larkdocx.TextElement{
									larkdocx.NewTextElementBuilder().
										TextRun(larkdocx.NewTextRunBuilder().
											Content(item).
											Build()).
										Build(),
								}).
								Build()).
							Build(),
					}).
					Build()).
				Build(),
		}
		
		_, err := client.CreateDocumentBlock(docID, parentID, orderedBlock)
		if err != nil {
			return fmt.Errorf("插入有序列表失败: %w", err)
		}
		
		time.Sleep(300 * time.Millisecond)
	}
	
	return nil
}
