# `07_persistence` | 持久层

- **目标**: SQL 心智、`database/sql` 模式、事务、GORM 对照（内存实现，无真实 DB）
- **前置**: `06_web_gin`
- **环境**: Go 最新稳定版（建议 ≥ 1.22）；`./scripts/verify_modules.sh 07_persistence`
- **推荐资料（明文 URL，已核验）**:
  - https://pkg.go.dev/database/sql
  - https://github.com/go-sql-driver/mysql
  - https://github.com/jackc/pgx
  - https://gorm.io/docs/

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_sql` | `L00_SqlBasics.go` | SQL CRUD 心智（内存表） | ✅ |
| `01_database_sql` | `L00_DatabaseSqlPattern.go` | database/sql 用法模式 | ✅ |
| `02_tx` | `L00_Transactions.go` | 事务与回滚 | ✅ |
| `03_gorm` | `L00_GormMentalModel.go` | ORM 心智对照 GORM | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| database/sql | https://pkg.go.dev/database/sql |
| MySQL 驱动 | https://github.com/go-sql-driver/mysql |
| pgx | https://github.com/jackc/pgx |
| GORM | https://gorm.io/docs/ |

**本目录总课数：4（已完成 4，待生成 0）**