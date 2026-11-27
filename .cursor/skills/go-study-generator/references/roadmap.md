# kgo-project 阶段路线

## 学习者画像

- **起点**：Go 小白（可有其他语言基础，但按零基础讲 Go）。
- **目标**：先达到**可就业 Go 后端**，再沿 `14→16→17→15→18→19` 冲刺各方向**顶尖精通**。
- **工具链**：Go 最新稳定版（建议 ≥ 1.22）。
- **与 kjv 的差异**：目录编号可对齐（便于跨语言对照），但精通顺序按 Go 语言特性定制，**不要**照搬 Java 的 `14→17→16→15→18→19`。
- **资料**：每知识点资料包见 [resources.md](resources.md)；阶段 README 须抄录明文 URL。

## 阶段编号（学习顺序）

| 编号 | 目录 | 内容 |
| --- | --- | --- |
| `00_` | `00_go_basics` | 安装、语法、类型、控制流、函数、包、指针 |
| `01_` | `01_idiomatic_go` | Effective Go、error、defer、接口、组合 |
| `02_` | `02_data_structures` | slice/map/string/rune、泛型入门 |
| `03_` | `03_concurrency` | goroutine、channel、select、Context、Mutex、worker |
| `04_` | `04_stdlib_http` | net/http、JSON、中间件心智、Client |
| `05_` | `05_modules_tooling` | go mod、go test、gofmt/vet |
| `06_` | `06_web_gin` | 路由心智、Gin REST、绑定校验、分层 |
| `07_` | `07_persistence` | SQL 心智、database/sql、事务、GORM 心智 |
| `08_` | `08_redis` | 缓存心智、go-redis 对照、锁入门 |
| `09_` | `09_message_queue` | 异步、至少一次与幂等、Kafka/NATS 心智 |
| `10_` | `10_grpc_microservices` | RPC 心智、gRPC/Protobuf、服务拆分 |
| `11_` | `11_testing` | 表驱动、benchmark、mock、HTTP 测试 |
| `12_` | `12_observability` | pprof、Prometheus 心智、OTel 心智 |
| `13_` | `13_project_capstone` | 综合小项目（鉴权+CRUD+缓存） |
| `14_` | `14_production_go` | 配置、优雅退出、超时、限流、错误码 |
| `15_` | `15_runtime_advanced` | 调度器、GC、逃逸、诊断 |
| `16_` | `16_concurrency_advanced` | pipeline、errgroup、扇入扇出、陷阱 |
| `17_` | `17_data_infra` | 索引/隔离、Redis 隐患、MQ 语义、池 |
| `18_` | `18_cloud_native` | 容器、探针、配置、client-go 心智 |
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

## 课表明细（生成时按此落盘；可增不可随意删核心课）

### `00_go_basics`（6+）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_hello` | `L00_HelloAndTypes.go` | main、基本类型、自测结构 |
| `01_control` | `L00_ControlFlow.go` | if / switch / for / range |
| `02_functions` | `L00_Functions.go` | 多返回值、命名返回、闭包、defer 入门 |
| `03_structs` | `L00_StructsAndMethods.go` | 结构体、方法、值/指针接收者 |
| `04_packages` | `L00_PackagesAndVisibility.go` | 包、导出、init 心智 |
| `05_pointers` | `L00_PointersAndNil.go` | 指针、nil、避免空指针 |

### `01_idiomatic_go`（5）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_errors` | `L00_ErrorHandling.go` | error 接口、wrapping、哨兵 |
| `01_defer` | `L00_DeferPanicRecover.go` | defer / panic / recover |
| `02_interfaces` | `L00_Interfaces.go` | 隐式实现、小接口 |
| `03_composition` | `L00_CompositionEmbedding.go` | 嵌入与组合优于继承 |
| `04_naming` | `L00_NamingAndPackages.go` | 命名、包设计直觉 |

### `02_data_structures`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_slices` | `L00_Slices.go` | 切片底层、append、共享数组陷阱 |
| `01_maps` | `L00_Maps.go` | map、零值、遍历 |
| `02_strings` | `L00_StringsAndRunes.go` | string / byte / rune |
| `03_generics` | `L00_GenericsIntro.go` | 类型参数、约束入门 |

### `03_concurrency`（6）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_goroutine` | `L00_Goroutines.go` | 启动、泄漏直觉 |
| `01_channel` | `L00_Channels.go` | 有缓冲/无缓冲、关闭 |
| `02_select` | `L00_Select.go` | select 多路复用 |
| `03_context` | `L00_Context.go` | 取消、超时、传值边界 |
| `04_mutex` | `L00_MutexAndAtomic.go` | Mutex、原子操作入门 |
| `05_worker` | `L00_WorkerPool.go` | worker pool 模式 |

### `04_stdlib_http`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_handler` | `L00_HttpHandler.go` | ServeMux、Handler |
| `01_json` | `L00_JsonApi.go` | encoding/json REST 心智 |
| `02_middleware` | `L00_Middleware.go` | 中间件链式 |
| `03_client` | `L00_HttpClient.go` | Client、超时 |

### `05_modules_tooling`（3）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_modules` | `L00_ModulesMentalModel.go` | module 路径、依赖心智 |
| `01_gotest` | `L00_GoTestIntro.go` | testing 包心智（可 go run 演示） |
| `02_fmt_vet` | `L00_FmtVetStatic.go` | gofmt / vet 习惯 |

### `06_web_gin`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_router` | `L00_RouterMentalModel.go` | 标准库迷你路由 → 对照 Gin |
| `01_rest` | `L00_RestCrudPattern.go` | REST CRUD 心智 |
| `02_binding` | `L00_BindingValidation.go` | 绑定与校验心智 |
| `03_layers` | `L00_HandlerServiceRepo.go` | 分层（Handler/Service/Repo） |

### `07_persistence`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_sql` | `L00_SqlBasics.go` | SQL CRUD 心智（内存表） |
| `01_database_sql` | `L00_DatabaseSqlPattern.go` | database/sql 用法模式 |
| `02_tx` | `L00_Transactions.go` | 事务与回滚 |
| `03_gorm` | `L00_GormMentalModel.go` | ORM 心智对照 GORM |

### `08_redis`（3）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_cache` | `L00_CacheAside.go` | Cache-Aside |
| `01_client` | `L00_GoRedisMental.go` | 客户端 API 心智 |
| `02_lock` | `L00_DistributedLockIntro.go` | 分布式锁入门与坑 |

### `09_message_queue`（3）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_async` | `L00_AsyncDecouple.go` | 异步解耦 |
| `01_idempotent` | `L00_AtLeastOnceIdempotent.go` | 至少一次 + 幂等 |
| `02_broker` | `L00_BrokerMentalModel.go` | Kafka/NATS 心智对照 |

### `10_grpc_microservices`（3）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_rpc` | `L00_RpcMentalModel.go` | RPC vs HTTP |
| `01_grpc` | `L00_GrpcProtobufMental.go` | gRPC/Protobuf 心智 |
| `02_split` | `L00_ServiceSplit.go` | 拆分边界与网关 |

### `11_testing`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_table` | `L00_TableDriven.go` | 表驱动 |
| `01_bench` | `L00_BenchmarkIdea.go` | benchmark 心智 |
| `02_mock` | `L00_InterfaceMock.go` | 接口 mock |
| `03_http_test` | `L00_HttpTest.go` | httptest |

### `12_observability`（3）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_pprof` | `L00_PprofMental.go` | pprof 心智 |
| `01_metrics` | `L00_MetricsPrometheus.go` | 指标心智 |
| `02_trace` | `L00_TraceOtelMental.go` | 追踪心智 |

### `13_project_capstone`（1+）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_app` | `L00_MiniApiCapstone.go` | 鉴权+CRUD+缓存综合（内存实现可演示） |

### `14_production_go`（5）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_config` | `L00_ConfigLoading.go` | 配置加载 |
| `01_shutdown` | `L00_GracefulShutdown.go` | 优雅退出 |
| `02_timeout` | `L00_TimeoutsAndDeadlines.go` | 超时与 deadline |
| `03_ratelimit` | `L00_RateLimit.go` | 限流 |
| `04_errors` | `L00_ErrorCodes.go` | 错误码与 API 错误体 |

### `16_concurrency_advanced`（5）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_pipeline` | `L00_Pipelines.go` | pipeline |
| `01_errgroup` | `L00_Errgroup.go` | errgroup |
| `02_fan` | `L00_FanInFanOut.go` | 扇入扇出 |
| `03_traps` | `L00_ConcurrencyTraps.go` | 常见并发陷阱 |
| `04_semaphore` | `L00_SemaphorePattern.go` | 信号量 / 限并发 |

### `17_data_infra`（5）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_index` | `L00_SqlIndexExplain.go` | 索引与 EXPLAIN 直觉 |
| `01_isolation` | `L00_IsolationLevels.go` | 隔离级别 |
| `02_redis` | `L00_RedisPitfalls.go` | 穿透/击穿/雪崩 |
| `03_mq` | `L00_MqDeliverySemantics.go` | 投递语义 |
| `04_pool` | `L00_ConnectionPool.go` | 连接池 |

### `15_runtime_advanced`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_scheduler` | `L00_SchedulerGPM.go` | G/P/M 心智 |
| `01_gc` | `L00_GarbageCollector.go` | GC 直觉 |
| `02_escape` | `L00_EscapeAnalysis.go` | 逃逸分析直觉 |
| `03_diag` | `L00_RuntimeDiagnostics.go` | 诊断习惯 |

### `18_cloud_native`（4）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_container` | `L00_ContainerAware.go` | 容器与 GOMAXPROCS |
| `01_probes` | `L00_HealthProbes.go` | 存活/就绪探针 |
| `02_config` | `L00_ConfigAndSecrets.go` | 配置与密钥心智 |
| `03_clientgo` | `L00_ClientGoMental.go` | client-go 心智 |

### `19_architecture`（5）

| 目录 | 文件 | 核心点 |
| --- | --- | --- |
| `00_ddd` | `L00_DddTactical.go` | DDD 战术 |
| `01_events` | `L00_DomainEvents.go` | 领域事件 |
| `02_cqrs` | `L00_CqrsSplit.go` | CQRS |
| `03_adr` | `L00_AdrChecklist.go` | ADR |
| `04_sre` | `L00_SreChecklist.go` | SRE 门禁 |

## 全自动生成批次

1. 一次只做一个 `NN_xxx`。
2. 写课 → 更新阶段 README（进度 + 资料包）→ 更新根 README 状态 → verify → 单独 commit。
3. 用户要求再 push。
4. 生成顺序：`00`→`13`，再 `14→16→17→15→18→19`。
