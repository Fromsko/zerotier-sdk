# ZeroTier MCP Windows 安装脚本

param(
    [string]$Version = "latest",
    [string]$InstallDir = "$env:LOCALAPPDATA\zerotier-mcp"
)

$ErrorActionPreference = "Stop"

Write-Host "📥 下载 ZeroTier MCP (Windows x86_64)..." -ForegroundColor Green

$BinaryName = "zerotier-mcp-windows-amd64.exe"
$DownloadUrl = "https://github.com/fromsko/zerotier-sdk/releases/download/$Version/$BinaryName"

# 创建安装目录
New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null

# 下载二进制
$OutputPath = Join-Path $InstallDir $BinaryName
Write-Host "   URL: $DownloadUrl"
Write-Host "   保存到: $OutputPath"

try {
    Invoke-WebRequest -Uri $DownloadUrl -OutFile $OutputPath
} catch {
    Write-Host "❌ 下载失败: $_" -ForegroundColor Red
    exit 1
}

# 创建快捷方式
$ShortcutPath = Join-Path $InstallDir "zerotier-mcp.exe"
if (Test-Path $ShortcutPath) {
    Remove-Item $ShortcutPath -Force
}
Copy-Item $OutputPath $ShortcutPath

Write-Host "✅ 安装完成！" -ForegroundColor Green
Write-Host ""
Write-Host "📍 安装位置: $InstallDir" -ForegroundColor Cyan
Write-Host ""
Write-Host "🔧 配置 Claude Desktop:" -ForegroundColor Yellow
Write-Host ""
Write-Host "编辑 %APPDATA%\Claude\claude_desktop_config.json:" -ForegroundColor Gray
Write-Host ""
Write-Host '  "mcpServers": {' -ForegroundColor Gray
Write-Host '    "zerotier": {' -ForegroundColor Gray
Write-Host "      `"command`": `"$ShortcutPath`"," -ForegroundColor Gray
Write-Host '      "env": {' -ForegroundColor Gray
Write-Host '        "ZT_CENTRAL_TOKEN": "your_api_token"' -ForegroundColor Gray
Write-Host '      }' -ForegroundColor Gray
Write-Host '    }' -ForegroundColor Gray
Write-Host '  }' -ForegroundColor Gray
