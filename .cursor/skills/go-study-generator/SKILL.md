---
name: go-study-generator
description: >-
  Generate structured Go learning modules for kgo-project (00_go_basics …
  19_architecture). Produces NN_ dirs, L00_*.go lessons with Chinese comments,
  runTests()+bootstrap()+main, and scripts/verify_modules.sh. Use when initializing
  or expanding any NN_xxx stage, running full-auto roadmap generation, or auditing lessons.
---

# Go Study Generator (`kgo-project`)

为 `kgo-project` 生成可直接运行的 Go 学习课：中文注释、直观版 + 进阶版、`runTests` + `bootstrap` + `main`、`./scripts/verify_modules.sh` 可过。

## 何时启用

- 「按 go-study-generator / kgo 学习规范初始化 `00_go_basics` / `06_web_gin` …」
- 「全自动生成下一阶段 / 把路线图全部做完 / 推进 14→16→17→15→18→19」
- 在 `kgo-project` 内新增编号学习目录或审计课是否合格

## Workflow

1. **读路线图，定目录名**
   - 根 [README.md](../../../README.md) 与 [references/roadmap.md](references/roadmap.md)。
   - 阶段目录：`<repo-root>/NN_<slug>/`。
   - 起点：**Go 小白**；先就业主线 `00_`～`13_`，再精通线 **`14→16→17→15→18→19`（Go 专用）**。
   - 注意：编号可与 `kjv-project` 对齐方便对照，但精通顺序**不要**照搬 Java 的 `14→17→16→15→18→19`。
   - 工具链：Go 最新稳定版（建议 ≥ 1.22）。

2. **读模板再写代码**
   - 必读 [references/templates.md](references/templates.md)。
   - 必读 [.cursor/rules/kgo-learning-standards.mdc](../../rules/kgo-learning-standards.mdc)。

3. **源码落盘约定**
   - 优先：`NN_slug/<topic>/L00_TopicName.go`（`package main`，verify 可 `go run`）。
   - 文件名与主要教学入口一致；同阶段风格一致。
   - 生产框架阶段：可用标准库迷你实现讲心智，README 写明与 Gin/gRPC/真实中间件对照。

4. **生成内容约束**
   - 顶部主题块、中文注释、有价值时直观版 + 进阶版。
   - `runTests()` 失败须可见失败（`panic`/`log.Fatal`/返回 error 后退出非 0）。
   - `bootstrap()` 演示后调用 `runTests()`；`main` → `bootstrap()`。
   - 离线可测；新依赖写入该阶段或根 `go.mod`。
   - 更新阶段 README 与根 README「已生成」表。
   - 外链必须明文完整 URL，且已验证可访问。

5. **分批提交（用户要求全自动时）**
   - 每完成一个 `NN_xxx`：verify 通过 → 单独 commit；全部完成再 push。
   - 不要 force push；不要无关重构。

6. **验证**

```bash
./scripts/verify_modules.sh 00_go_basics
./scripts/verify_modules.sh 00_go_basics/00_hello
```

## 用户一句话触发示例

```text
按 go-study-generator / kgo 学习规范初始化 00_go_basics
```

```text
全自动把 README 阶段路线里还没生成的目录都做完，分批次提交并 push
```

```text
请帮我全自动推进14 → 16 → 17 → 15 → 18 → 19，保证全面，细节完美
```

## Additional Resources

- [references/templates.md](references/templates.md)
- [references/roadmap.md](references/roadmap.md)
- 仓库 `scripts/verify_modules.sh`
