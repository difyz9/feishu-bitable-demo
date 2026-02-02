#!/bin/bash

# 飞书多维表格创建和操作脚本

set -e

echo "=================================================="
echo "🚀 飞书多维表格 - 创建表格并写入数据"
echo "=================================================="
echo ""

# 检查 Go 环境
if ! command -v go &> /dev/null; then
    echo "❌ 未检测到 Go 环境，请先安装 Go"
    exit 1
fi

echo "✅ Go 版本: $(go version)"
echo ""

# 检查配置文件
if [ ! -f "config.yaml" ]; then
    echo "❌ 配置文件 config.yaml 不存在"
    exit 1
fi

# 检查配置
if grep -q "your_app_id_here" config.yaml; then
    echo "❌ 请先配置 config.yaml 文件"
    echo ""
    echo "需要填写："
    echo "  - app_id: 飞书应用 ID"
    echo "  - app_secret: 飞书应用密钥"
    echo ""
    exit 1
fi

# 安装依赖
echo "📦 正在检查依赖..."
go mod tidy
echo "✅ 依赖检查完成"
echo ""

# 运行创建表格程序
echo "🔨 正在创建多维表格并写入数据..."
echo ""
go run main_create.go

echo ""
echo "=================================================="
echo "✅ 完成！"
echo "=================================================="
echo ""
echo "提示："
echo "  - 可以在飞书中查看新创建的多维表格"
echo "  - 程序已输出表格访问链接"
echo ""
