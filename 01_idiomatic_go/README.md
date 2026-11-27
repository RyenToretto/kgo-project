# `01_idiomatic_go` | 地道 Go

- **目标**: Effective Go、error、defer、接口、组合、命名
- **前置**: `00_go_basics`
- **环境**: Go ≥ 1.22；`./scripts/verify_modules.sh 01_idiomatic_go`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/doc/effective_go
  - https://go.dev/wiki/CodeReviewComments
  - https://github.com/uber-go/guide/blob/master/style.md
  - https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
  - https://go.dev/blog/go1.13-errors

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_errors/L00_ErrorHandling.go](./00_errors/L00_ErrorHandling.go) | error 包装与 Is | ✅ |
| `01_` | [01_defer/L00_DeferPanicRecover.go](./01_defer/L00_DeferPanicRecover.go) | defer / panic / recover | ✅ |
| `02_` | [02_interfaces/L00_Interfaces.go](./02_interfaces/L00_Interfaces.go) | 隐式接口 | ✅ |
| `03_` | [03_composition/L00_CompositionEmbedding.go](./03_composition/L00_CompositionEmbedding.go) | 嵌入组合 | ✅ |
| `04_` | [04_naming/L00_NamingAndPackages.go](./04_naming/L00_NamingAndPackages.go) | 命名与包直觉 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| Effective Go | https://go.dev/doc/effective_go |
| 错误包装 | https://go.dev/blog/go1.13-errors |
| 评审意见 | https://go.dev/wiki/CodeReviewComments |
| Uber 风格 | https://github.com/uber-go/guide/blob/master/style.md |
| 100 Go Mistakes | https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them |

**本目录总课数：5（已完成 5，待生成 0）**
