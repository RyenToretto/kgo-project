package main

import (
	"fmt"
	"go/format"
	"strings"
)

/*
[主题]: gofmt / go vet 习惯
[难度]: Easy
[前置]: 05_modules_tooling/00_modules

一、主题说明
gofmt（或 gofmt -s / goimports）统一格式；go vet 抓常见错误。
本课用 go/format 演示「格式化源码」心智，离线可跑。

二、写法对比
1. 直观版: 手工空格不统一。
2. 进阶版: format.Source。

三、真实用途
- 提交前格式化；CI 检查 gofmt / vet / staticcheck。
资料: https://go.dev/doc/code ；https://pkg.go.dev/go/format
*/
func messySource() string {
	return "package p\nfunc F(  x int  ) int {return x+1}\n"
}

func formatSourceBetter(src string) (string, error) {
	out, err := format.Source([]byte(src))
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func vetStyleHint(code string) []string {
	var hints []string
	if strings.Contains(code, "fmt.Sprintf") && strings.Contains(code, "Printf") {
		hints = append(hints, "检查 Printf 格式化动词是否匹配")
	}
	if strings.Contains(code, "Unlock()") && !strings.Contains(code, "defer") {
		hints = append(hints, "锁是否应 defer Unlock")
	}
	return hints
}

func runTests() {
	formatted, err := formatSourceBetter(messySource())
	if err != nil {
		panic(err.Error())
	}
	if !strings.Contains(formatted, "func F(x int)") {
		panic("gofmt 应规范参数空格: " + formatted)
	}
	hints := vetStyleHint("mu.Lock(); work(); mu.Unlock()")
	if len(hints) == 0 {
		panic("应提示 defer Unlock")
	}
	fmt.Println("✅ FmtVetStatic 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 FmtVetStatic ************\n")
	out, _ := formatSourceBetter(messySource())
	fmt.Println(out)
	runTests()
}

func main() {
	bootstrap()
}
