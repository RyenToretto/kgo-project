package main

import (
	"fmt"
	"strings"
)

/*
[主题]: ADR（架构决策记录）检查清单
[难度]: Easy
[前置]: 无（架构收官）

一、主题说明
重要技术决策写成 ADR：背景、决策、后果。便于后人理解「为什么」。

二、写法对比
1. 直观版: 口头约定。
2. 进阶版: 结构化 ADR 文本。

三、真实用途
- 技术委员会、团队知识沉淀。
资料: https://github.com/joelparkerhenderson/architecture-decision-record
*/
func decideOrally() string { return "我们用 Postgres" }

func renderADR(title, status, context, decision, consequences string) string {
	return fmt.Sprintf("# %s\nStatus: %s\n\n## Context\n%s\n\n## Decision\n%s\n\n## Consequences\n%s\n",
		title, status, context, decision, consequences)
}

func runTests() {
	if decideOrally() == "" {
		panic("oral")
	}
	doc := renderADR("ADR-001 Use Postgres", "Accepted", "Need relational TX", "Use Postgres", "Ops cost; strong consistency")
	if !strings.Contains(doc, "## Decision") || !strings.Contains(doc, "Postgres") {
		panic("adr")
	}
	fmt.Println("✅ AdrChecklist 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 AdrChecklist ************\n")
	fmt.Println(renderADR("ADR-001", "Accepted", "ctx", "dec", "cons"))
	runTests()
}

func main() {
	bootstrap()
}
