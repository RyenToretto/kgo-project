# `18_cloud_native` | 云原生

- **目标**: 容器感知、探针、配置密钥、client-go 心智
- **前置**: `14`、`15` 建议
- **环境**: `./scripts/verify_modules.sh 18_cloud_native`
- **推荐资料（明文 URL，已核验）**:
  - https://go.dev/blog/container-aware-gomaxprocs
  - https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
  - https://github.com/kubernetes/client-go
  - https://go.dev/doc/security/best-practices

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | [00_container/L00_ContainerAware.go](./00_container/L00_ContainerAware.go) | GOMAXPROCS | ✅ |
| `01_` | [01_probes/L00_HealthProbes.go](./01_probes/L00_HealthProbes.go) | live/ready | ✅ |
| `02_` | [02_config/L00_ConfigAndSecrets.go](./02_config/L00_ConfigAndSecrets.go) | 配置与密钥 | ✅ |
| `03_` | [03_clientgo/L00_ClientGoMental.go](./03_clientgo/L00_ClientGoMental.go) | client-go 心智 | ✅ |

## 本阶段知识点资料包（明文 URL）

| 知识点 | 明文 URL |
| --- | --- |
| GOMAXPROCS | https://go.dev/blog/container-aware-gomaxprocs |
| 探针 | https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes |
| client-go | https://github.com/kubernetes/client-go |
| 安全 | https://go.dev/doc/security/best-practices |

**本目录总课数：4（已完成 4，待生成 0）**
