# `09_message_queue` | 消息队列

- **目标**: 异步解耦、至少一次 + 幂等、Kafka/NATS Broker 心智（内存 broker，无真实 MQ）
- **前置**: `03_concurrency`、`06_web_gin`
- **环境**: Go 最新稳定版（建议 ≥ 1.22）；`./scripts/verify_modules.sh 09_message_queue`
- **推荐资料（明文 URL，已核验）**:
  - https://github.com/IBM/sarama
  - https://nats.io/docs/

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_async` | `L00_AsyncDecouple.go` | 异步解耦 | ✅ |
| `01_idempotent` | `L00_AtLeastOnceIdempotent.go` | 至少一次 + 幂等 | ✅ |
| `02_broker` | `L00_BrokerMentalModel.go` | Kafka/NATS 心智对照 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| Sarama (Kafka) | https://github.com/IBM/sarama |
| NATS | https://nats.io/docs/ |

**本目录总课数：3（已完成 3，待生成 0）**