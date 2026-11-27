package main

import (
	"bytes"
	"fmt"
	"runtime/pprof"
	"sort"
)

/*
[主题]: pprof 性能剖析心智
[难度]: Medium
[前置]: 11_testing/01_bench

一、主题说明
线上：import _ "net/http/pprof" 暴露 /debug/pprof；用 go tool pprof 看 CPU/堆。
本课用 runtime/pprof 写剖析缓冲，理解「采样调用栈」即可，不启真实 pprof HTTP。

二、写法对比
1. 直观版: 猜哪慢，无数据优化。
2. 进阶版: 先 profile 再改热点。

三、真实用途
- 协程泄漏、内存涨、CPU 尖峰。
资料: https://go.dev/blog/pprof ；https://go.dev/doc/diagnostics
*/

func hotNaive(n int) int {
	// 故意低效：重复分配
	sum := 0
	for i := 0; i < n; i++ {
		s := make([]byte, 1024)
		sum += len(s)
	}
	return sum
}

func hotBetter(n int) int {
	buf := make([]byte, 1024)
	sum := 0
	for i := 0; i < n; i++ {
		sum += len(buf)
	}
	return sum
}

func writeCPUProfile(buf *bytes.Buffer, work func()) error {
	if err := pprof.StartCPUProfile(buf); err != nil {
		return err
	}
	work()
	pprof.StopCPUProfile()
	return nil
}

func runTests() {
	if hotNaive(10) != hotBetter(10) {
		panic("same sum")
	}
	var prof bytes.Buffer
	if err := writeCPUProfile(&prof, func() { _ = hotBetter(500) }); err != nil {
		panic(err)
	}
	if prof.Len() == 0 {
		panic("profile empty")
	}
	// 堆 profile 心智
	hp := pprof.Lookup("heap")
	if hp == nil {
		panic("no heap")
	}
	var heapBuf bytes.Buffer
	if err := hp.WriteTo(&heapBuf, 0); err != nil {
		panic(err)
	}

	// 假「Top 栈」：记录两种热点函数名，模拟 pprof 排序输出
	names := []string{"hotNaive", "hotBetter"}
	sort.Strings(names)
	if names[0] != "hotBetter" || names[1] != "hotNaive" {
		panic("sort demo")
	}

	fmt.Println("✅ PprofMental 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 PprofMental ************\n")
	fmt.Println("习惯: CPU 用 StartCPUProfile；阻塞/互斥有专门 profile；结合 benchmark 与 -memprofile。")
	runTests()
}

func main() {
	bootstrap()
}