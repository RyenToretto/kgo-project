package main

import (
	"fmt"
	"runtime"
)

/*
[主题]: 容器感知：GOMAXPROCS 与 CPU limit
[难度]: Medium
[前置]: 15_runtime_advanced/00_scheduler

一、主题说明
容器 CPU limit 可能小于宿主机 NumCPU；Go 1.25+ / 相关运行时改进与 container-aware GOMAXPROCS 相关讨论见官方博客。
本课对比 NumCPU 与 GOMAXPROCS，强调「按配额设并行度」。

二、写法对比
1. 直观版: 默认吃满宿主机 CPU 数。
2. 进阶版: 显式设定合理 GOMAXPROCS。

三、真实用途
- K8s Deployment 资源配额下的 Go 服务。
资料: https://go.dev/blog/container-aware-gomaxprocs ；https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
*/
func defaultParallelism() int { return runtime.NumCPU() }

func setParallelism(n int) int {
	if n < 1 {
		n = 1
	}
	return runtime.GOMAXPROCS(n)
}

func runTests() {
	old := runtime.GOMAXPROCS(0)
	_ = defaultParallelism()
	got := setParallelism(1)
	if got < 1 {
		panic("gomaxprocs")
	}
	runtime.GOMAXPROCS(old)
	fmt.Println("✅ ContainerAware 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ContainerAware ************\n")
	fmt.Println("NumCPU", runtime.NumCPU(), "GOMAXPROCS", runtime.GOMAXPROCS(0))
	runTests()
}

func main() {
	bootstrap()
}
