package main

import (
	"fmt"
	"sync"
)

/*
[主题]: 信号量限并发
[难度]: Medium
[前置]: 03_concurrency/01_channel

一、主题说明
用缓冲 channel 作信号量：获取 token 再干活，结束归还。控制最大并发。

二、写法对比
1. 直观版: 无限制并发。
2. 进阶版: sem channel。

三、真实用途
- 限制对 DB/外部 API 的并发连接数。
资料: https://pkg.go.dev/golang.org/x/sync/semaphore ；https://go.dev/blog/pipelines
*/
func unbounded(n int, fn func()) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}

func withSemaphore(n, limit int, fn func()) {
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			fn()
		}()
	}
	wg.Wait()
}

func runTests() {
	var mu sync.Mutex
	maxInflight, cur := 0, 0
	withSemaphore(20, 3, func() {
		mu.Lock()
		cur++
		if cur > maxInflight {
			maxInflight = cur
		}
		mu.Unlock()
		mu.Lock()
		cur--
		mu.Unlock()
	})
	if maxInflight > 3 {
		panic("并发度应 <= 3")
	}
	unbounded(5, func() {})
	fmt.Println("✅ SemaphorePattern 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 SemaphorePattern ************\n")
	runTests()
}

func main() {
	bootstrap()
}
