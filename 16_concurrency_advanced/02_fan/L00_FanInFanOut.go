package main

import (
	"fmt"
	"sync"
)

/*
[主题]: 扇出 Fan-Out / 扇入 Fan-In
[难度]: Medium
[前置]: 16_concurrency_advanced/00_pipeline

一、主题说明
Fan-Out：多 worker 并行消费同一输入；Fan-In：多 channel 合并到一个。

二、写法对比
1. 直观版: 串行。
2. 进阶版: 多 worker + merge。

三、真实用途
- 并行爬取、分片计算再汇总。
资料: https://go.dev/blog/pipelines ；https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197
*/
func serialSum(nums []int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func fanOutSum(nums []int, workers int) int {
	in := make(chan int)
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()
	var wg sync.WaitGroup
	out := make(chan int, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			local := 0
			for n := range in {
				local += n
			}
			out <- local
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	total := 0
	for v := range out {
		total += v
	}
	return total
}

func runTests() {
	nums := []int{1, 2, 3, 4, 5}
	if serialSum(nums) != 15 {
		panic("serial")
	}
	if fanOutSum(nums, 3) != 15 {
		panic("fanOut")
	}
	fmt.Println("✅ FanInFanOut 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 FanInFanOut ************\n")
	fmt.Println(fanOutSum([]int{1, 2, 3}, 2))
	runTests()
}

func main() {
	bootstrap()
}
