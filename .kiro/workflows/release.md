# 发布工作流

## 概述

ZeroTier SDK 版本发布的标准工作流。

## 发布前检查

### 1. 代码质量

```bash
# 运行测试
make test

# 代码检查
make lint

# 构建验证
make build
```

### 2. 文档检查

- [ ] README.md 更新
- [ ] CHANGELOG.md 更新
- [ ] API 文档同步

### 3. 版本号

遵循语义化版本：

- `MAJOR.MINOR.PATCH`
- 重大变更：MAJOR
- 新功能：MINOR
- Bug 修复：PATCH

## 发布步骤

### 1. 更新版本号

更新 `mcp/server.go` 中的版本：

```go
s := mcp.New("zerotier", "x.y.z", opts...)
```

### 2. 更新 CHANGELOG

```markdown
## [x.y.z] - YYYY-MM-DD

### 新增

- 新功能描述

### 变更

- 变更描述

### 修复

- 修复描述
```

### 3. 构建 MCP 二进制

```bash
make build-mcp
```

### 4. 创建 Git Tag

```bash
git add .
git commit -m "Release vx.y.z"
git tag vx.y.z
git push origin main --tags
```

### 5. 创建 GitHub Release

- 上传 `dist/` 目录下的二进制文件
- 添加 SHA256 校验和
- 编写 Release Notes

## 发布后

### 1. 验证

- [ ] Go 模块可正常获取
- [ ] MCP 二进制可下载
- [ ] 安装脚本正常工作

### 2. 通知

- 更新相关文档
- 通知用户新版本

## 回滚

如需回滚：

```bash
# 删除本地 tag
git tag -d vx.y.z

# 删除远程 tag
git push origin :refs/tags/vx.y.z
```
