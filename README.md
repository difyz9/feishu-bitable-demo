# 飞书多维表格 Go SDK 示例

这是一个完整的 Go 语言实现，用于操作飞书多维表格（Lark Base），支持创建、读取、更新和删除记录等操作。

## 功能特性

✅ **完整的 CRUD 操作**
- 创建单个记录
- 批量创建记录
- 读取记录
- 更新记录
- 批量更新记录
- 删除记录
- 查询记录列表

✅ **创建多维表格和数据表** ⭐️ 新增
- 创建多维表格（Bitable App）
- 创建数据表（Table）
- 定义自定义字段
- 一键创建并写入数据

✅ **自动 Token 管理**
- 自动获取和刷新 Access Token
- Token 缓存机制

✅ **丰富的字段类型支持**
- 文本、数字、日期时间
- 复选框、单选/多选
- 链接、人员、电话
- 地理位置

## 项目结构

```
.
├── feishu/              # 飞书多维表格 SDK
│   ├── client.go        # 客户端和认证
│   ├── types.go         # 数据类型定义
│   ├── records.go       # 记录操作
│   ├── table.go         # 表格和数据表创建 ⭐️
│   └── helpers.go       # 辅助函数
├── main.go              # 测试和验证程序（操作已有表格）
├── main_create.go       # 创建表格并写入数据 ⭐️ 新增
├── config.yaml          # 配置文件
├── go.mod               # Go 模块配置
├── README.md            # 说明文档
├── CREATE_TABLE_GUIDE.md # 创建表格使用指南 ⭐️
└── PERMISSION_GUIDE.md  # 权限配置指南
```方式一：创建新的多维表格（推荐用于测试）⭐️

这种方式会自动创建一个新的多维表格和数据表，无需手动创建。

#### 1. 配置

编辑 `config.yaml`：

```yaml
feishu:
  app_id: "你的app_id"
  app_secret: "你的app_secret"
  folder_token: ""  # 可选，留空则创建在根目录
```

#### 2. 运行

```bash
# 使用脚本
./create_table.sh

# 或直接运行
go run main_create.go
```

程序会自动：
1. 创建多维表格
2. 创建包含8个字段的数据表
3. 写入示例数据
4. 执行各种操作测试

详细说明请查看 [CREATE_TABLE_GUIDE.md](CREATE_TABLE_GUIDE.md)

---

### 方式二：操作已有的多维表格

如果你已经有一个多维表格，使用这种方式。

#### 

## 快速开始

### 1. 准备工作

#### 1.1 创建飞书应用

1. 登录 [飞书开放平台](https://open.feishu.cn/)
2. 创建企业自建应用
3. 获取 `app_id` 和 `app_secret`

#### 1.2 配置权限

在应用管理后台，给应用添加以下权限：

**必须添加的权限：**
- ✅ **查看、编辑、管理多维表格** (`bitable:app:readonly` 和 `bitable:app`)
  - 在"权限管理" -> "多维表格"中开启
  - 包括以下权限：
    - 查看多维表格
    - 编辑多维表格
    - 管理多维表格

**重要提示：**
1. 添加权限后，需要在"版本管理与发布"中**创建新版本并发布**
2. 如果是企业自建应用，需要管理员审核通过
3. 发布后，应用才能获得相应的权限

#### 1.3 创建多维表格

1. 在飞书中创建一个多维表格
2. 创建以下字段（用于测试）：

| 字段名 | 类型 | 说明 |
|--------|------|------|
| 名称 | 文本 | 产品名称 |
| 数量 | 数字 | 产品数量 |
| 价格 | 数字 | 产品价格 |
| 描述 | 文本 | 产品描述 |
| 创建时间 | 日期 | 创建时间 |
| 是否上架 | 复选框 | 是否上架 |

#### 1.4 获取表格信息

- **app_token**: 打开飞书多维表格，从浏览器地址栏获取
  - 格式：`https://xxx.feishu.cn/base/bascnxxxxxx`
  - `bascnxxxxxx` 就是 `app_token`
  
- **table_id**: 点击表格后，从地址栏获取
  - 格式：`https://xxx.feishu.cn/base/bascnxxxxxx?table=tblxxxxxx`
  - `tblxxxxxx` 就是 `table_id`

### 2. 安装依赖

```bash
go mod tidy
```

### 3. 配置

编辑 `config.yaml` 文件，填入你的配置：

```yaml
feishu:
  app_id: "cli_xxxxxxxxxx"          # 你的应用 ID
  app_secret: "xxxxxxxxxxxxx"        # 你的应用密钥
  app_token: "bascnxxxxxxxxx"        # 多维表格 app_token
  table_id: "tblxxxxxxxxx"           # 表格 table_id
```

### 4. 运行测试

```bash
go run main.go
```

## 使用示例

### 示例 1：创建新的多维表格并写入数据 ⭐️

```go
package main

import (
    "feishu_bitable_demo/feishu"
    larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func main() {
    client := feishu.NewMultiTableClient(appID, appSecret)
    
    // 定义表格字段
    fields := []*larkbitable.AppTableCreateHeader{
        {
            FieldName: ptrString("产品名称"),
            Type:      ptrInt(1), // 1 = 文本
        },
        {
            FieldName: ptrString("价格"),
            Type:      ptrInt(2), // 2 = 数字
        },
    }
    
    // 创建多维表格和数据表
    appToken, tableID, err := client.CreateAppAndTable(
        "产品管理系统",
        "",  // folder_token，留空创建在根目录
        "产品列表",
        fields,
    )
    
    if err != nil {
        panic(err)
    }
    
    // 写入数据
    recordFields := map[string]interface{}{
        "产品名称": feishu.CreateTextField("iPhone 15 Pro"),
        "价格": feishu.CreateNumberField(7999.00),
    }
    
    recordID, _ := client.CreateRecord(appToken, tableID, recordFields)
}

func ptrString(s string) *string { return &s }
func ptrInt(i int) *int { return &i }
```

### 示例 2：操作已有表格 - 初始化客户端

```go
import "feishu_bitable_demo/feishu"

client := feishu.NewMultiTableClient(appID, appSecret)
```

### 创建记录

```go
fields := map[string]interface{}{
    "名称":   feishu.CreateTextField("测试产品"),
    "数量":   feishu.CreateNumberField(100),
    "价格":   feishu.CreateNumberField(299.99),
    "描述":   feishu.CreateTextField("这是一个测试产品"),
    "创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
    "是否上架": feishu.CreateCheckboxField(true),
}

recordID, err := client.CreateRecord(appToken, tableID, fields)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("创建成功，记录ID: %s\n", recordID)
```

### 读取记录

```go
fields, err := client.GetRecord(appToken, tableID, recordID)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("记录内容: %+v\n", fields)
```

### 更新记录

```go
fields := map[string]interface{}{
    "数量": feishu.CreateNumberField(200),
    "价格": feishu.CreateNumberField(399.99),
}

err := client.UpdateRecord(appToken, tableID, recordID, fields)
if err != nil {
    log.Fatal(err)
}
```

### 查询记录列表

```go
items, pageToken, hasMore, err := client.ListRecords(appToken, tableID, 10, "")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("查询到 %d 条记录\n", len(items))
for _, item := range items {
    fmt.Printf("%+v\n", item)
}
```

### 批量创建记录

```go
records := []feishu.CreateRecordRequest{
    {
        Fields: map[string]interface{}{
            "名称": feishu.CreateTextField("产品 A"),
            "数量": feishu.CreateNumberField(50),
            "价格": feishu.CreateNumberField(199.99),
        },
    },
    {
        Fields: map[string]interface{}{
            "名称": feishu.CreateTextField("产品 B"),
            "数量": feishu.CreateNumberField(75),
            "价格": feishu.CreateNumberField(249.99),
        },
    },
}

recordIDs, err := client.BatchCreateRecords(appToken, tableID, records)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("批量创建了 %d 条记录\n", len(recordIDs))
```

### 删除记录

```go
err := client.DeleteRecord(appToken, tableID, recordID)
if err != nil {
    log.Fatal(err)
}
```

## 字段类型辅助函数

SDK 提供了多种字段类型的辅助函数：

```go
// 文本
feishu.CreateTextField("文本内容")

// 数字
feishu.CreateNumberField(123.45)

// 日期时间（Unix 时间戳）
feishu.CreateDateTimeField(1609459200)

// 日期时间（time.Time）
feishu.CreateDateTimeFieldFromTime(time.Now())

// 复选框
feishu.CreateCheckboxField(true)

// 单选
feishu.CreateSingleSelectField("选项名称")

// 多选
feishu.CreateMultiSelectField([]string{"选项1", "选项2"})

// 链接
feishu.CreateURLField("https://example.com", "链接文本")

// 人员
feishu.CreateUserField([]string{"ou_xxxxx", "ou_yyyyy"})

// 电话
feishu.CreatePhoneField("13800138000")

// 地理位置
feishu.CreateLocationField("北京市朝阳区")
```

## 测试验证

运行 `main.go` 将执行以下测试：

1. ✅ 获取 Access Token
2. ✅ 创建单个记录
3. ✅ 读取记录
4. ✅ 更新记录
5. ✅ 查询记录列表
6. ✅ 批量创建记录
7. ✅ 批量更新记录
8. ✅ 删除记录

预期输出：

```
=================================================
🚀 飞书多维表格操作验证程序
=================================================

📝 步骤 1: 测试获取 Access Token
✅ 成功获取 Access Token: t-xxxxxxxxxxxxxx...

📝 步骤 2: 测试创建单个记录
✅ 成功创建记录，ID: recxxxxxx

📝 步骤 3: 测试读取记录
   记录字段: map[...]
✅ 成功读取记录

📝 步骤 4: 测试更新记录
✅ 成功更新记录

📝 步骤 5: 测试查询所有记录
   查询到 4 条记录
   是否有更多: false
✅ 成功查询记录

📝 步骤 6: 测试批量创建记录
✅ 成功批量创建 3 条记录

📝 步骤 7: 测试批量更新记录
✅ 成功批量更新记录

📝 步骤 8: 测试删除记录
✅ 成功删除记录 ID: recxxxxxx
✅ 成功清理批量创建的记录

=================================================
🎉 所有测试通过！飞书多维表格操作功能正常
=================================================
```

## API 文档

### Client 方法

#### `GetAccessToken() (string, error)`
获取访问令牌，自动处理缓存和刷新。

#### `CreateRecord(appToken, tableID string, fields map[string]interface{}) (string, error)`
创建单个记录，返回记录 ID。

#### `BatchCreateRecords(appToken, tableID string, records []CreateRecordRequest) ([]string, error)`
批量创建记录，返回记录 ID 列表。

#### `GetRecord(appToken, tableID, recordID string) (map[string]interface{}, error)`
获取单个记录的字段内容。

#### `UpdateRecord(appToken, tableID, recordID string, fields map[string]interface{}) error`
更新记录的字段。

#### `BatchUpdateRecords(appToken, tableID string, records []struct{...}) error`
批量更新多条记录。

#### `DeleteRecord(appToken, tableID, recordID string) error`
删除指定记录。

#### `ListRecords(appToken, tableID string, pageSize int, pageToken string) ([]map[string]interface{}, string, bool, error)`
查询记录列表，支持分页。返回：记录列表、下一页token、是否有更多、错误。

## 注意事项

1. **权限配置**：确保应用有足够的权限访问多维表格
2. **字段名称**：字段名称必须与飞书多维表格中的字段名完全一致
3. **字段类型**：确保字段类型与表格中定义的类型匹配
4. **速率限制**：注意飞书 API 的调用频率限制
5. **错误处理**：生产环境中应该添加更完善的错误处理和重试机制

## 常见问题

### Q1: 报错 "code=99991663" 或 "code=91403 Forbidden"
**A**: 权限不足，请检查：
1. 在飞书开放平台为应用添加多维表格权限（见上面权限配置步骤）
2. **添加权限后必须创建新版本并发布**
3. 确保应用已经通过审核
4. 在飞书多维表格设置中，将该应用添加为协作者
5. 可以在多维表格 -> 设置 -> 高级设置 -> 添加应用为协作者

### Q2: 报错 "code=99991668"  
**A**: app_token 或 table_id 错误，请检查配置。

### Q3: 字段写入失败
**A**: 检查字段名称是否与表格中的字段名完全一致（区分大小写）。

### Q4: Token 过期
**A**: 官方 SDK 会自动管理和刷新 Token，如果仍有问题，检查 app_id 和 app_secret 是否正确。

## 参考资料

- [飞书开放平台文档](https://open.feishu.cn/document/)
- [多维表格 API 文档](https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-record/list)

## License

MIT License

## 作者

Created with ❤️ for Feishu developers
