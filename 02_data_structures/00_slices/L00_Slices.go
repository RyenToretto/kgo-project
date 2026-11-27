package main

import "fmt"

/*
[主题]: 切片 slice：底层数组、append、共享陷阱
[难度]: Easy
[前置]: 00_go_basics

一、主题说明
slice 是对底层数组的视图（ptr/len/cap）。append 可能扩容换底层数组。
子切片共享数组：改一个可能影响另一个——高频面试与线上坑。

二、写法对比
1. 直观版: 直接 append 原切片。
2. 进阶版: copy 隔离 / 注意共享。

三、真实用途
- API 批量处理、缓冲区、几乎所有集合操作。
资料: https://gobyexample.com/slices ；https://go.dev/blog/slices-intro
*/
func appendNaive(in []int, v int) []int {
	return append(in, v)
}

func cloneThenAppend(in []int, v int) []int {
	out := make([]int, len(in), len(in)+1)
	copy(out, in)
	return append(out, v)
}

func sharedMutationDemo() (a0, b0 int) {
	base := []int{1, 2, 3, 4}
	a := base[:2]
	b := base[1:3]
	a[1] = 99 // 影响 base[1]，从而影响 b[0]
	return a[1], b[0]
}

func runTests() {
	s := []int{1, 2}
	s2 := appendNaive(s, 3)
	if len(s2) != 3 || s2[2] != 3 {
		panic("appendNaive 失败")
	}
	orig := []int{1, 2}
	iso := cloneThenAppend(orig, 9)
	iso[0] = 100
	if orig[0] != 1 {
		panic("clone 后修改不应影响原切片")
	}
	x, y := sharedMutationDemo()
	if x != 99 || y != 99 {
		panic("共享底层数组时两边都应看到 99")
	}
	fmt.Println("✅ Slices 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Slices ************\n")
	a, b := sharedMutationDemo()
	fmt.Println("shared:", a, b)
	runTests()
}

func main() {
	bootstrap()
}
