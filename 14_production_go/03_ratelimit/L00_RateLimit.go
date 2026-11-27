package main

import (
	"fmt"
	"sync"
	"time"
)

/*
[主题]: 限流（令牌桶简化版）
[难度]: Medium
[前置]: 03_concurrency/04_mutex

一、主题说明
保护自身与下游。本课实现简易固定窗口计数限流（教学用）；生产可用 golang.org/x/time/rate。

二、写法对比
1. 直观版: 无限制全放行。
2. 进阶版: 窗口内最多 N 次。

三、真实用途
- 网关、登录接口、昂贵 API。
资料: https://pkg.go.dev/golang.org/x/time/rate ；https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
*/
type limiterNaive struct{}

func (limiterNaive) Allow() bool { return true }

type fixedWindow struct {
	mu    sync.Mutex
	limit int
	count int
	start time.Time
	win   time.Duration
}

func newFixedWindow(limit int, win time.Duration) *fixedWindow {
	return &fixedWindow{limit: limit, win: win, start: time.Now()}
}

func (f *fixedWindow) Allow() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	now := time.Now()
	if now.Sub(f.start) >= f.win {
		f.start = now
		f.count = 0
	}
	if f.count >= f.limit {
		return false
	}
	f.count++
	return true
}

func runTests() {
	if !(limiterNaive{}).Allow() {
		panic("naive")
	}
	l := newFixedWindow(2, time.Second)
	if !l.Allow() || !l.Allow() || l.Allow() {
		panic("第三应拒绝")
	}
	fmt.Println("✅ RateLimit 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 RateLimit ************\n")
	runTests()
}

func main() {
	bootstrap()
}
