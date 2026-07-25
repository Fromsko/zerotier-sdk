# Examples

ZeroTier SDK 使用示例。

## 目录

```
example/
├── main.go          # 统一接口示例
├── client/          # 本地 Service API 示例
│   └── main.go
├── central/         # 云端 Central API v1 示例
│   └── main.go
└── central_v2/      # 云端 Central API v2 示例（可选）
    └── main.go
```

## 运行

```bash
# 统一接口
go run example/main.go

# 本地 API
go run example/client/main.go

# 云端 API
go run example/central/main.go
```

## 统一接口

```go
import zerotier "github.com/fromsko/zerotier-sdk"

// 本地
local := zerotier.NewClient()
local.Networks().List()

// 云端
cloud := zerotier.NewCentral("token")
cloud.Networks().List()
```

## 直接使用子模块

```go
import (
    "github.com/fromsko/zerotier-sdk/client"
    "github.com/fromsko/zerotier-sdk/central"
)

local := client.New()
cloud := central.New("token")
```
