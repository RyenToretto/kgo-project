package main

import "fmt"

/*
[主题]: channel：有缓冲 / 无缓冲、关闭与 range
[难度]: Easy
[前置]: 03_concurrency/00_goroutine

一、主题说明
channel 用于 goroutine 间通信。无缓冲要求收发同时就绪；有缓冲可暂存。
发送方关闭 channel；接收方用 range 或 `v, ok := <-ch`。

二、写法对比
1. 直观版: 无缓冲同步传递。
2. 进阶版: 有缓冲批量 + close + range。

三、真实用途
- 任务队列、结果收集、pipeline。
资料: https://go.dev/tour/concurrency ；https://gobyexample.com/channels
*/
func pingPongNaive() string {
	ch := make(chan string)
	go func() { ch <- "pong" }()
	return <-ch
}

func collectBetter(n int) []int {
	ch := make(chan int, n)
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
	var out []int
	for v := range ch {
		out = append(out, v)
	}
	return out
}

func runTests() {
	if pingPongNaive() != "pong" {
		panic("pingPong 失败")
	}
	got := collectBetter(3)
	if len(got) != 3 || got[0] != 1 || got[2] != 3 {
		panic("collectBetter 失败")
	}
	fmt.Println("✅ Channels 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Channels ************\n")
	fmt.Println(pingPongNaive(), collectBetter(2))
	runTests()
}

func main() {
	bootstrap()
}
