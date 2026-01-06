# Bug 修复工作流

## 概述

ZeroTier SDK Bug 修复的标准工作流。

## 步骤

### 1. 问题定位

- [ ] 复现问题
- [ ] 确定影响范围
- [ ] 定位问题代码

### 2. 编写测试

先编写能复现 Bug 的测试：

```go
func TestBugXxx(t *testing.T) {
    // 复现 Bug 的测试用例
    // 修复前应该失败
}
```

### 3. 修复代码

- 最小化修改
- 不引入新问题
- 保持向后兼容

### 4. 验证修复

```bash
# 运行相关测试
go test -v -run TestBugXxx ./...

# 运行所有测试
make test
```

### 5. 更新文档

- 更新 CHANGELOG
- 必要时更新 README

### 6. 提交

```bash
git add .
git commit -m "fix: 修复问题描述"
```

## 检查清单

- [ ] 问题已复现
- [ ] 测试已添加
- [ ] 修复已验证
- [ ] 无回归问题
- [ ] 文档已更新
