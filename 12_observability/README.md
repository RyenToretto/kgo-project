# `12_observability` | 可观测性

- **目标**: pprof、Prometheus 指标心智、OpenTelemetry 追踪心智
- **前置**: `06_`、`11_`
- **环境**: Go 最新稳定版（建议 ≥ 1.22）；`./scripts/verify_modules.sh 12_observability`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/doc/diagnostics
  - https://go.dev/blog/pprof
  - https://prometheus.io/docs/guides/go-application/
  - https://github.com/prometheus/client_golang
  - https://opentelemetry.io/docs/languages/go/

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_pprof/L00_PprofMental.go](./00_pprof/L00_PprofMental.go) | pprof 心智 | ✅ |
| `01_` | [01_metrics/L00_MetricsPrometheus.go](./01_metrics/L00_MetricsPrometheus.go) | 指标心智 | ✅ |
| `02_` | [02_trace/L00_TraceOtelMental.go](./02_trace/L00_TraceOtelMental.go) | 追踪心智 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| 诊断总览 | https://go.dev/doc/diagnostics |
| pprof | https://go.dev/blog/pprof |
| Prometheus Go | https://prometheus.io/docs/guides/go-application/ ；https://github.com/prometheus/client_golang |
| OpenTelemetry Go | https://opentelemetry.io/docs/languages/go/ |

**本目录总课数：3（已完成 3，待生成 0）**