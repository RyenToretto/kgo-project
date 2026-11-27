package main

import (
	"fmt"
	"strings"
)

/*
[主题]: go test 心智（用 runTests 模拟表驱动）
[难度]: Easy
[前置]: 05_modules_tooling/00_modules

一、主题说明
真实项目用 *_test.go + testing.T；本仓库学习课用 runTests 便于 go run 验证。
表驱动：切片用例循环断言——与 go test 习惯一致。

二、写法对比
1. 直观版: 多个独立 if 断言。
2. 进阶版: 表驱动用例。

三、真实用途
- 单元测试；`go test ./...` 是 CI 标配。
资料: https://go.dev/doc/tutorial/add-a-test ；https://pkg.go.dev/testing
*/
func add(a, b int) int { return a + b }

func runNaiveChecks() bool {
	if add(1, 1) != 2 {
		return false
	}
	if add(2, 3) != 5 {
		return false
	}
	return true
}

func runTableDriven() error {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{"1+1", 1, 1, 2},
		{"2+3", 2, 3, 5},
		{"负", -1, 1, 0},
	}
	for _, tc := range cases {
		if got := add(tc.a, tc.b); got != tc.want {
			return fmt.Errorf("%s: got %d want %d", tc.name, got, tc.want)
		}
	}
	return nil
}

func runTests() {
	if !runNaiveChecks() {
		panic("naive")
	}
	if err := runTableDriven(); err != nil {
		panic(err.Error())
	}
	hint := "真实文件命名: foo_test.go ，命令: go test ./..."
	if !strings.Contains(hint, "go test") {
		panic("hint")
	}
	fmt.Println("✅ GoTestIntro 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GoTestIntro ************\n")
	fmt.Println("表驱动 err=", runTableDriven())
	runTests()
}

func main() {
	bootstrap()
}
