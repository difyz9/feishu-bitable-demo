#!/bin/bash

# 飞书多维表格 Go SDK 快速入门脚本

set -e

echo "=================================================="
echo "🚀 飞书多维表格 Go SDK 快速入门"
echo "=================================================="
echo ""

# 检查 Go 环境
if ! command -v go &> /dev/null; then
    echo "❌ 未检测到 Go 环境，请先安装 Go"
    exit 1
fi

echo "✅ Go 版本: $(go version)"
echo ""

# 安装依赖
echo "📦 正在安装依赖..."
go mod tidy
echo "✅ 依赖安装完成"
echo ""

# 编译项目
echo "🔨 正在编译项目..."
go build -o feishu_test main.go
echo "✅ 编译完成"
echo ""

# 检查配置文件
if grep -q "your_app_id_here" config.yaml; then
    echo "⚠️  请先配置 config.yaml 文件"
    echo ""
    echo "需要填写以下信息："
    echo "  1. app_id: 飞书应用 ID"
    echo "  2. app_secret: 飞书应用密钥"
    echo "  3. app_token: 多维表格 app_token"
    echo "  4. table_id: 表格 table_id"
    echo ""
    echo "📖 详细说明请参考 README.md"
    exit 1
fi

# 运行测试
echo "🧪 运行测试程序..."
echo ""
./feishu_test

echo ""
echo "=================================================="
echo "🎉 测试完成！"
echo "=================================================="
