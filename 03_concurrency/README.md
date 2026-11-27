# `03_concurrency` | 并发基础

- **目标**: goroutine、channel、select、Context、Mutex、worker pool
- **前置**: `00`～`02`
- **环境**: Go ≥ 1.22；`./scripts/verify_modules.sh 03_concurrency`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/tour/concurrency
  - https://go.dev/blog/pipelines
  - https://go.dev/blog/context
  - https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_goroutine/L00_Goroutines.go](./00_goroutine/L00_Goroutines.go) | goroutine + WaitGroup | ✅ |
| `01_` | [01_channel/L00_Channels.go](./01_channel/L00_Channels.go) | channel 开关与 range | ✅ |
| `02_` | [02_select/L00_Select.go](./02_select/L00_Select.go) | select 超时 | ✅ |
| `03_` | [03_context/L00_Context.go](./03_context/L00_Context.go) | Context 取消 | ✅ |
| `04_` | [04_mutex/L00_MutexAndAtomic.go](./04_mutex/L00_MutexAndAtomic.go) | Mutex / atomic | ✅ |
| `05_` | [05_worker/L00_WorkerPool.go](./05_worker/L00_WorkerPool.go) | worker pool | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| Tour 并发 | https://go.dev/tour/concurrency |
| pipelines | https://go.dev/blog/pipelines |
| Context | https://go.dev/blog/context ；https://pkg.go.dev/context |
| 例子 | https://gobyexample.com/goroutines ；https://gobyexample.com/channels ；https://gobyexample.com/worker-pools |
| 书籍 | https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197 |

**本目录总课数：6（已完成 6，待生成 0）**
