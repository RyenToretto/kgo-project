# `14_production_go` | 生产级工程实践

- **目标**: 配置、优雅退出、超时、限流、错误码
- **前置**: 就业主线 `00`～`13`
- **环境**: `./scripts/verify_modules.sh 14_production_go`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/doc/security/best-practices
  - https://go.dev/blog/context
  - https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
  - https://pkg.go.dev/golang.org/x/time/rate

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_config/L00_ConfigLoading.go](./00_config/L00_ConfigLoading.go) | 环境变量配置 | ✅ |
| `01_` | [01_shutdown/L00_GracefulShutdown.go](./01_shutdown/L00_GracefulShutdown.go) | 优雅退出 | ✅ |
| `02_` | [02_timeout/L00_TimeoutsAndDeadlines.go](./02_timeout/L00_TimeoutsAndDeadlines.go) | 超时 | ✅ |
| `03_` | [03_ratelimit/L00_RateLimit.go](./03_ratelimit/L00_RateLimit.go) | 限流 | ✅ |
| `04_` | [04_errors/L00_ErrorCodes.go](./04_errors/L00_ErrorCodes.go) | API 错误码 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| 安全实践 | https://go.dev/doc/security/best-practices |
| Context | https://go.dev/blog/context |
| 100 Go Mistakes | https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them |
| rate | https://pkg.go.dev/golang.org/x/time/rate |

**本目录总课数：5（已完成 5，待生成 0）**
