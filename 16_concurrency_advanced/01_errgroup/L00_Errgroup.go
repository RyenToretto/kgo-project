package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

/*
[主题]: errgroup 心智（标准库迷你实现）
[难度]: Medium
[前置]: 03_concurrency/03_context

一、主题说明
errgroup：一组 goroutine，任一失败则 cancel 其余并返回首个 error。
真实库：golang.org/x/sync/errgroup。本课手写迷你版讲清心智。

二、写法对比
1. 直观版: WaitGroup + 忽略错误。
2. 进阶版: 迷你 errgroup。

三、真实用途
- 并行拉取多下游；聚合错误。
资料: https://pkg.go.dev/golang.org/x/sync/errgroup ；https://go.dev/blog/pipelines
*/
type miniGroup struct {
	cancel context.CancelFunc
	wg     sync.WaitGroup
	errOnce sync.Once
	err    error
}

func withContext(parent context.Context) (*miniGroup, context.Context) {
	ctx, cancel := context.WithCancel(parent)
	return &miniGroup{cancel: cancel}, ctx
}

func (g *miniGroup) Go(fn func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := fn(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				g.cancel()
			})
		}
	}()
}

func (g *miniGroup) Wait() error {
	g.wg.Wait()
	g.cancel()
	return g.err
}

func runAllNaive(fns ...func() error) {
	var wg sync.WaitGroup
	for _, fn := range fns {
		wg.Add(1)
		go func(f func() error) {
			defer wg.Done()
			_ = f()
		}(fn)
	}
	wg.Wait()
}

func runTests() {
	runAllNaive(func() error { return nil })
	g, ctx := withContext(context.Background())
	g.Go(func() error {
		<-ctx.Done()
		return nil
	})
	g.Go(func() error { return errors.New("boom") })
	if err := g.Wait(); err == nil || err.Error() != "boom" {
		panic("应返回 boom")
	}
	fmt.Println("✅ Errgroup 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Errgroup ************\n")
	runTests()
}

func main() {
	bootstrap()
}
