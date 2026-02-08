# 飞书云文档操作指南

## 📖 简介

本项目提供了完整的飞书云文档（Feishu Docs）操作功能，支持创建文档、添加内容、读取信息等操作。

## ✨ 功能特性

- ✅ **创建云文档** - 创建新的云文档
- ✅ **获取文档信息** - 读取文档基本信息
- ✅ **添加内容** - 支持多种内容块类型
- ✅ **读取内容** - 获取文档的所有块或纯文本内容
- ✅ **更新内容** - 修改文档块内容
- ✅ **删除内容** - 删除指定的文档块

## 📦 支持的内容块类型

| 块类型 | 说明 | 辅助函数 |
|--------|------|----------|
| 文本块 | 普通文本段落 | `CreateTextBlock(text)` |
| 标题块 | H1-H9 标题 | `CreateHeadingBlock(text, level)` |
| 无序列表 | 项目符号列表 | `CreateBulletListBlock(text)` |
| 有序列表 | 数字编号列表 | `CreateOrderedListBlock(text)` |
| 代码块 | 代码片段（支持语法高亮） | `CreateCodeBlock(code, language)` |

## 🔧 前置准备

### 1. 配置飞书应用权限

在飞书开放平台为你的应用添加以下权限：

#### 必需权限
- ✅ `docx:document` - 查看、编辑、创建和删除云文档
- ✅ `docx:document:readonly` - 查看云文档

#### 可选权限（推荐）
- `drive:drive` - 查看、评论、编辑和管理云空间中所有文件（用于指定文件夹创建文档）

### 2. 配置文件

在 `config.yaml` 中填写配置：

```yaml
feishu:
  app_id: "cli_xxxxxxxxxxxxx"
  app_secret: "xxxxxxxxxxxxxxxxxxxxx"
  folder_token: ""  # 可选，指定创建文档的文件夹
```

#### 如何获取 folder_token（可选）

1. 在飞书中打开目标文件夹
2. 从 URL 中获取 folder_token
   ```
   https://xxx.feishu.cn/drive/folder/fldcnxxxxxx
                                      ↑↑↑↑↑↑↑↑↑↑↑
                                      folder_token
   ```

如果不填写 `folder_token`，文档将创建在根目录下。

## 🚀 快速开始

### 运行云文档示例

```bash
# 方式一：直接运行
go run main_docs.go

# 方式二：使用运行脚本
./run_docs.sh
```

### 预期输出

```
=================================================
🚀 飞书云文档操作示例程序
=================================================

📝 步骤 1: 创建云文档
  📄 文档标题: 测试云文档 - 2026-02-08 10:30:00
  📋 文档 ID: doxcnxxxxxxxxxxxxxx
✅ 成功创建云文档，Document ID: doxcnxxxxxxxxxxxxxx

📝 步骤 2: 获取云文档信息
  📄 文档标题: 测试云文档 - 2026-02-08 10:30:00
  📋 文档 ID: doxcnxxxxxxxxxxxxxx
✅ 成功获取云文档信息

📝 步骤 3: 添加内容到云文档
  📦 成功添加 13 个内容块
✅ 成功添加内容到云文档

📝 步骤 4: 获取云文档所有块
  📦 块 ID: doxcnxxxxxx, 类型: 1
  📦 块 ID: doxcnxxxxxx, 类型: 3
  ...
✅ 成功获取文档块，共 15 个块

📝 步骤 5: 获取云文档纯文本内容
  📄 文档内容预览（前200字符）:
  📚 飞书云文档示例
这是一个由 Go SDK 创建的飞书云文档示例。
✨ 功能特性
...
✅ 成功获取文档内容

=================================================
🎉 所有云文档操作测试完成！
📄 文档链接: https://example.feishu.cn/docx/doxcnxxxxxxxxxxxxxx
=================================================
```

## 💡 代码示例

### 1. 创建云文档

```go
package main

import (
    "fmt"
    "feishu_bitable_demo/feishu"
)

func main() {
    // 初始化客户端
    client := feishu.NewMultiTableClient("your_app_id", "your_app_secret")
    
    // 创建云文档
    resp, err := client.CreateDocument("我的第一个文档", "")
    if err != nil {
        panic(err)
    }
    
    documentID := *resp.Data.Document.DocumentId
    fmt.Printf("文档 ID: %s\n", documentID)
}
```

### 2. 添加内容到文档

```go
// 获取文档的根块 ID
blocksResp, err := client.ListDocumentBlocks(documentID)
if err != nil {
    panic(err)
}
pageBlockID := *blocksResp.Data.Items[0].BlockId

// 创建不同类型的内容块
blocks := []*feishu.Block{
    feishu.CreateHeadingBlock("欢迎使用飞书云文档", 1),
    feishu.CreateTextBlock("这是一个文本段落。"),
    feishu.CreateBulletListBlock("第一个列表项"),
    feishu.CreateBulletListBlock("第二个列表项"),
    feishu.CreateCodeBlock("fmt.Println(\"Hello World\")", "go"),
}

// 添加到文档
resp, err := client.CreateDocumentBlock(documentID, pageBlockID, 2, blocks)
if err != nil {
    panic(err)
}
```

### 3. 读取文档内容

```go
// 获取文档信息
docResp, err := client.GetDocument(documentID)
if err != nil {
    panic(err)
}
fmt.Printf("文档标题: %s\n", *docResp.Data.Document.Title)

// 获取纯文本内容
contentResp, err := client.GetDocumentRawContent(documentID)
if err != nil {
    panic(err)
}
fmt.Printf("内容: %s\n", *contentResp.Data.Content)
```

### 4. 获取文档所有块

```go
resp, err := client.ListDocumentBlocks(documentID)
if err != nil {
    panic(err)
}

for _, block := range resp.Data.Items {
    fmt.Printf("块 ID: %s, 类型: %d\n", *block.BlockId, *block.BlockType)
}
```

### 5. 删除文档块

```go
blockIDs := []string{"doxcnxxxxxx", "doxcnxxxxxx"}
resp, err := client.DeleteDocumentBlocks(documentID, blockIDs)
if err != nil {
    panic(err)
}
fmt.Println("删除成功")
```

## 📋 API 参考

### 客户端方法

| 方法 | 说明 | 参数 |
|------|------|------|
| `CreateDocument(title, folderToken)` | 创建云文档 | title: 文档标题<br>folderToken: 文件夹 token（可选） |
| `GetDocument(documentID)` | 获取文档信息 | documentID: 文档 ID |
| `GetDocumentRawContent(documentID)` | 获取文档纯文本内容 | documentID: 文档 ID |
| `ListDocumentBlocks(documentID)` | 获取文档所有块 | documentID: 文档 ID |
| `CreateDocumentBlock(documentID, parentID, blockType, children)` | 创建文档块 | documentID: 文档 ID<br>parentID: 父块 ID<br>blockType: 块类型<br>children: 子块列表 |
| `UpdateDocumentBlock(documentID, blockID, block)` | 更新文档块 | documentID: 文档 ID<br>blockID: 块 ID<br>block: 新内容 |
| `DeleteDocumentBlocks(documentID, blockIDs)` | 删除文档块 | documentID: 文档 ID<br>blockIDs: 要删除的块 ID 列表 |

### 辅助函数

| 函数 | 说明 | 示例 |
|------|------|------|
| `CreateTextBlock(text)` | 创建文本块 | `CreateTextBlock("这是一段文本")` |
| `CreateHeadingBlock(text, level)` | 创建标题块 | `CreateHeadingBlock("标题", 1)` |
| `CreateBulletListBlock(text)` | 创建无序列表块 | `CreateBulletListBlock("列表项")` |
| `CreateOrderedListBlock(text)` | 创建有序列表块 | `CreateOrderedListBlock("第一项")` |
| `CreateCodeBlock(code, language)` | 创建代码块 | `CreateCodeBlock("code", "go")` |

## 🔍 块类型说明

| 类型值 | 说明 |
|--------|------|
| 1 | 页面块（Page Block）|
| 2 | 文本块（Text Block）|
| 3 | 标题块（Heading Block）|
| 4 | 无序列表（Bullet List）|
| 5 | 有序列表（Ordered List）|
| 6 | 代码块（Code Block）|
| 7 | 引用块（Quote Block）|
| 8 | 待办事项（Todo Block）|
| 9 | 表格（Table Block）|

## 🛠️ 常见问题

### Q1: 如何获取文档的访问链接？

文档创建后，可以通过以下格式访问：
```
https://your-domain.feishu.cn/docx/{documentID}
```

### Q2: 如何在指定位置插入内容块？

在调用 `CreateDocumentBlock` 时，可以指定 `Index` 参数：
- `-1`: 添加到末尾（默认）
- `0`: 添加到开头
- 其他数字: 添加到指定位置

### Q3: 如何创建更复杂的内容（如表格、图片等）？

查看飞书开放平台文档获取更多块类型：
https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/document-block/create

### Q4: 权限不足怎么办？

确保你的应用具有以下权限：
1. 在飞书开放平台添加权限
2. 发布应用版本
3. 在企业管理后台启用应用

## 📚 相关资源

- [飞书开放平台 - 云文档 API](https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/overview)
- [飞书 Go SDK 文档](https://github.com/larksuite/oapi-sdk-go)
- [项目 README](README.md)
- [权限配置指南](PERMISSION_GUIDE.md)

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License
