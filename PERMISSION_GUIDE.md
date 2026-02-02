# 飞书多维表格权限配置指南

## 问题：code=91403 Forbidden

这个错误表示应用没有足够的权限访问多维表格。

## 解决步骤

### 方法一：在飞书开放平台配置权限（推荐）

1. **登录飞书开放平台**
   - 访问 https://open.feishu.cn/
   - 找到你的应用

2. **添加权限**
   - 点击左侧菜单"权限管理"
   - 找到"多维表格" (Bitable) 相关权限
   - 开启以下权限：
     - `查看多维表格` (bitable:app:readonly)
     - `编辑多维表格` (bitable:app)
     - `管理多维表格`

3. **发布应用**（重要！）
   - 点击左侧菜单"版本管理与发布"
   - 点击"创建版本"
   - 填写版本说明
   - 提交审核（如果是企业应用需要管理员审核）
   - **等待审核通过并发布**
   - ⚠️ 只有发布后，权限才会生效！

4. **验证权限**
   - 发布成功后，重新运行程序测试

### 方法二：在多维表格中添加应用为协作者

如果方法一不行，尝试直接在多维表格中授权：

1. **打开你的飞书多维表格**

2. **进入设置**
   - 点击右上角的"..."按钮
   - 选择"设置"

3. **添加应用**
   - 找到"高级设置"或"应用访问"
   - 点击"添加应用"或"添加机器人"
   - 搜索并添加你的应用
   - 给予"编辑"或"管理"权限

4. **保存设置**

### 验证配置

运行测试程序：

```bash
go run main.go
```

成功的输出应该类似：

```
=================================================
🚀 飞书多维表格操作验证程序
=================================================

📝 步骤 1: 初始化飞书客户端
✅ 成功初始化客户端

📝 步骤 2: 测试创建单个记录
✅ 成功创建记录，ID: recxxxxxx
...
```

## 其他常见错误

### code=99991663
- 应用未授权
- 解决方法同上

### code=99991668
- app_token 或 table_id 错误
- 检查 config.yaml 中的配置是否正确

### code=99991661
- access_token 无效
- 检查 app_id 和 app_secret 是否正确

## 获取 app_token 和 table_id

### 获取 app_token

1. 打开飞书多维表格
2. 查看浏览器地址栏，格式为：
   ```
   https://xxx.feishu.cn/base/bascnxxxxxx?table=tblxxxxxx
   ```
3. `bascnxxxxxx` 就是 app_token

### 获取 table_id

1. 在多维表格中点击具体的某个表格
2. 查看地址栏，格式为：
   ```
   https://xxx.feishu.cn/base/bascnxxxxxx?table=tblxxxxxx&view=vewxxxx
   ```
3. `tblxxxxxx` 就是 table_id

## 完整的配置示例

```yaml
feishu:
  app_id: "cli_a9b2931b14b89bb4"      # 从飞书开放平台获取
  app_secret: "s27dfD78pod2CG2Uq..."   # 从飞书开放平台获取
  app_token: "VADqbPj3EaA9EIszDNNc..."  # 从多维表格 URL 获取
  table_id: "tblCwGFKVXY6TBJW"         # 从多维表格 URL 获取
```

## 需要帮助？

如果仍然遇到问题，请检查：

1. ✅ 应用是否已创建并获取了 app_id 和 app_secret
2. ✅ 权限是否已添加
3. ✅ 应用版本是否已发布
4. ✅ app_token 和 table_id 是否正确
5. ✅ 表格中的字段名称是否与代码中的一致

参考官方文档：
- https://open.feishu.cn/document/server-docs/docs/bitable-v1/notification
- https://open.feishu.cn/document/home/introduction-to-scope-and-authorization/list-of-permissions-by-scope
