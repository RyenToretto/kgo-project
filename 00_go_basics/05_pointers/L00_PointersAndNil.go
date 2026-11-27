package main

import "fmt"

/*
[主题]: 指针与 nil
[难度]: Easy
[前置]: 00_go_basics/03_structs

一、主题说明
& 取址，* 解引用。Go 没有指针运算。nil 是指针/切片/map/channel/函数/接口的零值之一。
对 nil 指针解引用会 panic——生产代码必须判空或保证非 nil。

二、写法对比
1. 直观版: 传值修改失败（拷贝）。
2. 进阶版: 传指针修改成功 + 安全解引用。

三、真实用途
- 大结构体避免拷贝；可选字段用指针区分「未设置」与零值。
资料: https://go.dev/tour/moretypes/1 ；https://gobyexample.com/pointers
*/
func setAgeByValue(age int) {
	age = 99
}

func setAgeByPointer(age *int) {
	if age == nil {
		return
	}
	*age = 99
}

func safeDeref(p *string) string {
	if p == nil {
		return "<nil>"
	}
	return *p
}

func runTests() {
	age := 18
	setAgeByValue(age)
	if age != 18 {
		panic("传值不应修改原变量")
	}
	setAgeByPointer(&age)
	if age != 99 {
		panic("传指针应修改为 99")
	}
	setAgeByPointer(nil) // 不应 panic
	if safeDeref(nil) != "<nil>" {
		panic("nil 应返回 <nil>")
	}
	s := "ok"
	if safeDeref(&s) != "ok" {
		panic("解引用错误")
	}
	fmt.Println("✅ PointersAndNil 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 PointersAndNil ************\n")
	x := 1
	fmt.Println("before", x)
	setAgeByPointer(&x)
	fmt.Println("after", x)
	runTests()
}

func main() {
	bootstrap()
}
