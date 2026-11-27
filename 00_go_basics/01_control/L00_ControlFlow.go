package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 控制流：if / switch / for / range
[难度]: Easy
[前置]: 00_go_basics/00_hello

一、主题说明
Go 没有 while；用 for 覆盖循环。if 可带短语句初始化。
switch 默认不贯穿（无需 break），可用 fallthrough 显式贯穿。

二、写法对比
1. 直观版: 多层 if-else 判断分数等级。
2. 进阶版: switch 表达式 + for range 处理切片。

三、真实用途
- API 参数校验、状态机分支、遍历请求列表。
资料: https://go.dev/tour/ ；https://gobyexample.com/
*/
func gradeNaive(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 60 {
		return "C"
	}
	return "D"
}

func gradeBetter(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 60:
		return "C"
	default:
		return "D"
	}
}

func sumPositive(nums []int) int {
	total := 0
	for _, n := range nums {
		if n <= 0 {
			continue
		}
		total += n
	}
	return total
}

func joinEvenIndexes(words []string) string {
	var parts []string
	for i, w := range words {
		if i%2 == 0 {
			parts = append(parts, w)
		}
	}
	return strings.Join(parts, ",")
}

func runTests() {
	if gradeNaive(95) != "A" || gradeBetter(95) != "A" {
		panic("95 应为 A")
	}
	if gradeNaive(75) != "C" || gradeBetter(75) != "C" {
		panic("75 应为 C")
	}
	if sumPositive([]int{1, -2, 3, 0, 4}) != 8 {
		panic("正数之和应为 8")
	}
	if joinEvenIndexes([]string{"a", "b", "c", "d"}) != "a,c" {
		panic("偶数下标应为 a,c")
	}
	fmt.Println("✅ ControlFlow 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ControlFlow ************\n")
	fmt.Println("grade(88)=", gradeBetter(88))
	fmt.Println("sumPositive=", sumPositive([]int{1, -1, 2}))
	runTests()
}

func main() {
	bootstrap()
}
