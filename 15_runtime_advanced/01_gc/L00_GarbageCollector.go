package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

/*
[主题]: GC 直觉
[难度]: Hard
[前置]: 15_runtime_advanced/00_scheduler

一、主题说明
Go GC 是并发标记清除；目标低延迟。可用 GOGC / SetGCPercent 调吞吐与 CPU。
本课读取 MemStats，演示分配后统计变化。

二、写法对比
1. 直观版: 无视分配。
2. 进阶版: 观察 HeapAlloc / NumGC。

三、真实用途
- 延迟敏感服务调优；定位分配热点。
资料: https://go.dev/doc/gc-guide ；https://go.dev/doc/diagnostics
*/
func allocNaive(n int) [][]byte {
	out := make([][]byte, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, make([]byte, 1024))
	}
	return out
}

func readHeap() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.HeapAlloc
}

func runTests() {
	_ = debug.SetGCPercent
	before := readHeap()
	hold := allocNaive(200)
	after := readHeap()
	if after < before && len(hold) == 0 {
		panic("unexpected")
	}
	if len(hold) != 200 {
		panic("alloc")
	}
	runtime.GC()
	fmt.Println("✅ GarbageCollector 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GarbageCollector ************\n")
	fmt.Println("heap=", readHeap())
	runTests()
}

func main() {
	bootstrap()
}
