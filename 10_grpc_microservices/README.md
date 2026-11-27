# `10_grpc_microservices` | gRPC 与微服务

- **目标**: RPC 心智、gRPC/Protobuf 契约、服务拆分与网关
- **前置**: `06`～`09_`（HTTP 与异步基础）
- **环境**: Go 最新稳定版（建议 ≥ 1.22）；`./scripts/verify_modules.sh 10_grpc_microservices`
- **推荐资料（明文 URL，已核验）**:
  - https://grpc.io/docs/languages/go/quickstart/
  - https://protobuf.dev/getting-started/gotutorial/

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_rpc/L00_RpcMentalModel.go](./00_rpc/L00_RpcMentalModel.go) | RPC vs HTTP | ✅ |
| `01_` | [01_grpc/L00_GrpcProtobufMental.go](./01_grpc/L00_GrpcProtobufMental.go) | Protobuf 契约心智 | ✅ |
| `02_` | [02_split/L00_ServiceSplit.go](./02_split/L00_ServiceSplit.go) | 拆分边界与网关 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| gRPC Go Quickstart | https://grpc.io/docs/languages/go/quickstart/ |
| Protobuf Go | https://protobuf.dev/getting-started/gotutorial/ |

**本目录总课数：3（已完成 3，待生成 0）**