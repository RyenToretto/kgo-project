package main

import (
	"fmt"
	"strings"
)

/*
[主题]: defer / panic / recover
[难度]: Easy
[前置]: 01_idiomatic_go/00_errors

一、主题说明
defer 保证清理。panic 表示不可恢复的程序错误；业务失败优先返回 error。
recover 只能在 defer 中抓住同 goroutine 的 panic。

二、写法对比
1. 直观版: 忘记清理导致「资源」未释放（用切片模拟）。
2. 进阶版: defer 清理 + 可选 recover 隔离 panic。

三、真实用途
- 文件/连接 Close；中间件兜底把 panic 转 500（仍应少用 panic）。
资料: https://go.dev/doc/effective_go ；https://go.dev/wiki/CodeReviewComments
*/
func leaky(log *[]string) {
	*log = append(*log, "open")
	*log = append(*log, "work")
	// 忘记 close
}

func withDefer(log *[]string) {
	*log = append(*log, "open")
	defer func() { *log = append(*log, "close") }()
	*log = append(*log, "work")
}

func safeRun(fn func()) (recovered string) {
	defer func() {
		if r := recover(); r != nil {
			recovered = fmt.Sprint(r)
		}
	}()
	fn()
	return ""
}

func runTests() {
	var a []string
	leaky(&a)
	if strings.Join(a, ",") != "open,work" {
		panic("leaky 轨迹错误")
	}
	var b []string
	withDefer(&b)
	if strings.Join(b, ",") != "open,work,close" {
		panic("defer 应在返回前 close")
	}
	msg := safeRun(func() { panic("boom") })
	if msg != "boom" {
		panic("recover 应抓住 boom")
	}
	if safeRun(func() {}) != "" {
		panic("无 panic 时 recovered 应为空")
	}
	fmt.Println("✅ DeferPanicRecover 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 DeferPanicRecover ************\n")
	fmt.Println("recovered=", safeRun(func() { panic("demo") }))
	runTests()
}

func main() {
	bootstrap()
}
