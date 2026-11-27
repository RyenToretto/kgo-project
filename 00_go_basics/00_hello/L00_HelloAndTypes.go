package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 第一个 Go 程序：main、基本类型与自测结构
[难度]: Easy
[前置]: 无（Go 小白第一课）

一、主题说明
Go 程序从 package main 的 func main 开始执行。
本课建立「写逻辑 → bootstrap 演示 → runTests 断言」的学习习惯，与仓库其它课一致。

二、写法对比
1. 直观版 addNaive: 直接用 + 号。
2. 进阶版 addBetter: 同样简单加法，强调命名与可测函数拆分。

三、真实用途
- 所有 CLI、HTTP 服务、worker 都有入口。
- 数值与类型选择在金额、计数、ID 场景非常关键（金额常用 int64 分或 decimal 库）。
*/
func addNaive(left, right int) int {
	return left + right
}

func addBetter(left, right int) int {
	return left + right
}

func describePrimitives() string {
	age := 18
	price := 19.9
	ok := true
	grade := 'A' // rune
	return fmt.Sprintf("age=%d, price=%v, ok=%v, grade=%c", age, price, ok, grade)
}

func runTests() {
	if addNaive(1, 2) != 3 {
		panic("addNaive(1,2) 应为 3")
	}
	if addBetter(2, 3) != 5 {
		panic("addBetter(2,3) 应为 5")
	}
	if !strings.Contains(describePrimitives(), "age=18") {
		panic("describePrimitives 应包含 age=18")
	}
	fmt.Println("✅ HelloAndTypes 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 HelloAndTypes ************\n")
	fmt.Println("基本类型示例:", describePrimitives())
	fmt.Println("1 + 2 =", addNaive(1, 2))
	runTests()
}

func main() {
	bootstrap()
}
