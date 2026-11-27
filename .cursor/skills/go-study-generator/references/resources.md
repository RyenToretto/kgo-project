# 已核验资料包（明文 URL）

> 规则：写入 README / 课内注释的外链必须是本表或同次 HTTP 核验通过的完整 URL；禁止 Markdown 仅锚文字、禁止未核验假链接。  
> 核验方式：`curl -sL -o /dev/null -w "%{http_code}" <url>`（可加浏览器 UA）；目标优先 **200**。个别站点对爬虫 403（如部分 O'Reilly 页）时，改用本表中的 **Amazon / Manning / go.dev/wiki/Books** 可核验入口，勿贴死链。  
> 更新日：2026-03-27；本机 Go 建议 ≥ 1.22。

## 总览（小白 → 就业 → 顶尖）

| 用途 | 明文 URL |
| --- | --- |
| 安装 | https://go.dev/doc/install |
| 学习入口 | https://go.dev/learn/ |
| 文档总览 | https://go.dev/doc/ |
| Tour | https://go.dev/tour/ |
| Go by Example | https://gobyexample.com/ |
| Effective Go | https://go.dev/doc/effective_go |
| 语言规范 | https://go.dev/ref/spec |
| 包索引 | https://pkg.go.dev/ |
| 源码 | https://github.com/golang/go |
| 书单索引 | https://go.dev/wiki/Books |
| GOPL 官网 | https://www.gopl.io/ |
| 官方视频频道 | https://www.youtube.com/c/Golang |
| 100 Go Mistakes（Manning） | https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them |
| Concurrency in Go（Amazon） | https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197 |
| Efficient Go（Amazon） | https://www.amazon.com/Efficient-Go-Data-Driven-Performance-Optimization/dp/1098105710 |

## 按阶段知识点资料包

### `00_go_basics`

| 知识点 | 明文 URL |
| --- | --- |
| 安装与 Hello | https://go.dev/doc/install ；https://go.dev/doc/tutorial/getting-started |
| 交互语法 | https://go.dev/tour/ ；https://gobyexample.com/ |
| 写 Go 代码习惯 | https://go.dev/doc/code |
| 书籍补强 | https://www.gopl.io/ （第 1–5 章） |

### `01_idiomatic_go`

| 知识点 | 明文 URL |
| --- | --- |
| Effective Go | https://go.dev/doc/effective_go |
| Code Review Comments | https://go.dev/wiki/CodeReviewComments |
| Uber Go Style Guide | https://github.com/uber-go/guide/blob/master/style.md |
| 常见错误书 | https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them |

### `02_data_structures`

| 知识点 | 明文 URL |
| --- | --- |
| slice | https://gobyexample.com/slices ；https://go.dev/blog/slices-intro |
| map / string | https://gobyexample.com/maps ；https://gobyexample.com/strings-and-runes |
| 泛型 | https://go.dev/doc/tutorial/generics ；https://go.dev/blog/go1.18 |

### `03_concurrency`

| 知识点 | 明文 URL |
| --- | --- |
| Tour 并发 | https://go.dev/tour/concurrency |
| pipelines | https://go.dev/blog/pipelines |
| Context | https://go.dev/blog/context |
| errgroup | https://pkg.go.dev/golang.org/x/sync/errgroup |
| 书籍 | https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197 ；https://go.dev/wiki/Books |

### `04_stdlib_http`

| 知识点 | 明文 URL |
| --- | --- |
| net/http | https://pkg.go.dev/net/http |
| HTTP 例子 | https://gobyexample.com/http-servers |
| JSON | https://pkg.go.dev/encoding/json |

### `05_modules_tooling`

| 知识点 | 明文 URL |
| --- | --- |
| Modules | https://go.dev/ref/mod ；https://go.dev/doc/modules/managing-dependencies |
| 如何写代码 | https://go.dev/doc/code |
| 测试入门 | https://go.dev/doc/tutorial/add-a-test |

### `06_web_gin`

| 知识点 | 明文 URL |
| --- | --- |
| 官方 Gin 教程 | https://go.dev/doc/tutorial/web-service-gin |
| Gin 文档 | https://gin-gonic.com/en/docs/ |
| Gin 仓库 | https://github.com/gin-gonic/gin |

### `07_persistence`

| 知识点 | 明文 URL |
| --- | --- |
| database/sql | https://pkg.go.dev/database/sql |
| MySQL 驱动 | https://github.com/go-sql-driver/mysql |
| pgx | https://github.com/jackc/pgx |
| GORM | https://gorm.io/docs/ |

### `08_redis`

| 知识点 | 明文 URL |
| --- | --- |
| Redis Go 客户端指南 | https://redis.io/docs/latest/develop/clients/go/ |
| go-redis | https://github.com/redis/go-redis |

### `09_message_queue`

| 知识点 | 明文 URL |
| --- | --- |
| Sarama (Kafka) | https://github.com/IBM/sarama |
| NATS | https://nats.io/docs/ |

### `10_grpc_microservices`

| 知识点 | 明文 URL |
| --- | --- |
| gRPC Go Quickstart | https://grpc.io/docs/languages/go/quickstart/ |
| Protobuf Go | https://protobuf.dev/getting-started/gotutorial/ |

### `11_testing`

| 知识点 | 明文 URL |
| --- | --- |
| testing 包 | https://pkg.go.dev/testing |
| 加测试教程 | https://go.dev/doc/tutorial/add-a-test |
| mock | https://github.com/uber-go/mock |

### `12_observability`

| 知识点 | 明文 URL |
| --- | --- |
| 诊断总览 | https://go.dev/doc/diagnostics |
| pprof | https://go.dev/blog/pprof |
| Prometheus Go | https://prometheus.io/docs/guides/go-application/ ；https://github.com/prometheus/client_golang |
| OpenTelemetry Go | https://opentelemetry.io/docs/languages/go/ |

### `13_project_capstone`

| 知识点 | 明文 URL |
| --- | --- |
| Gin 服务教程 | https://go.dev/doc/tutorial/web-service-gin |
| 安全实践 | https://go.dev/doc/security/best-practices |
| GOPL | https://www.gopl.io/ |

### `14_production_go`

| 知识点 | 明文 URL |
| --- | --- |
| 安全最佳实践 | https://go.dev/doc/security/best-practices |
| Context 超时 | https://go.dev/blog/context |
| 100 Go Mistakes | https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them |

### `15_runtime_advanced`

| 知识点 | 明文 URL |
| --- | --- |
| 诊断 | https://go.dev/doc/diagnostics |
| pprof | https://go.dev/blog/pprof |
| GC 指南 | https://go.dev/doc/gc-guide |
| Efficient Go（书） | https://www.amazon.com/Efficient-Go-Data-Driven-Performance-Optimization/dp/1098105710 ；https://go.dev/wiki/Books |

### `16_concurrency_advanced`

| 知识点 | 明文 URL |
| --- | --- |
| pipelines | https://go.dev/blog/pipelines |
| errgroup | https://pkg.go.dev/golang.org/x/sync/errgroup |
| Concurrency in Go（书） | https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197 |

### `17_data_infra`

| 知识点 | 明文 URL |
| --- | --- |
| database/sql | https://pkg.go.dev/database/sql |
| Redis 客户端 | https://redis.io/docs/latest/develop/clients/go/ |
| NATS / Kafka | https://nats.io/docs/ ；https://github.com/IBM/sarama |

### `18_cloud_native`

| 知识点 | 明文 URL |
| --- | --- |
| container-aware GOMAXPROCS | https://go.dev/blog/container-aware-gomaxprocs |
| client-go | https://github.com/kubernetes/client-go |
| K8s 探针 | https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes |

### `19_architecture`

| 知识点 | 明文 URL |
| --- | --- |
| ADR | https://github.com/joelparkerhenderson/architecture-decision-record |
| SRE 书籍 | https://sre.google/books/ |
| 安全实践 | https://go.dev/doc/security/best-practices |

## 视频（穿插，主线仍以本仓库手敲课为准）

| 主题 | 明文 URL |
| --- | --- |
| Go 官方 YouTube | https://www.youtube.com/c/Golang |
