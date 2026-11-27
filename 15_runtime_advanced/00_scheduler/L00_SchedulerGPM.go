package main

import (
	"fmt"
	"runtime"
)

/*
[主题]: 调度器 G/P/M 心智
[难度]: Hard
[前置]: 03_concurrency 、16_concurrency_advanced

一、主题说明
G=goroutine，M=OS thread，P=逻辑处理器（最多 GOMAXPROCS 个运行）。
本课用 API 观察 GOMAXPROCS / NumGoroutine，不改运行时。

二、写法对比
1. 直观版: 「goroutine=线程」误解。
2. 进阶版: 用 GOMAXPROCS 理解并行度上限。

三、真实用途
- 容器 CPU limit、调优阻塞密集服务。
资料: https://go.dev/doc/diagnostics ；https://go.dev/blog/container-aware-gomaxprocs ；https://www.amazon.com/Efficient-Go-Data-Driven-Performance-Optimization/dp/1098105710
*/
func mythThreads() string {
	return "goroutine==OS thread" // 错误心智
}

func describeScheduler() string {
	return fmt.Sprintf("GOMAXPROCS=%d NumCPU=%d NumGoroutine=%d",
		runtime.GOMAXPROCS(0), runtime.NumCPU(), runtime.NumGoroutine())
}

func runTests() {
	if mythThreads() == "" {
		panic("myth")
	}
	s := describeScheduler()
	if runtime.GOMAXPROCS(0) < 1 {
		panic("GOMAXPROCS")
	}
	if s == "" {
		panic("empty")
	}
	fmt.Println("✅ SchedulerGPM 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 SchedulerGPM ************\n")
	fmt.Println(describeScheduler())
	runTests()
}

func main() {
	bootstrap()
}
