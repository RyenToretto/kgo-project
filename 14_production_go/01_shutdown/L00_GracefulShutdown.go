package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
[主题]: 优雅退出（Graceful Shutdown）心智
[难度]: Medium
[前置]: 03_concurrency/03_context

一、主题说明
收到 SIGTERM 后停止接新请求、排空进行中工作、再退出。
本课用 context cancel + WaitGroup 模拟，不真监听信号。

二、写法对比
1. 直观版: 直接 os.Exit。
2. 进阶版: cancel → wait inflight → done。

三、真实用途
- HTTP Server.Shutdown、K8s rolling update。
资料: https://go.dev/blog/context ；https://pkg.go.dev/net/http#Server.Shutdown
*/
func exitNaive(running *int) {
	*running = 0
}

func graceful(ctx context.Context, workers int, workMs int) (finished int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Duration(workMs) * time.Millisecond):
				mu.Lock()
				finished++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return finished
}

func runTests() {
	v := 1
	exitNaive(&v)
	if v != 0 {
		panic("naive")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if graceful(ctx, 3, 50) != 0 {
		panic("已取消不应完成长任务")
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()
	if graceful(ctx2, 3, 10) != 3 {
		panic("充足时间应完成 3")
	}
	fmt.Println("✅ GracefulShutdown 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GracefulShutdown ************\n")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	fmt.Println("finished=", graceful(ctx, 2, 5))
	runTests()
}

func main() {
	bootstrap()
}
