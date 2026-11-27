package main

import "fmt"

/*
[主题]: 逃逸分析直觉
[难度]: Hard
[前置]: 15_runtime_advanced/01_gc

一、主题说明
变量可能分配在栈或堆。返回局部变量地址、接口装箱等常导致逃逸到堆 → 增加 GC 压力。
用 `go build -gcflags=-m` 查看；本课用行为对比讲心智。

二、写法对比
1. 直观版: 返回指针（可能逃逸）。
2. 进阶版: 尽量值传递小对象 / 复用缓冲。

三、真实用途
- 热路径优化；配合 pprof allocs。
资料: https://go.dev/doc/diagnostics ；https://www.amazon.com/Efficient-Go-Data-Driven-Performance-Optimization/dp/1098105710
*/
func returnPtrNaive(v int) *int {
	return &v // 典型逃逸候选
}

func returnValueBetter(v int) int {
	return v
}

func runTests() {
	p := returnPtrNaive(7)
	if *p != 7 {
		panic("ptr")
	}
	if returnValueBetter(7) != 7 {
		panic("value")
	}
	fmt.Println("✅ EscapeAnalysis 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 EscapeAnalysis ************\n")
	fmt.Println(*returnPtrNaive(1), returnValueBetter(2))
	runTests()
}

func main() {
	bootstrap()
}
