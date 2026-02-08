package feishu

import (
	"context"
	"fmt"

	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
)

// Block 类型别名，方便使用
type Block = larkdocx.Block

// CreateDocument 创建云文档
func (c *MultiTableClient) CreateDocument(title, folderToken string) (*larkdocx.CreateDocumentResp, error) {
	// 构建请求
	reqBuilder := larkdocx.NewCreateDocumentReqBuilder()
	bodyBuilder := larkdocx.NewCreateDocumentReqBodyBuilder().Title(title)
	
	if folderToken != "" {
		bodyBuilder = bodyBuilder.FolderToken(folderToken)
	}
	
	req := reqBuilder.Body(bodyBuilder.Build()).Build()

	// 发起请求
	resp, err := c.client.Docx.Document.Create(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("创建云文档失败: %w", err)
	}

	// 检查响应
	if !resp.Success() {
		return nil, fmt.Errorf("创建云文档失败: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	return resp, nil
}

// GetDocument 获取云文档基本信息
func (c *MultiTableClient) GetDocument(documentID string) (*larkdocx.GetDocumentResp, error) {
	// 构建请求
	req := larkdocx.NewGetDocumentReqBuilder().
		DocumentId(documentID).
		Build()

	// 发起请求
	resp, err := c.client.Docx.Document.Get(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("获取云文档信息失败: %w", err)
	}

	// 检查响应
	if !resp.Success() {
		return nil, fmt.Errorf("获取云文档信息失败: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	return resp, nil
}

// GetDocumentRawContent 获取云文档纯文本内容
func (c *MultiTableClient) GetDocumentRawContent(documentID string) (*larkdocx.RawContentDocumentResp, error) {
	// 构建请求
	req := larkdocx.NewRawContentDocumentReqBuilder().
		DocumentId(documentID).
		Build()

	// 发起请求
	resp, err := c.client.Docx.Document.RawContent(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("获取云文档内容失败: %w", err)
	}

	// 检查响应
	if !resp.Success() {
		return nil, fmt.Errorf("获取云文档内容失败: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	return resp, nil
}

// ListDocumentBlocks 获取云文档所有块（blocks）
func (c *MultiTableClient) ListDocumentBlocks(documentID string) (*larkdocx.ListDocumentBlockResp, error) {
	// 构建请求
	req := larkdocx.NewListDocumentBlockReqBuilder().
		DocumentId(documentID).
		Build()

	// 发起请求
	resp, err := c.client.Docx.DocumentBlock.List(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("获取文档块列表失败: %w", err)
	}

	// 检查响应
	if !resp.Success() {
		return nil, fmt.Errorf("获取文档块列表失败: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	return resp, nil
}

// CreateTextBlock 创建文本块的辅助函数
func CreateTextBlock(text string) *larkdocx.Block {
	blockType := 2 // 2 表示文本块
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Text(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(text).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateDocumentBlock 在云文档中创建块（插入子块）
func (c *MultiTableClient) CreateDocumentBlock(documentID, blockID string, index int, children []*larkdocx.Block) (*larkdocx.CreateDocumentBlockChildrenResp, error) {
	// 构建请求
	req := larkdocx.NewCreateDocumentBlockChildrenReqBuilder().
		DocumentId(documentID).
		BlockId(blockID).
		Body(
			larkdocx.NewCreateDocumentBlockChildrenReqBodyBuilder().
				Children(children).
				Index(index).
				Build(),
		).
		Build()

	// 发起请求
	resp, err := c.client.Docx.DocumentBlockChildren.Create(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("创建文档块失败: %w", err)
	}

	// 检查响应
	if !resp.Success() {
		return nil, fmt.Errorf("创建文档块失败: code=%d, msg=%s", resp.Code, resp.Msg)
	}

	return resp, nil
}

// CreateTextParagraph 创建简单的文本段落（简化版）
func CreateTextParagraph(content string) string {
	// 返回简单的文本内容，用于演示
	return content
}

// UpdateDocumentBlock 更新文档块内容（通过追加新块的方式）
// 注意：由于 SDK API 限制，这里使用追加新块的方式来"更新"内容
func (c *MultiTableClient) UpdateDocumentBlock(documentID, parentBlockID string, newBlock *larkdocx.Block) error {
	// 创建新块
	_, err := c.CreateDocumentBlock(documentID, parentBlockID, -1, []*larkdocx.Block{newBlock})
	if err != nil {
		return fmt.Errorf("添加新块失败: %w", err)
	}
	
	return nil
}

// BuildRichTextElements 构建富文本元素数组（支持多种样式）
func BuildRichTextElements(segments []TextSegment) []*larkdocx.TextElement {
	elements := make([]*larkdocx.TextElement, 0, len(segments))
	
	for _, seg := range segments {
		element := larkdocx.NewTextElementBuilder().
			TextRun(larkdocx.NewTextRunBuilder().
				Content(seg.Content).
				TextElementStyle(seg.Style).
				Build()).
			Build()
		elements = append(elements, element)
	}
	
	return elements
}

// TextSegment 文本片段（带样式）
type TextSegment struct {
	Content string
	Style   *larkdocx.TextElementStyle
}

// NewTextSegment 创建普通文本片段
func NewTextSegment(content string) TextSegment {
	return TextSegment{
		Content: content,
		Style:   larkdocx.NewTextElementStyleBuilder().Build(),
	}
}

// NewBoldTextSegment 创建加粗文本片段
func NewBoldTextSegment(content string) TextSegment {
	return TextSegment{
		Content: content,
		Style: larkdocx.NewTextElementStyleBuilder().
			Bold(true).
			Build(),
	}
}

// NewItalicTextSegment 创建斜体文本片段
func NewItalicTextSegment(content string) TextSegment {
	return TextSegment{
		Content: content,
		Style: larkdocx.NewTextElementStyleBuilder().
			Italic(true).
			Build(),
	}
}

// NewStrikethroughTextSegment 创建删除线文本片段
func NewStrikethroughTextSegment(content string) TextSegment {
	return TextSegment{
		Content: content,
		Style: larkdocx.NewTextElementStyleBuilder().
			Strikethrough(true).
			Build(),
	}
}

// NewUnderlineTextSegment 创建下划线文本片段
func NewUnderlineTextSegment(content string) TextSegment {
	return TextSegment{
		Content: content,
		Style: larkdocx.NewTextElementStyleBuilder().
			Underline(true).
			Build(),
	}
}

// NewInlineCodeSegment 创建内联代码片段
func NewInlineCodeSegment(content string) TextSegment {
	return TextSegment{
		Content: content,
		Style: larkdocx.NewTextElementStyleBuilder().
			InlineCode(true).
			Build(),
	}
}

// NewColoredTextSegment 创建带颜色的文本片段
func NewColoredTextSegment(content string, textColor int, bgColor int) TextSegment {
	builder := larkdocx.NewTextElementStyleBuilder()
	
	if textColor > 0 {
		builder = builder.TextColor(textColor)
	}
	if bgColor > 0 {
		builder = builder.BackgroundColor(bgColor)
	}
	
	return TextSegment{
		Content: content,
		Style:   builder.Build(),
	}
}

// NewLinkSegment 创建链接文本片段
func NewLinkSegment(content, url string) TextSegment {
	return TextSegment{
		Content: content,
		Style: larkdocx.NewTextElementStyleBuilder().
			Link(larkdocx.NewLinkBuilder().
				Url(url).
				Build()).
			Build(),
	}
}

// CreateHeading1Block 创建一级标题块
func CreateHeading1Block(title string) *larkdocx.Block {
	blockType := 3 // 3 表示一级标题
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Heading1(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(title).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateHeading2Block 创建二级标题块
func CreateHeading2Block(title string) *larkdocx.Block {
	blockType := 4 // 4 表示二级标题
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Heading2(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(title).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateStyledTextBlock 创建带样式的文本块
func CreateStyledTextBlock(segments ...string) *larkdocx.Block {
	// 解析 segments 并创建文本元素（这里简化处理）
	// 实际使用时需要传入 TextSegment 数组
	blockType := 2 // 2 表示文本块
	elements := make([]*larkdocx.TextElement, 0)
	
	for _, seg := range segments {
		elements = append(elements, larkdocx.NewTextElementBuilder().
			TextRun(
				larkdocx.NewTextRunBuilder().
					Content(seg).
					Build(),
			).
			Build())
	}
	
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Text(
			larkdocx.NewTextBuilder().
				Elements(elements).
				Build(),
		).
		Build()
}

// CreateBulletBlock 创建无序列表项
func CreateBulletBlock(content string) *larkdocx.Block {
	blockType := 6 // 6 表示无序列表
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Bullet(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(content).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateOrderedBlock 创建有序列表项
func CreateOrderedBlock(content string, number int) *larkdocx.Block {
	blockType := 7 // 7 表示有序列表
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Ordered(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(content).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateCodeBlock 创建代码块
func CreateCodeBlock(code string, language int) *larkdocx.Block {
	blockType := 8 // 8 表示代码块
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Code(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(code).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateQuoteBlock 创建引用块
func CreateQuoteBlock(content string) *larkdocx.Block {
	blockType := 10 // 10 表示引用块
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Quote(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(content).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}

// CreateTodoBlock 创建待办事项
func CreateTodoBlock(content string, checked bool) *larkdocx.Block {
	blockType := 17 // 17 表示待办事项
	return larkdocx.NewBlockBuilder().
		BlockType(blockType).
		Todo(
			larkdocx.NewTextBuilder().
				Elements([]*larkdocx.TextElement{
					larkdocx.NewTextElementBuilder().
						TextRun(
							larkdocx.NewTextRunBuilder().
								Content(content).
								Build(),
						).
						Build(),
				}).
				Build(),
		).
		Build()
}


// 辅助函数：创建不同样式的文本片段（用于 CreateStyledTextBlock）
func PlainText(content string) string {
	return content
}

func BoldText(content string) string {
	// 这里简化处理，实际应该返回带样式的元素
	return content
}

func ItalicText(content string) string {
	return content
}

func UnderlineText(content string) string {
	return content
}

func StrikethroughText(content string) string {
	return content
}

func InlineCodeText(content string) string {
	return content
}

func ColoredText(content string, color int) string {
	return content
}

func LinkText(content, url string) string {
	return content
}

