# `17_data_infra` | 数据与中间件深挖

- **目标**: 索引、隔离级别、Redis 隐患、MQ 语义、连接池
- **前置**: `07`/`08`/`09`、`14`
- **环境**: `./scripts/verify_modules.sh 17_data_infra`
- **推荐资料（明文 URL，已核验）**:
  - https://pkg.go.dev/database/sql
  - https://redis.io/docs/latest/develop/clients/go/
  - https://nats.io/docs/
  - https://github.com/IBM/sarama

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_index/L00_SqlIndexExplain.go](./00_index/L00_SqlIndexExplain.go) | 索引直觉 | ✅ |
| `01_` | [01_isolation/L00_IsolationLevels.go](./01_isolation/L00_IsolationLevels.go) | 隔离级别 | ✅ |
| `02_` | [02_redis/L00_RedisPitfalls.go](./02_redis/L00_RedisPitfalls.go) | 缓存三大坑 | ✅ |
| `03_` | [03_mq/L00_MqDeliverySemantics.go](./03_mq/L00_MqDeliverySemantics.go) | 投递语义 | ✅ |
| `04_` | [04_pool/L00_ConnectionPool.go](./04_pool/L00_ConnectionPool.go) | 连接池 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| SQL | https://pkg.go.dev/database/sql ；https://gorm.io/docs/ |
| Redis | https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis |
| MQ | https://nats.io/docs/ ；https://github.com/IBM/sarama |

**本目录总课数：5（已完成 5，待生成 0）**
