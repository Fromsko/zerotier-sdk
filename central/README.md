# Central - 云端 Central API

管理 ZeroTier Central 云端控制面板。

- `central` 包：旧版 Central API v1（`https://api.zerotier.com/api/v1`）
- `central/v2` 包：新版 Central API v2（`https://central.zerotier.com/api/v2`）

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

## Central V2

```go
import centralv2 "github.com/fromsko/zerotier-sdk/central/v2"

client, _ := centralv2.NewClientWithToken("your_service_account_token")

// 组织
client.ListOrgsWithResponse(context.Background(), nil)

// 网络组
client.ListNetworkGroupsWithResponse(context.Background(), nil)

// 网络
client.ListNetworksWithResponse(context.Background(), "org_or_group_id", nil)
client.CreateNetworkWithResponse(context.Background(), "group_id", nil)

// 成员
client.ListNetworkMembersWithResponse(context.Background(), "network_id", nil)
client.AuthorizeNetworkMemberWithResponse(context.Background(), "network_id", "member_id", nil)
```

### 组织 / 邀请 / 用户 / Token

```go
// 组织
c.Organization().Get()
c.Organization().Members()

// 邀请
c.Invitations().List()
c.Invitations().Create(req)
c.Invitations().Accept(id)
c.Invitations().Decline(id)

// 用户
c.Users().List()
c.Users().Get(id)

// Token
c.Users().Tokens(userID).List()
c.Users().Tokens(userID).Create(name)
c.Users().Tokens(userID).Delete(id)

// 随机 Token
c.RandomToken()
```

详见 `central/org.go`、`central/invitation.go`、`central/user.go`、`central/token.go`。

## 速率限制

- 付费用户：100 请求/秒
- 免费用户：20 请求/秒
