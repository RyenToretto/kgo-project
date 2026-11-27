# `15_runtime_advanced` | 运行时 / GC / 诊断

- **目标**: 调度器、GC、逃逸、诊断习惯
- **前置**: `16_concurrency_advanced`（推荐按精通线先学 16）
- **环境**: `./scripts/verify_modules.sh 15_runtime_advanced`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/doc/diagnostics
  - https://go.dev/doc/gc-guide
  - https://go.dev/blog/pprof
  - https://www.amazon.com/Efficient-Go-Data-Driven-Performance-Optimization/dp/1098105710

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_scheduler/L00_SchedulerGPM.go](./00_scheduler/L00_SchedulerGPM.go) | G/P/M | ✅ |
| `01_` | [01_gc/L00_GarbageCollector.go](./01_gc/L00_GarbageCollector.go) | GC 直觉 | ✅ |
| `02_` | [02_escape/L00_EscapeAnalysis.go](./02_escape/L00_EscapeAnalysis.go) | 逃逸 | ✅ |
| `03_` | [03_diag/L00_RuntimeDiagnostics.go](./03_diag/L00_RuntimeDiagnostics.go) | 诊断清单 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| 诊断 | https://go.dev/doc/diagnostics |
| GC | https://go.dev/doc/gc-guide |
| pprof | https://go.dev/blog/pprof |
| Efficient Go | https://www.amazon.com/Efficient-Go-Data-Driven-Performance-Optimization/dp/1098105710 |
| container GOMAXPROCS | https://go.dev/blog/container-aware-gomaxprocs |

**本目录总课数：4（已完成 4，待生成 0）**
