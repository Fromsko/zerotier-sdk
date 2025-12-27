# Client - 本地 Service API

管理本地 ZeroTier 节点（localhost:9993）。

## 安装

```go
import "github.com/fromsko/zerotier-sdk/client"
```

## 快速开始

```go
c := client.New() // 自动读取系统 token

status, _ := c.Status()
fmt.Println(status.Address, status.Online)
```

## 配置选项

```go
c := client.New(
    client.WithBaseURL("http://localhost:9993"),
    client.WithToken("your-token"),
    client.WithTokenFile("/path/to/authtoken.secret"),
    client.WithTimeout(30 * time.Second),
)
```

## Token 位置

| 系统    | 路径                                                      |
| ------- | --------------------------------------------------------- |
| Windows | `C:\ProgramData\ZeroTier\One\authtoken.secret`            |
| macOS   | `~/Library/Application Support/ZeroTier/authtoken.secret` |
| Linux   | `/var/lib/zerotier-one/authtoken.secret`                  |

## API

### 节点状态

```go
status, _ := c.Status()
// status.Address, status.Version, status.Online
```

### 网络

```go
// 列表
networks, _ := c.Networks().List()

// 加入
c.Networks().Join("network_id")

// 离开
c.Networks().Leave("network_id")

// 更新设置
settings := client.NewNetworkSettings().
    AllowDNS(true).
    AllowManaged(true).
    Build()
c.Networks().Update("network_id", settings)
```

### Peers

```go
peers, _ := c.Peers().List()
peer, _ := c.Peers().Get("peer_id")
```

### 控制器（自托管）

```go
// 状态
status, _ := c.Controller().Status()

// 创建网络
config := client.NewControllerNetworkConfig().
    Name("my-network").
    Private(true).
    AddRoute("10.0.0.0/24", nil).
    AddIPPool("10.0.0.1", "10.0.0.254").
    V4AssignMode(true).
    Build()
c.Controller().CreateNetwork(nodeID, config)

// 授权成员
memberConfig := client.NewControllerMemberConfig().
    Authorized(true).
    Build()
c.Controller().UpdateMember(networkID, memberID, memberConfig)
```
