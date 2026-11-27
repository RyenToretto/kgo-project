package main

import (
	"fmt"
	"strings"
)

/*
[主题]: client-go 心智（不引入真实依赖）
[难度]: Medium
[前置]: 18_cloud_native/01_probes

一、主题说明
client-go 用 rest.Config 连接 API Server，对 Pod/Deployment 做 CRUD/Watch。
本课模拟「List Pods → 过滤 Ready」心智，避免拉取庞大依赖。

二、写法对比
1. 直观版: kubectl 手工。
2. 进阶版: 程序化 List/Filter。

三、真实用途
- Operator、控制器、内部运维工具。
资料: https://github.com/kubernetes/client-go ；https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
*/
type pod struct {
	Name  string
	Ready bool
}

func kubectlNaive() string { return "kubectl get pods" }

func listReady(pods []pod) []string {
	var names []string
	for _, p := range pods {
		if p.Ready {
			names = append(names, p.Name)
		}
	}
	return names
}

func runTests() {
	if !strings.Contains(kubectlNaive(), "kubectl") {
		panic("naive")
	}
	got := listReady([]pod{{"a", true}, {"b", false}, {"c", true}})
	if len(got) != 2 || got[0] != "a" || got[1] != "c" {
		panic("listReady")
	}
	fmt.Println("✅ ClientGoMental 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ClientGoMental ************\n")
	runTests()
}

func main() {
	bootstrap()
}
