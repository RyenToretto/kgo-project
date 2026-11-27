package main

import (
	"fmt"
	"strings"
	"time"
)

/*
[主题]: benchmark 心智（手工计时模拟）
[难度]: Easy
[前置]: 11_testing/00_table

一、主题说明
go test -bench 用 testing.B 测纳秒/op；学习课用手工循环+计时理解「测什么」。
关注：分配次数、避免在循环里重复 setup。

二、写法对比
1. 直观版: 只测一次，易被 JIT/缓存误导。
2. 进阶版: 预热 + 多次取中位数（简化版）。

三、真实用途
- 热点路径优化前要有基准；配合 pprof。
资料: https://pkg.go.dev/testing#hdr-Benchmarks
*/

func concatNaive(parts []string) string {
	var s string
	for _, p := range parts {
		s += p
	}
	return s
}

func concatBetter(parts []string) string {
	return strings.Join(parts, "")
}

func measureOps(op func(), warmup, rounds int) time.Duration {
	for i := 0; i < warmup; i++ {
		op()
	}
	start := time.Now()
	for i := 0; i < rounds; i++ {
		op()
	}
	return time.Since(start) / time.Duration(rounds)
}

func runTests() {
	parts := make([]string, 100)
	for i := range parts {
		parts[i] = "x"
	}
	a := concatNaive(parts)
	b := concatBetter(parts)
	if len(a) != len(b) || a != b {
		panic("结果应一致")
	}
	dNaive := measureOps(func() { _ = concatNaive(parts) }, 10, 200)
	dBetter := measureOps(func() { _ = concatBetter(parts) }, 10, 200)
	if dBetter > dNaive*2 {
		// 小切片可能差异不大；大切片 Join 通常更稳
		fmt.Println("提示: 大字符串 + 应首选 strings.Builder/Join")
	}
	fmt.Println("✅ BenchmarkIdea 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 BenchmarkIdea ************\n")
	fmt.Println("真实命令: go test -bench=. -benchmem ./...")
	runTests()
}

func main() {
	bootstrap()
}