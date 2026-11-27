package main

import (
	"fmt"
	"strings"
)

/*
[主题]: SRE 门禁检查清单
[难度]: Medium
[前置]: 12_observability 、14_production_go

一、主题说明
SLO/错误预算、告警、演练、变更门禁。上线前用清单自检。

二、写法对比
1. 直观版: 「感觉能上」。
2. 进阶版: 可勾选门禁项。

三、真实用途
- 发布评审、SRE 协作。
资料: https://sre.google/books/ ；https://go.dev/doc/security/best-practices
*/
func feelReady() bool { return true }

func releaseGates(hasSLO, hasAlert, hasRollback, hasDash bool) []string {
	var missing []string
	if !hasSLO {
		missing = append(missing, "SLO")
	}
	if !hasAlert {
		missing = append(missing, "alert")
	}
	if !hasRollback {
		missing = append(missing, "rollback")
	}
	if !hasDash {
		missing = append(missing, "dashboard")
	}
	return missing
}

func runTests() {
	if !feelReady() {
		panic("feel")
	}
	miss := releaseGates(true, false, true, true)
	if len(miss) != 1 || miss[0] != "alert" {
		panic("gates")
	}
	if len(releaseGates(true, true, true, true)) != 0 {
		panic("all green")
	}
	joined := strings.Join(releaseGates(false, false, false, false), ",")
	if !strings.Contains(joined, "SLO") {
		panic("list")
	}
	fmt.Println("✅ SreChecklist 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 SreChecklist ************\n")
	fmt.Println("missing=", releaseGates(true, true, false, true))
	runTests()
}

func main() {
	bootstrap()
}
