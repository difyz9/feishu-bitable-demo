# 创建多维表格并写入数据 - 使用指南

## 📖 概述

本项目提供了两种使用方式：

1. **main.go** - 操作已有的多维表格
2. **main_create.go** - 创建新的多维表格并写入数据 ✨（推荐用于测试）

## 🚀 快速开始：创建新表格并写入数据

### 步骤 1：配置应用凭证

编辑 `config.yaml`：

```yaml
feishu:
  app_id: "你的app_id"
  app_secret: "你的app_secret"
  folder_token: ""  # 可选，留空则创建在根目录
```

### 步骤 2：运行创建程序

```bash
go run main_create.go
```

### 步骤 3：查看结果

程序会自动：
1. ✅ 创建一个新的多维表格
2. ✅ 创建包含8个字段的数据表
3. ✅ 写入示例数据
4. ✅ 执行各种操作（读取、更新、查询、删除）
5. ✅ 输出表格访问链接

## 📋 程序功能

### main_create.go 完整流程

```
📝 步骤 1: 创建多维表格和数据表
   ├─ 创建多维表格（App）
   ├─ 创建数据表（Table）
   └─ 定义8个字段：产品名称、库存数量、单价、状态等

📝 步骤 2: 写入单条记录
   └─ 写入 iPhone 15 Pro 产品信息

📝 步骤 3: 批量写入记录
   └─ 批量写入 5 个产品（MacBook、iPad、AirPods等）

📝 步骤 4: 读取记录
   └─ 读取并显示记录详情

📝 步骤 5: 更新记录
   └─ 更新库存和价格信息

📝 步骤 6: 查询所有记录
   └─ 列出所有记录

📝 步骤 7: 删除测试记录
   └─ 清理测试数据
```

## 📊 创建的表格结构

程序会自动创建包含以下字段的表格：

| 字段名 | 类型 | 说明 |
|--------|------|------|
| 产品名称 | 文本 | 产品的名称 |
| 库存数量 | 数字 | 库存数量 |
| 单价 | 数字 | 产品价格 |
| 状态 | 单选 | 在售/预售/下架 |
| 标签 | 多选 | 热销/新品/推荐等 |
| 创建时间 | 日期 | 创建时间戳 |
| 是否上架 | 复选框 | 是否上架 |
| 产品描述 | 文本 | 产品详细描述 |

## 🎯 预期输出

```
=================================================
🚀 飞书多维表格完整操作示例
=================================================

📝 步骤 1: 创建多维表格和数据表
✅ 成功创建多维表格
   App Token: bascnXXXXXXXXXXXX
   Table ID: tblXXXXXXXX

📝 步骤 2: 写入单条记录
✅ 成功写入记录，ID: recXXXXXX

📝 步骤 3: 批量写入记录
✅ 成功批量写入 5 条记录

📝 步骤 4: 读取记录
   记录内容: map[产品名称:iPhone 15 Pro ...]
✅ 成功读取记录

📝 步骤 5: 更新记录
✅ 成功更新记录

📝 步骤 6: 查询所有记录
   共查询到 6 条记录
✅ 成功查询记录

📝 步骤 7: 删除测试记录
✅ 成功删除记录 ID: recXXXXXX
✅ 成功清理批量创建的记录

=================================================
🎉 所有测试完成！
=================================================

📊 多维表格访问地址：
https://your-domain.feishu.cn/base/bascnXXXX?table=tblXXXX
```

## 🔧 API 使用示例

### 1. 创建多维表格和数据表

```go
// 定义字段
fields := []*larkbitable.AppTableCreateHeader{
    {
        FieldName: ptrString("产品名称"),
        Type:      ptrInt(1), // 1 = 文本
    },
    {
        FieldName: ptrString("库存数量"),
        Type:      ptrInt(2), // 2 = 数字
    },
}

// 创建
appToken, tableID, err := client.CreateAppAndTable(
    "产品管理系统",
    folderToken,
    "产品列表",
    fields,
)
```

### 2. 写入记录

```go
fields := map[string]interface{}{
    "产品名称": feishu.CreateTextField("iPhone 15 Pro"),
    "库存数量": feishu.CreateNumberField(100),
    "单价":   feishu.CreateNumberField(7999.00),
    "创建时间": feishu.CreateDateTimeFieldFromTime(time.Now()),
    "是否上架": feishu.CreateCheckboxField(true),
}

recordID, err := client.CreateRecord(appToken, tableID, fields)
```

## 📚 字段类型参考

| 类型代码 | 字段类型 | 说明 |
|---------|---------|------|
| 1 | 文本 | Text |
| 2 | 数字 | Number |
| 3 | 单选 | Single Select |
| 4 | 多选 | Multiple Select |
| 5 | 日期 | Date |
| 7 | 复选框 | Checkbox |
| 11 | 人员 | Person |
| 13 | 电话 | Phone |
| 15 | 超链接 | URL |
| 17 | 附件 | Attachment |
| 18 | 单向关联 | Link |
| 20 | 公式 | Formula |
| 21 | 双向关联 | Two-way Link |
| 22 | 地理位置 | Location |
| 1001 | 创建时间 | Created Time |
| 1002 | 修改时间 | Modified Time |
| 1003 | 创建人 | Created By |
| 1004 | 修改人 | Modified By |
| 1005 | 自动编号 | Auto Number |

## ⚠️ 权限要求

确保应用具有以下权限：
- ✅ 创建多维表格
- ✅ 编辑多维表格
- ✅ 查看多维表格

如果遇到权限错误，请参考 `PERMISSION_GUIDE.md`

## 🔍 故障排除

### 错误：code=91403 Forbidden
- 应用没有创建多维表格的权限
- 解决：在飞书开放平台添加权限并发布应用

### 错误：code=1254044
- folder_token 无效
- 解决：留空或提供有效的文件夹 token

### 表格创建成功但找不到
- 检查飞书云空间根目录
- 使用程序输出的访问链接直接打开

## 📖 两种模式对比

| 特性 | main.go | main_create.go |
|------|---------|----------------|
| 用途 | 操作已有表格 | 创建新表格 |
| 需要配置 | app_token, table_id | 仅需 app_id, app_secret |
| 适用场景 | 生产环境 | 测试、演示 |
| 创建表格 | ❌ | ✅ |
| 数据操作 | ✅ | ✅ |

## 🎓 学习路径

1. **入门**：运行 `main_create.go` 了解完整流程
2. **理解**：查看代码理解每个步骤
3. **实践**：修改字段和数据适配你的需求
4. **应用**：使用 `main.go` 操作生产环境的表格

## 🔗 相关文档

- [飞书开放平台](https://open.feishu.cn/)
- [多维表格 API](https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-record/list)
- [字段类型文档](https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-field/guide)

---

**开始创建你的第一个多维表格吧！** 🚀
