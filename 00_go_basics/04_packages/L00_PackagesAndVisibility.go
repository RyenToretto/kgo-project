package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 包与可见性（导出规则）心智
[难度]: Easy
[前置]: 00_go_basics/03_structs

一、主题说明
同一目录同一 package。标识符首字母大写 = 导出（其他包可访问）；小写 = 包内私有。
真实多包项目用 module path import；本课用「模拟导出 API」演示规则（单文件可 go run）。

二、写法对比
1. 直观版: 全部小写函数挤在一起。
2. 进阶版: 公开 Facade + 私有 helper（命名约定）。

三、真实用途
- 库设计：只导出稳定 API；内部实现保持小写。
资料: https://go.dev/doc/code ；https://go.dev/ref/spec#Exported_identifiers
*/

// PublicGreeter 模拟「导出」类型：给其他包用的稳定 API。
type PublicGreeter struct {
	Prefix string
}

func (g PublicGreeter) Hello(name string) string {
	return formatGreeting(g.Prefix, name)
}

// formatGreeting 模拟包内私有 helper（小写开头）。
func formatGreeting(prefix, name string) string {
	p := strings.TrimSpace(prefix)
	n := strings.TrimSpace(name)
	if p == "" {
		p = "Hi"
	}
	if n == "" {
		n = "gopher"
	}
	return p + ", " + n + "!"
}

func runTests() {
	g := PublicGreeter{Prefix: "Hello"}
	if g.Hello("Go") != "Hello, Go!" {
		panic("Hello 输出错误")
	}
	if formatGreeting("  Yo ", "  Ada ") != "Yo, Ada!" {
		panic("私有 helper 应 trim")
	}
	if formatGreeting("", "") != "Hi, gopher!" {
		panic("默认值错误")
	}
	fmt.Println("✅ PackagesAndVisibility 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 PackagesAndVisibility ************\n")
	fmt.Println(PublicGreeter{Prefix: "你好"}.Hello("同学"))
	runTests()
}

func main() {
	bootstrap()
}
