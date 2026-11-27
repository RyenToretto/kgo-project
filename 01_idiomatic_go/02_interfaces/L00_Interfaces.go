package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 接口：隐式实现与小接口
[难度]: Easy
[前置]: 01_idiomatic_go/00_errors

一、主题说明
Go 接口由方法集定义；类型只要实现方法即满足接口（无需 implements 关键字）。
偏好小接口（如 io.Reader）。

二、写法对比
1. 直观版: 针对具体类型写函数。
2. 进阶版: 依赖接口，便于替换实现与测试。

三、真实用途
- 存储/支付/通知等多实现；单元测试用假对象。
资料: https://go.dev/doc/effective_go ；https://go.dev/tour/methods
*/
type Greeter interface {
	Greet(name string) string
}

type PlainGreeter struct{}

func (PlainGreeter) Greet(name string) string {
	return "Hello, " + name
}

type LoudGreeter struct{}

func (LoudGreeter) Greet(name string) string {
	return strings.ToUpper("hello, " + name)
}

func greetConcrete(g PlainGreeter, name string) string {
	return g.Greet(name)
}

func greetAny(g Greeter, name string) string {
	return g.Greet(name)
}

func runTests() {
	if greetConcrete(PlainGreeter{}, "Ada") != "Hello, Ada" {
		panic("concrete 错误")
	}
	if greetAny(PlainGreeter{}, "Ada") != "Hello, Ada" {
		panic("Plain 接口调用错误")
	}
	if greetAny(LoudGreeter{}, "Ada") != "HELLO, ADA" {
		panic("Loud 接口调用错误")
	}
	fmt.Println("✅ Interfaces 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Interfaces ************\n")
	fmt.Println(greetAny(LoudGreeter{}, "Go"))
	runTests()
}

func main() {
	bootstrap()
}
