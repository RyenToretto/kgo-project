package main

import (
	"errors"
	"fmt"
)

/*
[主题]: 地道 error 处理：哨兵、包装、判断
[难度]: Easy
[前置]: 00_go_basics

一、主题说明
Go 用返回值传递错误，而不是异常主路径。errors.Is / errors.As 用于判断包装链。

二、写法对比
1. 直观版: 字符串比较错误信息（脆弱）。
2. 进阶版: 哨兵 error + fmt.Errorf("%w") 包装。

三、真实用途
- 所有 API/DB 调用；业务层把底层错误包装成可判断的领域错误。
资料: https://go.dev/doc/effective_go ；https://go.dev/blog/go1.13-errors ；https://github.com/uber-go/guide/blob/master/style.md
*/
var ErrNotFound = errors.New("not found")

func findNaive(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("not found")
	}
	return "ok", nil
}

func findBetter(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("user %d: %w", id, ErrNotFound)
	}
	return "ok", nil
}

func runTests() {
	_, err := findNaive(-1)
	if err == nil || err.Error() != "not found" {
		panic("naive 应返回 not found 字符串错误")
	}
	_, err = findBetter(-1)
	if !errors.Is(err, ErrNotFound) {
		panic("better 应用 errors.Is 命中 ErrNotFound")
	}
	_, err = findBetter(1)
	if err != nil {
		panic("id=1 不应出错")
	}
	fmt.Println("✅ ErrorHandling 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ErrorHandling ************\n")
	_, err := findBetter(-7)
	fmt.Println("errors.Is NotFound?", errors.Is(err, ErrNotFound), "msg=", err)
	runTests()
}

func main() {
	bootstrap()
}
