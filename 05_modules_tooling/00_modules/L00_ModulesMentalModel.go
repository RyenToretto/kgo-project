package main

import (
	"fmt"
	"strings"
)

/*
[主题]: Go Modules 心智模型
[难度]: Easy
[前置]: 00_go_basics/04_packages

一、主题说明
go.mod 声明 module path 与 Go 版本；依赖有版本选择算法。
本课用纯逻辑模拟「模块路径 / 依赖版本」概念，不改仓库根 go.mod。

二、写法对比
1. 直观版: 硬编码 import 路径字符串。
2. 进阶版: 解析 module 路径片段与版本比较直觉。

三、真实用途
- 每个真实项目都有 go.mod；协作与 CI 依赖可复现构建。
资料: https://go.dev/ref/mod ；https://go.dev/doc/modules/managing-dependencies ；https://go.dev/doc/code
*/
func moduleImportPathNaive(module, pkg string) string {
	return module + "/" + pkg
}

func joinModulePath(module string, elems ...string) string {
	parts := []string{strings.Trim(module, "/")}
	for _, e := range elems {
		e = strings.Trim(e, "/")
		if e != "" {
			parts = append(parts, e)
		}
	}
	return strings.Join(parts, "/")
}

// preferNewer 极简版本比较：只处理 vN.M.P 数字段，够用讲心智。
func preferNewer(a, b string) string {
	pa := parseSemver(a)
	pb := parseSemver(b)
	for i := 0; i < 3; i++ {
		if pa[i] != pb[i] {
			if pa[i] > pb[i] {
				return a
			}
			return b
		}
	}
	return a
}

func parseSemver(v string) [3]int {
	v = strings.TrimPrefix(v, "v")
	var out [3]int
	fmt.Sscanf(v, "%d.%d.%d", &out[0], &out[1], &out[2])
	return out
}

func runTests() {
	if moduleImportPathNaive("example.com/app", "internal/user") != "example.com/app/internal/user" {
		panic("naive path")
	}
	if joinModulePath("example.com/app", "internal", "user") != "example.com/app/internal/user" {
		panic("join path")
	}
	if preferNewer("v1.2.0", "v1.10.0") != "v1.10.0" {
		panic("1.10 > 1.2")
	}
	fmt.Println("✅ ModulesMentalModel 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ModulesMentalModel ************\n")
	fmt.Println(joinModulePath("github.com/you/kgo", "pkg", "auth"))
	fmt.Println("newer=", preferNewer("v0.9.0", "v1.0.0"))
	runTests()
}

func main() {
	bootstrap()
}
