#!/bin/bash

echo "🚀 启动飞书云文档操作示例程序..."
echo ""

# 检查配置文件
if [ ! -f "config.yaml" ]; then
    echo "❌ 错误: config.yaml 文件不存在"
    echo "💡 请先复制 config.example.yaml 为 config.yaml 并填写配置"
    echo ""
    echo "  cp config.example.yaml config.yaml"
    echo ""
    exit 1
fi

# 检查是否填写了配置
if grep -q "你的app_id" config.yaml || grep -q "your_app_id_here" config.yaml; then
    echo "❌ 错误: 请先在 config.yaml 中填写正确的飞书应用配置"
    echo "💡 需要配置以下内容："
    echo "  - app_id: 从飞书开放平台获取"
    echo "  - app_secret: 从飞书开放平台获取"
    echo "  - folder_token: （可选）云空间文件夹 token"
    echo ""
    exit 1
fi

# 运行程序
go run main_docs.go
