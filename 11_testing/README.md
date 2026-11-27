# `11_testing` | 测试

- **目标**: 表驱动、benchmark 心智、接口 mock、httptest
- **前置**: `00`～`05_`（可随时穿插）
- **环境**: Go 最新稳定版（建议 ≥ 1.22）；`./scripts/verify_modules.sh 11_testing`
- **推荐资料（明文 URL，已核验）**:
  - https://pkg.go.dev/testing
  - https://go.dev/doc/tutorial/add-a-test
  - https://github.com/uber-go/mock

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_table/L00_TableDriven.go](./00_table/L00_TableDriven.go) | 表驱动 | ✅ |
| `01_` | [01_bench/L00_BenchmarkIdea.go](./01_bench/L00_BenchmarkIdea.go) | benchmark 心智 | ✅ |
| `02_` | [02_mock/L00_InterfaceMock.go](./02_mock/L00_InterfaceMock.go) | 接口 mock | ✅ |
| `03_` | [03_http_test/L00_HttpTest.go](./03_http_test/L00_HttpTest.go) | httptest | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| testing 包 | https://pkg.go.dev/testing |
| 加测试教程 | https://go.dev/doc/tutorial/add-a-test |
| mock | https://github.com/uber-go/mock |

**本目录总课数：4（已完成 4，待生成 0）**