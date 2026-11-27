package main

import (
	"context"
	"fmt"
	"time"
)

/*
[主题]: context：取消与超时
[难度]: Medium
[前置]: 03_concurrency/02_select

一、主题说明
Context 在调用链传递截止时间与取消信号。HTTP/gRPC/DB 客户端几乎都接受 ctx。
不要把 ctx 放结构体长期存；作首参传入。

二、写法对比
1. 直观版: 无取消的死循环工作。
2. 进阶版: WithTimeout + Done。

三、真实用途
- 请求级超时、级联取消、追踪元数据（值谨慎使用）。
资料: https://go.dev/blog/context ；https://pkg.go.dev/context
*/
func workForeverNaive(stop <-chan struct{}) string {
	for {
		select {
		case <-stop:
			return "stopped"
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func workWithContext(ctx context.Context) string {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err().Error()
		default:
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func runTests() {
	stop := make(chan struct{})
	go func() {
		time.Sleep(20 * time.Millisecond)
		close(stop)
	}()
	if workForeverNaive(stop) != "stopped" {
		panic("naive stop")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()
	msg := workWithContext(ctx)
	if msg != context.DeadlineExceeded.Error() {
		panic("期望 deadline exceeded, got=" + msg)
	}
	fmt.Println("✅ Context 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Context ************\n")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Millisecond)
	defer cancel()
	fmt.Println(workWithContext(ctx))
	runTests()
}

func main() {
	bootstrap()
}
