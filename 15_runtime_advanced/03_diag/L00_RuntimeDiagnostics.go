package main

import (
	"fmt"
	"runtime"
	"strings"
)

/*
[主题]: 运行时诊断习惯（pprof 入口心智）
[难度]: Medium
[前置]: 12_observability（可对照）

一、主题说明
线上问题：CPU/内存/阻塞。标准做法：net/http/pprof 或 go tool pprof。
本课组装诊断清单字符串，强调先复现再抓 profile。

二、写法对比
1. 直观版: 只看日志猜。
2. 进阶版: 诊断检查清单。

三、真实用途
- 生产排障、性能回归。
资料: https://go.dev/blog/pprof ；https://go.dev/doc/diagnostics
*/
func guessFromLogs() string { return "maybe slow?" }

func diagnosticChecklist() []string {
	return []string{
		"复现路径与流量特征",
		"go tool pprof http://host:port/debug/pprof/profile",
		"heap: /debug/pprof/heap",
		"goroutine: /debug/pprof/goroutine",
		fmt.Sprintf("当前 NumGoroutine=%d", runtime.NumGoroutine()),
	}
}

func runTests() {
	if guessFromLogs() == "" {
		panic("guess")
	}
	items := diagnosticChecklist()
	joined := strings.Join(items, "|")
	if !strings.Contains(joined, "pprof") {
		panic("checklist")
	}
	fmt.Println("✅ RuntimeDiagnostics 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 RuntimeDiagnostics ************\n")
	for _, s := range diagnosticChecklist() {
		fmt.Println("-", s)
	}
	runTests()
}

func main() {
	bootstrap()
}
