package main

import (
	"fmt"
	"time"
)

/*
[主题]: select 多路复用
[难度]: Easy
[前置]: 03_concurrency/01_channel

一、主题说明
select 同时等待多个 channel 操作；可配 default 非阻塞；常与 time.After 做超时。

二、写法对比
1. 直观版: 只读一个 channel（可能永久阻塞）。
2. 进阶版: select + 超时。

三、真实用途
- RPC 超时、心跳、多路事件合并。
资料: https://go.dev/tour/concurrency ；https://gobyexample.com/select
*/
func onlyWait(ch <-chan string) string {
	return <-ch
}

func waitWithTimeout(ch <-chan string, d time.Duration) string {
	select {
	case v := <-ch:
		return v
	case <-time.After(d):
		return "timeout"
	}
}

func runTests() {
	ch := make(chan string, 1)
	ch <- "ok"
	if onlyWait(ch) != "ok" {
		panic("onlyWait")
	}
	slow := make(chan string)
	if waitWithTimeout(slow, 30*time.Millisecond) != "timeout" {
		panic("应超时")
	}
	fast := make(chan string, 1)
	fast <- "ready"
	if waitWithTimeout(fast, time.Second) != "ready" {
		panic("应读到 ready")
	}
	fmt.Println("✅ Select 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Select ************\n")
	fmt.Println(waitWithTimeout(make(chan string), 20*time.Millisecond))
	runTests()
}

func main() {
	bootstrap()
}
