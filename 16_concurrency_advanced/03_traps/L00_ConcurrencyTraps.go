package main

import (
	"fmt"
	"sync"
)

/*
[主题]: 常见并发陷阱
[难度]: Medium
[前置]: 03_concurrency

一、主题说明
循环变量捕获（Go 1.22 前经典坑；1.22+ per-iteration）、向已关闭 channel 发送、忘记关闭导致泄漏。

二、写法对比
1. 直观版: 展示危险写法并隔离。
2. 进阶版: 正确关闭与拷贝循环变量习惯。

三、真实用途
- Code review 高频；100 Go Mistakes 大量篇幅。
资料: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them ；https://go.dev/blog/loopvar-preview
*/
func captureSafe(n int) []int {
	out := make([]int, 0, n)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < n; i++ {
		wg.Add(1)
		i := i // 显式拷贝，兼容旧心智
		go func() {
			defer wg.Done()
			mu.Lock()
			out = append(out, i)
			mu.Unlock()
		}()
	}
	wg.Wait()
	return out
}

func closeOnceSend() (panicked bool) {
	ch := make(chan int, 1)
	close(ch)
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	ch <- 1
	return false
}

func runTests() {
	got := captureSafe(5)
	if len(got) != 5 {
		panic("应收集 5 个")
	}
	sum := 0
	for _, v := range got {
		sum += v
	}
	if sum != 10 { // 0+1+2+3+4
		panic("循环变量捕获错误")
	}
	if !closeOnceSend() {
		panic("向关闭 channel 发送应 panic")
	}
	fmt.Println("✅ ConcurrencyTraps 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ConcurrencyTraps ************\n")
	runTests()
}

func main() {
	bootstrap()
}
