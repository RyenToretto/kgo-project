package main

import (
	"fmt"
	"unicode/utf8"
)

/*
[主题]: string / byte / rune
[难度]: Easy
[前置]: 02_data_structures/01_maps

一、主题说明
string 是不可变字节序列；中文等多字节字符用 rune（Unicode code point）。
len(s) 是字节数；用 utf8.RuneCountInString 或 range 按字符计数。

二、写法对比
1. 直观版: 用 len 当「字符数」（对中文错误）。
2. 进阶版: rune / utf8 正确计数与截取直觉。

三、真实用途
- 用户昵称长度校验、日志截断、国际化文本。
资料: https://gobyexample.com/strings-and-runes ；https://go.dev/blog/strings
*/
func byteLenNaive(s string) int {
	return len(s)
}

func runeCountBetter(s string) int {
	return utf8.RuneCountInString(s)
}

func firstRune(s string) (rune, bool) {
	for _, r := range s {
		return r, true
	}
	return 0, false
}

func runTests() {
	s := "Go你好"
	if byteLenNaive(s) == runeCountBetter(s) {
		panic("中英混合时字节数应不等于字符数")
	}
	if runeCountBetter(s) != 4 {
		panic("Go你好 应为 4 个 rune")
	}
	r, ok := firstRune("你好")
	if !ok || r != '你' {
		panic("首 rune 应为 你")
	}
	fmt.Println("✅ StringsAndRunes 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 StringsAndRunes ************\n")
	s := "Go你好"
	fmt.Printf("bytes=%d runes=%d\n", byteLenNaive(s), runeCountBetter(s))
	runTests()
}

func main() {
	bootstrap()
}
