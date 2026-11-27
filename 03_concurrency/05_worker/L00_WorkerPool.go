package main

import (
	"fmt"
	"sync"
)

/*
[主题]: Worker Pool 模式
[难度]: Medium
[前置]: 03_concurrency/01_channel

一、主题说明
固定数量 worker 从 jobs channel 取任务，结果写入 results——控制并发度，避免无限 goroutine。

二、写法对比
1. 直观版: 每个任务一个 goroutine。
2. 进阶版: 固定池 + WaitGroup。

三、真实用途
- 批量下载、消息消费、CPU/IO 密集限流。
资料: https://go.dev/blog/pipelines ；https://gobyexample.com/worker-pools
*/
func fanOutNaive(jobs []int, fn func(int) int) []int {
	out := make([]int, len(jobs))
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for i, j := range jobs {
		i, j := i, j
		go func() {
			defer wg.Done()
			out[i] = fn(j)
		}()
	}
	wg.Wait()
	return out
}

func workerPool(jobs []int, workers int, fn func(int) int) []int {
	if workers < 1 {
		workers = 1
	}
	jobCh := make(chan int)
	resCh := make(chan int, len(jobs))
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for j := range jobCh {
				resCh <- fn(j)
			}
		}()
	}
	go func() {
		for _, j := range jobs {
			jobCh <- j
		}
		close(jobCh)
		wg.Wait()
		close(resCh)
	}()
	var out []int
	for v := range resCh {
		out = append(out, v)
	}
	return out
}

func runTests() {
	jobs := []int{1, 2, 3, 4}
	double := func(x int) int { return x * 2 }
	a := fanOutNaive(jobs, double)
	if len(a) != 4 || a[0] != 2 || a[3] != 8 {
		panic("fanOutNaive")
	}
	b := workerPool(jobs, 2, double)
	sum := 0
	for _, v := range b {
		sum += v
	}
	if sum != 20 {
		panic("workerPool sum 应为 20")
	}
	fmt.Println("✅ WorkerPool 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 WorkerPool ************\n")
	fmt.Println(workerPool([]int{1, 2, 3}, 2, func(x int) int { return x + 1 }))
	runTests()
}

func main() {
	bootstrap()
}
