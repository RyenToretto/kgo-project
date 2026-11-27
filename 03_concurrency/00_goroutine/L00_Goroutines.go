package main

import (
	"fmt"
	"sync"
	"time"
)

/*
[主题]: goroutine 启动与等待
[难度]: Easy
[前置]: 00～02

一、主题说明
go f() 启动轻量并发。main 退出会结束所有 goroutine——必须用 WaitGroup 等同步等待。

二、写法对比
1. 直观版: sleep 碰运气（脆弱）。
2. 进阶版: sync.WaitGroup 正确等待。

三、真实用途
- 并发请求、后台任务；泄漏（忘 Wait）是线上常见问题。
资料: https://go.dev/tour/concurrency ；https://gobyexample.com/goroutines
*/
func runWithSleep(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		go func() { count++ }() // 故意有 race，仅演示「不可靠」
	}
	time.Sleep(20 * time.Millisecond)
	return count // 结果不确定，测试不用它断言正确性
}

func runWithWaitGroup(n int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	count := 0
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return count
}

func runTests() {
	if runWithWaitGroup(10) != 10 {
		panic("WaitGroup 版本应精确得到 10")
	}
	_ = runWithSleep(5) // 仅演示，不断言
	fmt.Println("✅ Goroutines 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Goroutines ************\n")
	fmt.Println("wg count=", runWithWaitGroup(5))
	runTests()
}

func main() {
	bootstrap()
}
