package main

import (
	"context"
	"fmt"
	"time"
)

/*
[主题]: 超时与 Deadline
[难度]: Medium
[前置]: 03_concurrency/03_context

一、主题说明
每个出站调用都应有超时；用 context.WithTimeout 贯穿调用链。
超时错误要可识别，便于监控与重试策略。

二、写法对比
1. 直观版: 无超时阻塞。
2. 进阶版: WithTimeout 包裹。

三、真实用途
- DB/Redis/HTTP/gRPC 客户端。
资料: https://go.dev/blog/context ；https://go.dev/doc/security/best-practices
*/
func blockNaive(d time.Duration) {
	time.Sleep(d)
}

func callWithTimeout(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func runTests() {
	blockNaive(5 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	if err := callWithTimeout(ctx, 5*time.Millisecond); err != nil {
		panic("短工作不应超时")
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2()
	if err := callWithTimeout(ctx2, 50*time.Millisecond); err == nil {
		panic("长工作应超时")
	}
	fmt.Println("✅ TimeoutsAndDeadlines 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 TimeoutsAndDeadlines ************\n")
	runTests()
}

func main() {
	bootstrap()
}
