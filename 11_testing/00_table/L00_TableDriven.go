package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 表驱动测试
[难度]: Easy
[前置]: 05_modules_tooling/01_gotest

一、主题说明
把输入/期望放进切片，循环断言；新增用例只加一行。
真实项目写在 *_test.go 里用 testing.T；本课 runTests 便于 go run。

二、写法对比
1. 直观版: 每个场景复制一段 if panic。
2. 进阶版: 表驱动 + 子用例名。

三、真实用途
- 边界值、错误路径；与 go test 子测试 t.Run 同构。
资料: https://pkg.go.dev/testing ；https://go.dev/doc/tutorial/add-a-test
*/

func slugifyNaive(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}

func slugifyBetter(s string) (string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return "", fmt.Errorf("empty")
	}
	s = strings.ToLower(s)
	var b strings.Builder
	prevDash := false
	for _, r := range s {
		if r == ' ' || r == '-' {
			if !prevDash && b.Len() > 0 {
				b.WriteByte('-')
				prevDash = true
			}
			continue
		}
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			prevDash = false
		}
	}
	out := strings.Trim(b.String(), "-")
	if out == "" {
		return "", fmt.Errorf("no alnum")
	}
	return out, nil
}

func runTests() {
	if slugifyNaive("Hello World") != "hello-world" {
		panic("naive")
	}

	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"正常", "Go Lang", "go-lang", false},
		{"多空格", "a   b", "a-b", false},
		{"空串", "", "", true},
		{"纯符号", "!!!", "", true},
	}
	for _, tc := range cases {
		got, err := slugifyBetter(tc.in)
		if tc.wantErr {
			if err == nil {
				panic(tc.name + " 应失败")
			}
			continue
		}
		if err != nil || got != tc.want {
			panic(fmt.Sprintf("%s: got %q err %v", tc.name, got, err))
		}
	}
	fmt.Println("✅ TableDriven 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 TableDriven ************\n")
	fmt.Println("slugify:", slugifyNaive("KGO Project"))
	runTests()
}

func main() {
	bootstrap()
}