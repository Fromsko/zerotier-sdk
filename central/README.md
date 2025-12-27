# Central - 云端 Central API

管理 ZeroTier Central 云端控制面板（api.zerotier.com）。

## 安装

```go
import "github.com/fromsko/zerotier-sdk/central"
```

## 获取 Token

1. 登录 [my.zerotier.com](https://my.zerotier.com)
2. Account → 创建 API Token

## 快速开始

```go
c := central.New("your_api_token")

networks, _ := c.Networks().List()
for _, n := range networks {
    fmt.Println(n.ID, n.Config.Name)
}
```

## 配置选项

```go
c := central.New("token",
    central.WithBaseURL("https://api.zerotier.com/api/v1"),
    central.WithTimeout(30 * time.Second),
)
```

## API

### 状态

```go
status, _ := c.Status()
// status.APIVersion, status.User.DisplayName
```

### 网络

```go
// 列表
networks, _ := c.Networks().List()

// 详情
network, _ := c.Networks().Get("network_id")

// 创建
config := central.NewNetworkConfig().
    Name("My Network").
    Private(true).
    AddRoute("10.0.0.0/24", nil).
    AddIPPool("10.0.0.1", "10.0.0.254").
    V4AssignMode(true).
    Build()
c.Networks().Create(config)

// 更新
c.Networks().Update("network_id", config)

// 删除
c.Networks().Delete("network_id")
```

### 成员

```go
// 列表
members, _ := c.Networks().Members("network_id").List()

// 授权
c.Networks().Members("network_id").Authorize("member_id")

// 取消授权
c.Networks().Members("network_id").Deauthorize("member_id")

// 更新
config := central.NewMemberConfig().
    Name("my-device").
    Authorized(true).
    IPAssignments("10.0.0.100").
    Build()
c.Networks().Members("network_id").Update("member_id", config)

// 删除
c.Networks().Members("network_id").Delete("member_id")
```

## 速率限制

- 付费用户：100 请求/秒
- 免费用户：20 请求/秒
