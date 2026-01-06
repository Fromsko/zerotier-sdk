---
inclusion: fileMatch
fileMatchPattern: "**/*_test.go"
---

# 测试开发规范

## 测试文件结构

```go
package xxx_test

import (
    "testing"
    "github.com/fromsko/zerotier-sdk/xxx"
)

func TestXxx(t *testing.T) {
    // 测试代码
}
```

## 表驱动测试

```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:  "正常情况",
            input: "test",
            want:  "expected",
        },
        {
            name:    "错误情况",
            input:   "",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Function(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("got = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Mock HTTP 服务

```go
func setupTestServer(t *testing.T) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path {
        case "/status":
            json.NewEncoder(w).Encode(map[string]interface{}{
                "address": "1234567890",
                "online":  true,
            })
        default:
            w.WriteHeader(http.StatusNotFound)
        }
    }))
}

func TestClient(t *testing.T) {
    server := setupTestServer(t)
    defer server.Close()

    client := xxx.New(xxx.WithBaseURL(server.URL))
    // 测试...
}
```

## 运行测试

```bash
# 运行所有测试
make test

# 运行特定包测试
go test -v ./client/...

# 运行特定测试
go test -v -run TestFunction ./...

# 带覆盖率
go test -cover ./...
```
