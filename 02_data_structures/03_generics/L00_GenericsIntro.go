package main

import "fmt"

/*
[主题]: 泛型入门：类型参数与约束
[难度]: Medium
[前置]: 02_data_structures/00_slices

一、主题说明
Go 1.18+ 支持泛型。用类型参数减少样板；约束限制可用操作。
本课纯标准库，自定义 Ordered 约束，无第三方依赖。

二、写法对比
1. 直观版: 分别为 int 写 Max。
2. 进阶版: 泛型 Max[T Ordered] + MapSlice。

三、真实用途
- 通用容器、工具函数、减少复制粘贴。
资料: https://go.dev/doc/tutorial/generics ；https://go.dev/blog/go1.18
*/
type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

func maxIntNaive(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MapSlice[T any, R any](in []T, f func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func runTests() {
	if maxIntNaive(3, 5) != 5 {
		panic("naive max")
	}
	if Max(3, 5) != 5 || Max("a", "b") != "b" {
		panic("generic Max 错误")
	}
	got := MapSlice([]int{1, 2}, func(x int) int { return x * 2 })
	if len(got) != 2 || got[0] != 2 || got[1] != 4 {
		panic("MapSlice 错误")
	}
	fmt.Println("✅ GenericsIntro 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GenericsIntro ************\n")
	fmt.Println(Max(10, 20), Max("go", "zz"))
	runTests()
}

func main() {
	bootstrap()
}
