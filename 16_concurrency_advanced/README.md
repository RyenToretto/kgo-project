# `16_concurrency_advanced` | 并发深挖

- **目标**: pipeline、errgroup、扇入扇出、陷阱、信号量
- **前置**: `03_concurrency`、`14_production_go`
- **环境**: `./scripts/verify_modules.sh 16_concurrency_advanced`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/blog/pipelines
  - https://pkg.go.dev/golang.org/x/sync/errgroup
  - https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197
  - https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_pipeline/L00_Pipelines.go](./00_pipeline/L00_Pipelines.go) | pipeline | ✅ |
| `01_` | [01_errgroup/L00_Errgroup.go](./01_errgroup/L00_Errgroup.go) | errgroup 心智 | ✅ |
| `02_` | [02_fan/L00_FanInFanOut.go](./02_fan/L00_FanInFanOut.go) | 扇入扇出 | ✅ |
| `03_` | [03_traps/L00_ConcurrencyTraps.go](./03_traps/L00_ConcurrencyTraps.go) | 并发陷阱 | ✅ |
| `04_` | [04_semaphore/L00_SemaphorePattern.go](./04_semaphore/L00_SemaphorePattern.go) | 信号量 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| pipelines | https://go.dev/blog/pipelines |
| errgroup | https://pkg.go.dev/golang.org/x/sync/errgroup |
| semaphore | https://pkg.go.dev/golang.org/x/sync/semaphore |
| 书籍 | https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197 |
| loopvar | https://go.dev/blog/loopvar-preview |

**本目录总课数：5（已完成 5，待生成 0）**
