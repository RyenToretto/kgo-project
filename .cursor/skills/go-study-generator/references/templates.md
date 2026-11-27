# `kgo-project` 学习文件标准模板

## 1. 单课 `.go` 模板（package main，可被 verify go run）

```go
package main

import "fmt"

/*
[主题]: 一句话标题
[难度]: Easy / Medium / Hard
[前置]: 无 / 00_go_basics/...

一、主题说明
...

二、写法对比
1. 直观版: ...
2. 进阶版: ...

三、真实用途
- 业务后端 / 并发 / 数据访问: ...
*/
func addNaive(left, right int) int {
	// 中文注释
	return left + right
}

func addBetter(left, right int) int {
	return left + right
}

func runTests() {
	if addNaive(1, 2) != 3 {
		panic("1+2 应为 3")
	}
	if addBetter(2, 3) != 5 {
		panic("2+3 应为 5")
	}
	fmt.Println("✅ HelloAndTypes 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ HelloAndTypes ************\n")
	fmt.Println("演示:", addNaive(10, 20))
	runTests()
}

func main() {
	bootstrap()
}
```

## 2. 阶段 `README.md` 模板

```markdown
# `NN_slug` | 中文名

- **目标**: ...
- **前置**: ...
- **环境**: Go 最新稳定版；`./scripts/verify_modules.sh NN_slug`
- **推荐资料（明文 URL）**:
  - https://...

## 章节进度

| 编号 | 文件 | 核心点 | 状态 |
| --- | --- | --- | --- |
| `00_` | ... | ... | ✅ / 📝 |

**本目录总课数：N（已完成 x，待生成 y）**
```

## 3. 命名

- 目录：`00_hello/`、`01_control/`…
- 文件：`L00_HelloAndTypes.go`
