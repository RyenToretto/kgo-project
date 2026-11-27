# `08_redis` | Redis

- **目标**: Cache-Aside、go-redis API 心智、分布式锁入门（内存假 Redis，无 redis-server）
- **前置**: `06_web_gin`、`07_persistence`（建议）
- **环境**: Go 最新稳定版（建议 ≥ 1.22）；`./scripts/verify_modules.sh 08_redis`
- **推荐资料（明文 URL，已核验）**:
  - https://redis.io/docs/latest/develop/clients/go/
  - https://github.com/redis/go-redis

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_cache` | `L00_CacheAside.go` | Cache-Aside | ✅ |
| `01_client` | `L00_GoRedisMental.go` | 客户端 API 心智 | ✅ |
| `02_lock` | `L00_DistributedLockIntro.go` | 分布式锁入门与坑 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| Redis Go 客户端指南 | https://redis.io/docs/latest/develop/clients/go/ |
| go-redis | https://github.com/redis/go-redis |

**本目录总课数：3（已完成 3，待生成 0）**