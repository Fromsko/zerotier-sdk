#!/bin/bash
# ZeroTier MCP 安装脚本

set -e

VERSION="${1:-latest}"
INSTALL_DIR="${2:-$HOME/.local/bin}"

# 检测操作系统和架构
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
  x86_64)
    ARCH="amd64"
    ;;
  aarch64)
    ARCH="arm64"
    ;;
  *)
    echo "❌ 不支持的架构: $ARCH"
    exit 1
    ;;
esac

case "$OS" in
  linux)
    OS="linux"
    ;;
  darwin)
    OS="darwin"
    ;;
  *)
    echo "❌ 不支持的操作系统: $OS"
    exit 1
    ;;
esac

BINARY_NAME="zerotier-mcp-${OS}-${ARCH}"
DOWNLOAD_URL="https://github.com/fromsko/zerotier-sdk/releases/download/${VERSION}/${BINARY_NAME}"

echo "📥 下载 ZeroTier MCP ($OS/$ARCH)..."
echo "   URL: $DOWNLOAD_URL"

# 创建安装目录
mkdir -p "$INSTALL_DIR"

# 下载二进制
if command -v curl &> /dev/null; then
  curl -L -o "$INSTALL_DIR/$BINARY_NAME" "$DOWNLOAD_URL"
elif command -v wget &> /dev/null; then
  wget -O "$INSTALL_DIR/$BINARY_NAME" "$DOWNLOAD_URL"
else
  echo "❌ 需要 curl 或 wget"
  exit 1
fi

# 设置执行权限
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# 创建符号链接
ln -sf "$INSTALL_DIR/$BINARY_NAME" "$INSTALL_DIR/zerotier-mcp"

echo "✅ 安装完成！"
echo ""
echo "📍 安装位置: $INSTALL_DIR/zerotier-mcp"
echo ""
echo "🔧 配置 Claude Desktop:"
echo ""
echo "编辑 ~/.config/Claude/claude_desktop_config.json:"
echo ""
echo '  "mcpServers": {'
echo '    "zerotier": {'
echo "      \"command\": \"$INSTALL_DIR/zerotier-mcp\","
echo '      "env": {'
echo '        "ZT_CENTRAL_TOKEN": "your_api_token"'
echo '      }'
echo '    }'
echo '  }'
