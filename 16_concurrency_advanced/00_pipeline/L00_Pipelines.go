package main

import "fmt"

/*
[主题]: Pipeline：多阶段 channel 串联
[难度]: Medium
[前置]: 03_concurrency

一、主题说明
把处理拆成 stage，每个 stage 从 in 读、写 out，形成流水线。

二、写法对比
1. 直观版: 单函数循环处理。
2. 进阶版: gen → square → sum pipeline。

三、真实用途
- ETL、日志处理、流式计算。
资料: https://go.dev/blog/pipelines
*/
func processNaive(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n * n
	}
	return sum
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}

func runTests() {
	nums := []int{1, 2, 3}
	if processNaive(nums) != 14 {
		panic("naive")
	}
	if sum(sq(gen(1, 2, 3))) != 14 {
		panic("pipeline")
	}
	fmt.Println("✅ Pipelines 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Pipelines ************\n")
	fmt.Println(sum(sq(gen(2, 3))))
	runTests()
}

func main() {
	bootstrap()
}
