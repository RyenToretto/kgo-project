package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 函数：多返回值、命名返回、闭包、defer 入门
[难度]: Easy
[前置]: 00_go_basics/01_control

一、主题说明
Go 函数可返回多个值（常见 error）。defer 在函数返回前按 LIFO 执行，常用于关资源。

二、写法对比
1. 直观版: 分开写除法与错误判断。
2. 进阶版: 多返回值 + defer 记录调用痕迹 + 闭包累加器。

三、真实用途
- 几乎所有业务函数：`(result, error)`；HTTP/DB 用 defer Close。
资料: https://go.dev/doc/effective_go ；https://gobyexample.com/
*/
func divideNaive(a, b int) (int, string) {
	if b == 0 {
		return 0, "除数不能为 0"
	}
	return a / b, ""
}

func divideBetter(a, b int) (quot int, err error) {
	defer func() {
		// defer 可看到命名返回值；这里仅演示时机
		_ = quot
	}()
	if b == 0 {
		return 0, fmt.Errorf("除数不能为 0")
	}
	return a / b, nil
}

func makeCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func withCleanup(trace *[]string, name string, body func()) {
	*trace = append(*trace, "enter:"+name)
	defer func() {
		*trace = append(*trace, "exit:"+name)
	}()
	body()
}

func runTests() {
	q, msg := divideNaive(10, 2)
	if q != 5 || msg != "" {
		panic("10/2 应为 5")
	}
	_, msg = divideNaive(1, 0)
	if msg == "" {
		panic("除 0 应有错误信息")
	}
	q2, err := divideBetter(9, 3)
	if err != nil || q2 != 3 {
		panic("9/3 应为 3")
	}
	_, err = divideBetter(1, 0)
	if err == nil {
		panic("除 0 应返回 error")
	}
	c := makeCounter()
	if c() != 1 || c() != 2 {
		panic("闭包计数器错误")
	}
	var trace []string
	withCleanup(&trace, "job", func() {
		trace = append(trace, "work")
	})
	got := strings.Join(trace, ",")
	if got != "enter:job,work,exit:job" {
		panic("defer 顺序错误: " + got)
	}
	fmt.Println("✅ Functions 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Functions ************\n")
	q, err := divideBetter(8, 2)
	fmt.Println("8/2=", q, "err=", err)
	runTests()
}

func main() {
	bootstrap()
}
