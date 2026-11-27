# kgo-project 阶段路线

## 学习者画像

- **起点**：Go 小白（可有其他语言基础，但按零基础讲 Go）。
- **目标**：先达到**可就业 Go 后端**，再沿 `14→16→17→15→18→19` 冲刺各方向顶尖精通。
- **工具链**：Go 最新稳定版（建议 ≥ 1.22）。
- **与 kjv 的差异**：目录编号可对齐（便于跨语言对照），但精通顺序按 Go 语言特性定制，**不要**照搬 Java 的 `14→17→16→15→18→19`。

## 阶段编号（学习顺序）

| 编号 | 目录 | 内容 |
| --- | --- | --- |
| `00_` | `00_go_basics` | 安装、语法、类型、控制流、函数、包 |
| `01_` | `01_idiomatic_go` | Effective Go、error、defer、接口、组合 |
| `02_` | `02_data_structures` | slice/map/string/rune、泛型入门 |
| `03_` | `03_concurrency` | goroutine、channel、select、Context |
| `04_` | `04_stdlib_http` | net/http、JSON、中间件心智 |
| `05_` | `05_modules_tooling` | go mod、go test、gofmt/vet/静态检查 |
| `06_` | `06_web_gin` | Gin REST、绑定校验、分层 |
| `07_` | `07_persistence` | database/sql、GORM、事务 |
| `08_` | `08_redis` | go-redis、缓存、锁入门 |
| `09_` | `09_message_queue` | 异步、Kafka/NATS 心智、幂等 |
| `10_` | `10_grpc_microservices` | gRPC、服务拆分、网关心智 |
| `11_` | `11_testing` | testing、表驱动、benchmark、mock |
| `12_` | `12_observability` | pprof、Prometheus、OpenTelemetry |
| `13_` | `13_project_capstone` | 综合小项目（鉴权+CRUD+缓存） |
| `14_` | `14_production_go` | 配置、优雅退出、超时、限流、错误码 |
| `15_` | `15_runtime_advanced` | 调度器、GC、逃逸、诊断 |
| `16_` | `16_concurrency_advanced` | pipeline、errgroup、并发陷阱 |
| `17_` | `17_data_infra` | SQL/索引/隔离、Redis 隐患、MQ 语义 |
| `18_` | `18_cloud_native` | 容器、K8s client、探针、配置 |
| `19_` | `19_architecture` | DDD、事件、CQRS、ADR、SRE |

精通推荐学习顺序（与编号不同，Go 专用）：`14 → 16 → 17 → 15 → 18 → 19`。

| 顺序 | 目录 | 为何排在这里 |
| --- | --- | --- |
| 1 | `14_production_go` | 先能稳定上线（超时、优雅退出、配置） |
| 2 | `16_concurrency_advanced` | Go 核心卖点，紧接生产实践加深 |
| 3 | `17_data_infra` | 数据/缓存/MQ 生产事故高发区 |
| 4 | `15_runtime_advanced` | 调度器/GC：写过大量并发后再学更有体感 |
| 5 | `18_cloud_native` | 容器与编排，偏部署与协同 |
| 6 | `19_architecture` | 架构与 SRE，收官 |

## 全自动生成批次

1. 一次只做一个 `NN_xxx`。
2. 写课 → 更新阶段 README → verify → 单独 commit。
3. 用户要求再 push。
