# `04_stdlib_http` | 标准库 HTTP

- **目标**: Handler、JSON API、中间件、Client 超时
- **前置**: `00`～`03`（Context 建议先学）
- **环境**: `./scripts/verify_modules.sh 04_stdlib_http`
- **推荐资料（明文 URL，已核验）**:
  - https://pkg.go.dev/net/http
  - https://gobyexample.com/http-servers
  - https://pkg.go.dev/encoding/json

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_handler/L00_HttpHandler.go](./00_handler/L00_HttpHandler.go) | ServeMux / Handler | ✅ |
| `01_` | [01_json/L00_JsonApi.go](./01_json/L00_JsonApi.go) | JSON REST 心智 | ✅ |
| `02_` | [02_middleware/L00_Middleware.go](./02_middleware/L00_Middleware.go) | 中间件包装 | ✅ |
| `03_` | [03_client/L00_HttpClient.go](./03_client/L00_HttpClient.go) | Client 超时 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| net/http | https://pkg.go.dev/net/http |
| HTTP 例子 | https://gobyexample.com/http-servers |
| JSON | https://pkg.go.dev/encoding/json ；https://gobyexample.com/json |
| Context | https://go.dev/blog/context |

**本目录总课数：4（已完成 4，待生成 0）**
