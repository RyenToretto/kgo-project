# kgo-project

Go **从零到顶尖精通**的个人学习仓库。前提：你是 **Go 小白**（可以有其他语言基础，但 Go 从语法讲起）。目录按 `00_`、`01_`… 学习阶段递增。

姊妹项目：`kjv-project`（Java）、`kpy_project`（Python）、`kts-project`（TypeScript）。三者同级目录，可并行，但 **Go 请严格按本仓库阶段顺序**。

## 目录（规划 / 已生成）

| 路径 | 阶段 | 内容 | 状态 |
| --- | --- | --- | --- |
| `00_go_basics/` | 阶段 1 | Go 语法入门 | ✅ 已生成（6 课） |
| `01_idiomatic_go/` | 阶段 2 | 地道 Go（error/接口/组合） | 📝 待生成 |
| `02_data_structures/` | 阶段 2 | slice/map/string 与泛型 | 📝 待生成 |
| `03_concurrency/` | 阶段 3 | goroutine / channel / Context | 📝 待生成 |
| `04_stdlib_http/` | 阶段 4 | 标准库 net/http | 📝 待生成 |
| `05_modules_tooling/` | 阶段 4 | go mod 与工程工具链 | 📝 待生成 |
| `06_web_gin/` | 阶段 5 | Gin Web API | 📝 待生成 |
| `07_persistence/` | 阶段 5 | database/sql 与 GORM | 📝 待生成 |
| `08_redis/` | 阶段 6 | Redis（go-redis） | 📝 待生成 |
| `09_message_queue/` | 阶段 6 | 消息队列入门 | 📝 待生成 |
| `10_grpc_microservices/` | 阶段 7 | gRPC 与微服务心智 | 📝 待生成 |
| `11_testing/` | 穿插 | testing / benchmark / mock | 📝 待生成 |
| `12_observability/` | 阶段 7 | pprof / Prometheus / OTel | 📝 待生成 |
| `13_project_capstone/` | 毕业 | 综合小项目 | 📝 待生成 |
| `14_production_go/` | 精通 | 生产级工程实践 | 📝 待生成 |
| `15_runtime_advanced/` | 精通 | 运行时 / GC / 诊断 | 📝 待生成 |
| `16_concurrency_advanced/` | 精通 | 并发深挖 | 📝 待生成 |
| `17_data_infra/` | 精通 | 数据与中间件深挖 | 📝 待生成 |
| `18_cloud_native/` | 精通 | 云原生 | 📝 待生成 |
| `19_architecture/` | 精通 | 架构与 SRE | 📝 待生成 |

## 环境

- Go **最新稳定版**（建议 ≥ 1.22）。安装：https://go.dev/doc/install
- 验证入门课：

```bash
cd kgo-project
chmod +x scripts/verify_modules.sh
./scripts/verify_modules.sh 00_go_basics
```

## 一、Go 小白到精通：还该学什么（推荐度从高到低）

下表按**推荐度**从高到低；**目录编号按学习顺序**。推荐度 ≠ 学习顺序（例如 Gin 很值钱，但仍必须先学语法与并发）。

判定：必须/强相关就业或生产；可以系统学；有必要（投入产出高于「只刷语法题」）。

**外链一律明文完整 URL（已 HTTP 核验），不要只看锚文字。**

| 排名 | 推荐度 | 主题 / 代表资料（明文 URL） | 目录 | 核心内容 | 为什么有必要 |
| --- | --- | --- | --- | --- | --- |
| 1 | ⭐⭐⭐⭐⭐+ | Gin 官方教程与文档：https://go.dev/doc/tutorial/web-service-gin ；https://gin-gonic.com/en/docs/ ；https://github.com/gin-gonic/gin | `06_web_gin` | REST、绑定、中间件、分层 | **国内 Go 业务岗常见基本盘**；不要零基础直接跳这里。 |
| 2 | ⭐⭐⭐⭐⭐ | 官方入门：https://go.dev/doc/tutorial/getting-started ；交互课：https://go.dev/tour/ ；例子：https://gobyexample.com/ | `00_go_basics` | 类型、控制流、函数、包 | **地基**。不会这个，后面全是抄作业。 |
| 3 | ⭐⭐⭐⭐⭐ | 并发：https://go.dev/tour/concurrency ；https://go.dev/blog/pipelines ；https://go.dev/blog/context | `03_concurrency` | goroutine、channel、select、Context | Go 的核心卖点；面试与高并发服务分水岭。 |
| 4 | ⭐⭐⭐⭐⭐ | 持久化：https://pkg.go.dev/database/sql ；https://github.com/go-sql-driver/mysql ；https://github.com/jackc/pgx ；https://gorm.io/docs/ | `07_persistence` | SQL、驱动、GORM、事务 | 业务系统几乎都落库。 |
| 5 | ⭐⭐⭐⭐☆ | 地道写法：https://go.dev/doc/effective_go ；Uber 风格：https://github.com/uber-go/guide/blob/master/style.md ；评审意见：https://go.dev/wiki/CodeReviewComments | `01_idiomatic_go` | error、接口、组合、命名 | 写出「像 Go」的代码，而不是 C/Java 翻译腔。 |
| 6 | ⭐⭐⭐⭐☆ | 工程工具：https://go.dev/doc/code ；https://go.dev/ref/mod ；https://go.dev/doc/modules/managing-dependencies | `05_modules_tooling` | module、依赖、构建 | 现代 Go 项目的日常。 |
| 7 | ⭐⭐⭐⭐☆ | Redis：https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis | `08_redis` | 缓存、过期、锁入门 | 性能与会话/秒杀标配。 |
| 8 | ⭐⭐⭐⭐ | 标准库 HTTP：https://pkg.go.dev/net/http ；https://gobyexample.com/http-servers | `04_stdlib_http` | Handler、JSON、中间件心智 | 先懂标准库再上框架，排障更稳。 |
| 9 | ⭐⭐⭐⭐ | 测试：https://pkg.go.dev/testing ；https://go.dev/doc/tutorial/add-a-test ；https://github.com/uber-go/mock | `11_testing` | 表驱动、benchmark、mock | 没有测试谈不上精通与协作。 |
| 10 | ⭐⭐⭐⭐ | 可观测：https://go.dev/doc/diagnostics ；https://go.dev/blog/pprof ；https://prometheus.io/docs/guides/go-application/ ；https://github.com/prometheus/client_golang ；https://opentelemetry.io/docs/languages/go/ | `12_observability` | pprof、指标、追踪 | 线上排障与 SRE 基本功。 |
| 11 | ⭐⭐⭐☆ | gRPC：https://grpc.io/docs/languages/go/quickstart/ | `10_grpc_microservices` | Protobuf、服务间调用 | 微服务与内部 RPC 常见。 |
| 12 | ⭐⭐⭐☆ | 消息队列：https://github.com/IBM/sarama ；https://nats.io/docs/ | `09_message_queue` | 异步、削峰、幂等 | 业务复杂后必遇。 |
| 13 | ⭐⭐⭐☆ | 数据结构与泛型：https://gobyexample.com/slices ；https://go.dev/doc/tutorial/generics ；https://go.dev/blog/go1.18 | `02_data_structures` | slice/map、泛型 | 日写最高频；泛型减少样板。 |
| 14 | ⭐⭐⭐ | 综合项目 + 书籍：https://www.gopl.io/ ；书单索引：https://go.dev/wiki/Books | `13_project_capstone` | 鉴权+CRUD+缓存 | **作品集**；《The Go Programming Language》(Donovan/Kernighan) 系统补强。 |
| 15 | ⭐⭐⭐ | 语言规范与包站：https://go.dev/ref/spec ；https://pkg.go.dev/ ；源码：https://github.com/golang/go | 穿插查阅 | 规范、标准库、源码 | 精通期对照用，不必一上来通读。 |
| 16 | ⭐⭐ | 视频频道：https://www.youtube.com/c/Golang | 穿插 | 官方演讲与专题 | 辅助理解并发/运行时；主线仍以手敲课为准。 |
| 17 | ⭐⭐ | 过早追新框架皮毛 / 无基础直接上 K8s 运维细节 | 不单独建目录 | — | 机会成本高；主线完成前不要分心。 |

### 各知识点「资料包」速查（阶段内还会再细化）

| 知识点簇 | 明文资料 URL |
| --- | --- |
| 安装与第一课 | https://go.dev/doc/install ；https://go.dev/doc/tutorial/getting-started ；https://go.dev/tour/ |
| 语法例子 | https://gobyexample.com/ |
| 地道与评审 | https://go.dev/doc/effective_go ；https://go.dev/wiki/CodeReviewComments ；https://github.com/uber-go/guide/blob/master/style.md |
| 模块依赖 | https://go.dev/ref/mod ；https://go.dev/doc/modules/managing-dependencies |
| 并发专题 | https://go.dev/blog/pipelines ；https://go.dev/blog/context ；https://pkg.go.dev/golang.org/x/sync/errgroup |
| Web | https://go.dev/doc/tutorial/web-service-gin ；https://gin-gonic.com/en/docs/ |
| 数据库 | https://pkg.go.dev/database/sql ；https://gorm.io/docs/ ；https://github.com/go-sql-driver/mysql ；https://github.com/jackc/pgx |
| Redis | https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis |
| gRPC | https://grpc.io/docs/languages/go/quickstart/ |
| MQ | https://github.com/IBM/sarama ；https://nats.io/docs/ |
| 测试 | https://pkg.go.dev/testing ；https://go.dev/doc/tutorial/add-a-test ；https://github.com/uber-go/mock |
| 诊断与监控 | https://go.dev/doc/diagnostics ；https://go.dev/blog/pprof ；https://prometheus.io/docs/guides/go-application/ ；https://opentelemetry.io/docs/languages/go/ |
| 安全实践 | https://go.dev/doc/security/best-practices |
| 云原生 Go | https://github.com/kubernetes/client-go ；https://go.dev/blog/container-aware-gomaxprocs |
| 经典书 | https://www.gopl.io/ ；https://go.dev/wiki/Books |
| 官方视频 | https://www.youtube.com/c/Golang |

## 二、阶段路线（学习顺序）

| 阶段 | 建议周期 | 目录 | 做什么 | 验收标准 |
| --- | --- | --- | --- | --- |
| 1. 语法 | 2～4 周 | `00_go_basics` | 每天 1～2 课，手敲 + verify | 能独立写小程序与包结构 |
| 2. 地道与数据 | 2～3 周 | `01_`、`02_` | error/接口 + slice/map/泛型 | 会正确处理 error；说清切片底层直觉 |
| 3. 并发 | 2～4 周 | `03_concurrency` | channel/Context 练习 | 能写 worker pool；理解取消与泄漏 |
| 4. HTTP 与工具 | 2～3 周 | `04_`、`05_`、`11_`（开始） | 标准库服务 + module + 测试 | 能手写简易 API；会 `go test` |
| 5. Web 与存储 | 1～2 月 | `06_`、`07_` | Gin + SQL/GORM CRUD | 独立完成带校验的 REST + DB |
| 6. 缓存与消息 | 2～4 周 | `08_`、`09_` | 接到上阶段项目 | 会缓存查询；理解异步与幂等 |
| 7. 微服务与观测 | 1～2 月 | `10_`、`12_`、`13_` | gRPC + 可观测 + 毕业项目 | 作品集可演示；能看 pprof/指标 |
| 8. 顶尖精通 | 持续 | `14→16→17→15→18→19`（Go 专用，勿照搬 Java） | 生产 → 并发深挖 → 中间件 → 运行时 → 云原生 → 架构 | 各方向能讲清原理并对照真实组件 |

## 三、用 Cursor 全自动生成（rules + skill）

| 类型 | 路径 | 作用 |
| --- | --- | --- |
| Rule | `.cursor/rules/kgo-learning-standards.mdc` | 编号、`L00_`、小白友好、verify、明文链接、顶尖精通终点 |
| Skill | `.cursor/skills/go-study-generator/SKILL.md` | 按阶段生成课、模板、分批提交 |
| 课表 | `.cursor/skills/go-study-generator/references/roadmap.md` | 各阶段课表明细与精通顺序 |
| 资料包 | `.cursor/skills/go-study-generator/references/resources.md` | 已 HTTP 核验的明文 URL（文档/书/视频） |

**用法：** Cursor 打开 **`kgo-project` 根目录**，粘贴下面「全自动 Prompt」。

### 推荐粘贴的「全自动」Prompt（复制整段）

```text
请严格按 go-study-generator skill 与 .cursor/rules/kgo-learning-standards.mdc 执行。

我的起点：Go 小白（可能有一点其他语言基础，但请按零基础讲 Go）。
工具链：使用 Go 最新稳定版（建议 ≥ 1.22）；生成的代码与讲解按当前稳定版，不要用过时写法。
目标：按 README「阶段路线」顺序，把尚未生成的目录全部自动生成完，使我按顺序学完后达到可就业的 Go 后端水平，并继续精通各方向（语法/并发扎实、会 Gin+SQL、会 Redis/MQ 入门、会测试与可观测、有综合小项目，再完成 14→19）。

生成顺序（跳过已完成的课，补齐缺课）：
00_go_basics →
01_idiomatic_go →
02_data_structures →
03_concurrency →
04_stdlib_http →
05_modules_tooling →
06_web_gin →
07_persistence →
08_redis →
09_message_queue →
10_grpc_microservices →
11_testing →
12_observability →
13_project_capstone →
（以下为 Go 精通线，顺序勿改成 Java 的 14→17→16→…）
14_production_go →
16_concurrency_advanced →
17_data_infra →
15_runtime_advanced →
18_cloud_native →
19_architecture

硬性要求：
1. 一次只做一个 NN_xxx；写完更新该阶段 README 与根 README「已生成」表。
2. 每课：L00_Xxx.go 命名、中文注释、直观版+进阶版（有价值时）、runTests()+bootstrap()+main、可用 ./scripts/verify_modules.sh 跑通（package main 入门课优先）。
3. 每阶段 verify 全绿后单独 git commit；全部完成后 git push。
4. 内容必须由浅入深，禁止默认我已会 Gin/gRPC/K8s。
5. 不要无关重构；不要 force push。
6. README 外链必须明文完整 URL，禁止未验证假链接。

现在从下一个未完成的阶段开始全自动执行，直到全部完成。
```

### 精通线追加 Prompt

```text
请帮我全自动推进14 → 16 → 17 → 15 → 18 → 19，保证全面，细节完美。
（这是 Go 专用精通顺序：生产 → 并发深挖 → 数据中间件 → 运行时 → 云原生 → 架构；不要用 Java 仓的 14→17→16→15→18→19。）
严格按 go-study-generator 与 kgo-learning-standards：每阶段 verify 全绿后单独 commit，全部完成再 push，不要 force push。
```

### 为何精通线不是 Java 那套？

| | Java (`kjv`) | Go (`kgo`) |
| --- | --- | --- |
| 编号 `00`～`19` | 阶段目录骨架 | **可对齐**（方便跨语言对照） |
| 精通顺序 | `14→17→16→15→18→19` | **`14→16→17→15→18→19`** |
| 差异原因 | 先巩固数据层再挖 JUC/JVM | Go 并发是语言核心，应早于运行时深挖；运行时在大量并发实践后再学更合适 |

## 四、和 kjv / kpy / kts 的关系

- `kjv-project`：Java 后端主线
- `kgo-project`：Go 后端主线（本仓库）
- `kpy_project`：Python / AI
- `kts-project`：TypeScript / 前端

不要一上来只抄 Gin 教程；请按本 README 阶段顺序推进。
