# MCP - Model Context Protocol 集成

为 ZeroTier SDK 提供 MCP 服务支持，可与 Claude、Cursor 等 AI 工具集成。

## 安装

### 方式一：使用安装脚本（推荐）

**Linux/macOS:**

```bash
curl -fsSL https://raw.githubusercontent.com/fromsko/zerotier-sdk/main/scripts/install-mcp.sh | bash
# 或指定版本
curl -fsSL https://raw.githubusercontent.com/fromsko/zerotier-sdk/main/scripts/install-mcp.sh | bash -s v1.1.0
```

**Windows (PowerShell):**

```powershell
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/fromsko/zerotier-sdk/main/scripts/install-mcp.ps1" -OutFile install-mcp.ps1
.\install-mcp.ps1
```

### 方式二：手动下载

从 [Releases](https://github.com/fromsko/zerotier-sdk/releases) 页面下载对应平台的二进制文件。

### 方式三：从源码构建

```bash
git clone https://github.com/fromsko/zerotier-sdk.git
cd zerotier-sdk
make build-mcp
# 二进制文件在 dist/ 目录
```

## 快速开始

```go
package main

import (
    "log"
    "github.com/fromsko/zerotier-sdk/mcp"
)

func main() {
    // 创建 MCP 服务（仅本地 API）
    s := mcp.New("zerotier", "1.1.0")

    // 启动服务
    if err := s.ServeStdio(); err != nil {
        log.Fatal(err)
    }
}
```

## 配置选项

```go
// 同时启用本地和云端 API
s := mcp.New("zerotier", "1.1.0",
    mcp.WithCentralToken("your_api_token"),
    mcp.WithCentralV2Token("your_service_account_token"),
)

// 自定义客户端
s := mcp.New("zerotier", "1.1.0",
    mcp.WithLocalClient(myLocalClient),
    mcp.WithCentralClient(myCentralClient),
)
```

## 可用工具

### 本地 API

| 工具                 | 描述                   |
| -------------------- | ---------------------- |
| `zt_status`          | 获取节点状态           |
| `zt_networks`        | 列出已加入的网络       |
| `zt_network_info`    | 获取网络详情           |
| `zt_update_network`  | 更新网络设置           |
| `zt_join`            | 加入网络               |
| `zt_leave`           | 离开网络               |
| `zt_peers`           | 列出 Peers             |
| `zt_peer_info`       | 获取 Peer 详情         |
| `zt_controller_*`    | 本地控制器网络/成员管理 |

### 云端 API（Central v1，需配置 ZT_CENTRAL_TOKEN）

| 工具                                          | 描述                      |
| --------------------------------------------- | ------------------------- |
| `zt_central_networks`                         | 列出云端网络              |
| `zt_central_network_info`                     | 网络详情                  |
| `zt_central_create_network`                   | 创建网络                  |
| `zt_central_update_network`                   | 更新网络                  |
| `zt_central_delete_network`                   | 删除网络                  |
| `zt_central_members`                          | 列出网络成员              |
| `zt_central_member_info`                      | 成员详情                  |
| `zt_central_authorize`                        | 授权成员                  |
| `zt_central_deauthorize`                      | 取消授权                  |
| `zt_central_authorize_with_ip`                | 授权并设置 IP/名称        |
| `zt_central_set_name`                         | 设置成员名称              |
| `zt_central_set_ip`                           | 设置成员 IP               |
| `zt_central_delete_member`                    | 删除成员                  |
| `zt_central_organization`                     | 组织信息（当前组织）      |
| `zt_central_organization_members`             | 组织成员列表（需 org_id） |
| `zt_central_invitations`                      | 邀请列表                  |
| `zt_central_create_invitation`                | 创建邀请                  |
| `zt_central_accept_invitation`                | 接受邀请                  |
| `zt_central_decline_invitation`               | 拒绝邀请                  |
| `zt_central_user`                             | 用户信息（需 user_id）    |
| `zt_central_random_token`                     | 获取随机 Token            |
| `zt_central_create_token`                     | 创建用户 Token            |
| `zt_central_delete_token`                     | 删除用户 Token            |
| `zt_central_set_network_user_permissions`     | 设置网络用户权限          |

### Central V2 工具（需配置 ZT_CENTRAL_V2_TOKEN）

| 工具                                      | 描述           |
| ----------------------------------------- | -------------- |
| `zt_central_v2_orgs`                      | 列出组织       |
| `zt_central_v2_network_groups`            | 网络组列表     |
| `zt_central_v2_network_group`             | 网络组详情     |
| `zt_central_v2_networks`                  | 网络列表       |
| `zt_central_v2_network`                   | 网络详情       |
| `zt_central_v2_network_members`           | 成员列表       |
| `zt_central_v2_network_member`            | 成员详情       |
| `zt_central_v2_authorize_member`          | 授权成员       |
| `zt_central_v2_deauthorize_member`        | 取消授权       |
| `zt_central_v2_create_network`            | 创建网络       |
| `zt_central_v2_delete_network`            | 删除网络       |
| `zt_central_v2_create_network_group`      | 创建网络组     |

### 批量与排行榜（Central v1）

| 工具                         | 描述                    |
| ---------------------------- | ----------------------- |
| `zt_central_batch_authorize`     | 批量授权成员            |
| `zt_central_batch_deauthorize`   | 批量取消授权            |
| `zt_central_batch_delete`        | 批量删除成员            |
| `zt_central_batch_rename`        | 批量重命名（支持模板）  |
| `zt_central_batch_set_ip`        | 批量设置 IP             |
| `zt_central_members_rank`        | 成员排行榜/按字段排序   |

## Claude Desktop 配置

编辑 `claude_desktop_config.json`（或 Cursor / Cherry Studio 等对应配置）：

```json
{
  "mcpServers": {
    "zerotier": {
      "command": "/path/to/zerotier-mcp",
      "env": {
        "ZT_CENTRAL_TOKEN": "your_api_token",
        "ZT_CENTRAL_V2_TOKEN": "your_service_account_token",
        "ZT_LOCAL_TOKEN": "your_local_token"
      }
    }
  }
}
```

## 构建 MCP 服务

```bash
cd cmd/mcp
go build -o zerotier-mcp
```

## 本地构建

使用 Makefile 快速构建：

```bash
# 构建所有平台
make build-mcp

# 构建特定平台
make build-mcp-linux
make build-mcp-darwin
make build-mcp-windows
```

## 环境变量

| 变量                  | 说明                       | 必需 |
| --------------------- | -------------------------- | ---- |
| `ZT_CENTRAL_TOKEN`    | Central v1 API Token       | 否   |
| `ZT_CENTRAL_V2_TOKEN` | Central V2 Service Account | 否   |
| `ZT_LOCAL_TOKEN`      | 本地节点 API Token         | 否   |

## 故障排除

### 权限错误

```bash
chmod +x zerotier-mcp-linux-amd64
```

### 验证完整性

```bash
sha256sum -c SHA256SUMS
```

### 测试连接

```bash
export ZT_CENTRAL_TOKEN="your_token"
export ZT_CENTRAL_V2_TOKEN="your_v2_token"
export ZT_LOCAL_TOKEN="your_local_token"
./zerotier-mcp-linux-amd64
# 按 Ctrl+C 退出
```
