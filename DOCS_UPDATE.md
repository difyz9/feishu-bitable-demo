# 云文档功能更新说明

## 📌 更新概述

本次更新为项目添加了完整的飞书云文档操作功能，现在您可以通过 Go SDK 创建、编辑和管理飞书云文档。

## ✨ 新增功能

### 1. 核心功能模块（feishu/docs.go）

| 功能 | 方法 | 说明 |
|------|------|------|
| 创建云文档 | `CreateDocument` | 在指定文件夹创建新的云文档 |
| 获取文档信息 | `GetDocument` | 读取文档的基本信息（标题、ID等） |
| 获取文档内容 | `GetDocumentRawContent` | 获取文档的纯文本内容 |
| 列出文档块 | `ListDocumentBlocks` | 获取文档所有内容块的列表 |
| 创建文档块 | `CreateDocumentBlock` | 向文档中添加新的内容块 |
| 更新文档块 | `UpdateDocumentBlock` | 修改已有文档块的内容 |
| 删除文档块 | `DeleteDocumentBlocks` | 删除指定的文档块 |

### 2. 辅助函数

便捷的内容块创建函数，简化开发流程：

```go
// 创建文本段落
block := feishu.CreateTextBlock("这是一段文本")

// 创建标题（支持 H1-H9）
heading := feishu.CreateHeadingBlock("这是标题", 1)

// 创建无序列表
bullet := feishu.CreateBulletListBlock("列表项内容")

// 创建有序列表
ordered := feishu.CreateOrderedListBlock("第一项")

// 创建代码块（支持语法高亮）
code := feishu.CreateCodeBlock("fmt.Println(\"Hello\")", "go")
```

### 3. 示例程序（main_docs.go）

完整的云文档操作演示程序，包括：

1. ✅ 创建云文档
2. ✅ 获取文档信息
3. ✅ 添加多种类型的内容（标题、文本、列表、代码块）
4. ✅ 读取文档所有块
5. ✅ 获取文档纯文本内容
6. ✅ （可选）删除文档块

运行方式：
```bash
# 方式一：使用脚本
./run_docs.sh

# 方式二：直接运行
go run main_docs.go
```

### 4. 文档资料

#### DOCS_GUIDE.md
详细的云文档操作指南，包含：
- 前置准备和权限配置
- 快速开始指南
- 完整的代码示例
- API 参考文档
- 块类型说明
- 常见问题解答

#### 更新的配置文件
`config.example.yaml` 中新增了云文档相关的配置说明和权限要求。

## 🔧 配置要求

### 1. 权限配置

在飞书开放平台为应用添加以下权限：

**必需权限：**
- ✅ `docx:document` - 查看、编辑、创建和删除云文档
- ✅ `docx:document:readonly` - 查看云文档

**可选权限：**
- `drive:drive` - 管理云空间文件（用于指定文件夹创建文档）

### 2. 配置文件

在 `config.yaml` 中添加：

```yaml
feishu:
  app_id: "cli_xxxxxxxxxxxxx"
  app_secret: "xxxxxxxxxxxxxxxxxxxxx"
  folder_token: ""  # 可选，指定创建文档的文件夹
```

## 📦 项目结构变化

```diff
feishu_golang/
├── feishu/
│   ├── client.go
│   ├── records.go
│   ├── table.go
+   ├── docs.go              # ⭐️ 新增：云文档操作
│   ├── helpers.go
│   └── types.go
├── main.go
├── main_create.go
+├── main_docs.go            # ⭐️ 新增：云文档示例程序
├── config.yaml
├── config.example.yaml      # 更新：添加云文档配置说明
├── README.md                # 更新：添加云文档功能介绍
├── CREATE_TABLE_GUIDE.md
+├── DOCS_GUIDE.md           # ⭐️ 新增：云文档操作指南
├── PERMISSION_GUIDE.md
├── PROJECT_SUMMARY.md       # 更新：添加云文档模块说明
├── run.sh
+├── run_docs.sh             # ⭐️ 新增：云文档运行脚本
└── create_table.sh
```

## 🚀 快速体验

### 步骤 1：配置权限
在飞书开放平台添加 `docx:document` 权限并发布应用版本。

### 步骤 2：填写配置
```bash
cp config.example.yaml config.yaml
# 编辑 config.yaml，填入 app_id 和 app_secret
```

### 步骤 3：运行示例
```bash
./run_docs.sh
```

程序将自动创建一个包含丰富内容的云文档，包括标题、文本、列表、代码块等。

## 💡 使用场景

云文档功能适用于以下场景：

1. **自动生成报告** - 从数据库或多维表格提取数据，生成格式化的文档报告
2. **文档模板化** - 批量创建具有统一格式的文档
3. **知识库管理** - 自动创建和更新技术文档、API 文档等
4. **会议纪要** - 自动整理和格式化会议内容
5. **数据展示** - 将数据以文档形式展示，支持代码、表格等多种格式

## 📊 代码示例

### 创建一个完整的项目文档

```go
package main

import (
    "feishu_bitable_demo/feishu"
)

func main() {
    // 初始化客户端
    client := feishu.NewMultiTableClient("app_id", "app_secret")
    
    // 创建文档
    resp, _ := client.CreateDocument("项目技术文档", "")
    docID := *resp.Data.Document.DocumentId
    
    // 获取页面块 ID
    blocks, _ := client.ListDocumentBlocks(docID)
    pageBlockID := *blocks.Data.Items[0].BlockId
    
    // 添加内容
    content := []*feishu.Block{
        feishu.CreateHeadingBlock("项目概述", 1),
        feishu.CreateTextBlock("这是一个飞书 API 集成项目。"),
        
        feishu.CreateHeadingBlock("技术栈", 2),
        feishu.CreateBulletListBlock("Go 1.21+"),
        feishu.CreateBulletListBlock("飞书开放平台 SDK"),
        feishu.CreateBulletListBlock("YAML 配置管理"),
        
        feishu.CreateHeadingBlock("快速开始", 2),
        feishu.CreateOrderedListBlock("安装依赖"),
        feishu.CreateOrderedListBlock("配置文件"),
        feishu.CreateOrderedListBlock("运行程序"),
        
        feishu.CreateHeadingBlock("代码示例", 2),
        feishu.CreateCodeBlock(`package main

import "fmt"

func main() {
    fmt.Println("Hello, Feishu!")
}`, "go"),
    }
    
    client.CreateDocumentBlock(docID, pageBlockID, 2, content)
}
```

## 🔗 相关资源

- [飞书云文档 API 文档](https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/overview)
- [DOCS_GUIDE.md](DOCS_GUIDE.md) - 完整使用指南
- [README.md](README.md) - 项目说明文档
- [飞书开放平台](https://open.feishu.cn/)

## 📝 注意事项

1. **权限要求**：确保应用已添加云文档相关权限并发布版本
2. **文件夹 Token**：如果不指定 folder_token，文档将创建在根目录
3. **块类型**：不同的块类型有不同的参数要求，详见 DOCS_GUIDE.md
4. **错误处理**：建议在生产环境中添加完善的错误处理逻辑

## 🎉 总结

本次更新为项目增加了完整的云文档操作能力，与已有的多维表格功能形成互补，为飞书自动化办公提供了更多可能性。

现在您可以：
- ✅ 管理多维表格数据
- ✅ 创建和编辑云文档
- ✅ 构建完整的飞书自动化工作流

欢迎使用和反馈！
