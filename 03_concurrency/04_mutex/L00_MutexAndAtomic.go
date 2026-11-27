package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
[主题]: Mutex 与原子操作入门
[难度]: Easy
[前置]: 03_concurrency/00_goroutine

一、主题说明
共享内存需要同步。Mutex 保护临界区；简单计数可用 atomic。
优先用 channel 通信；共享状态时再上锁。

二、写法对比
1. 直观版: 无锁并发写（错误）。
2. 进阶版: Mutex / atomic。

三、真实用途
- 连接池计数、本地缓存、限流器内部状态。
资料: https://gobyexample.com/mutexes ；https://pkg.go.dev/sync
*/
func unsafeCount(n int) int {
	count := 0
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			count++ // data race
		}()
	}
	wg.Wait()
	return count
}

func mutexCount(n int) int {
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

func atomicCount(n int) int64 {
	var v int64
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&v, 1)
		}()
	}
	wg.Wait()
	return v
}

func runTests() {
	if mutexCount(100) != 100 {
		panic("mutexCount")
	}
	if atomicCount(100) != 100 {
		panic("atomicCount")
	}
	_ = unsafeCount(50) // 故意不断言
	fmt.Println("✅ MutexAndAtomic 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 MutexAndAtomic ************\n")
	fmt.Println(mutexCount(20), atomicCount(20))
	runTests()
}

func main() {
	bootstrap()
}
