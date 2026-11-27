package main

import "fmt"

/*
[主题]: map：读写、零值、逗号 ok
[难度]: Easy
[前置]: 02_data_structures/00_slices

一、主题说明
map 是引用类型；未 make 的 nil map 可读（得零值）但写会 panic。
用 `v, ok := m[k]` 区分「没有键」与「零值」。

二、写法对比
1. 直观版: 直接下标取值。
2. 进阶版: ok 判断 + 安全写入。

三、真实用途
- 缓存、索引、配置字典、去重。
资料: https://gobyexample.com/maps ；https://go.dev/blog/maps
*/
func getNaive(m map[string]int, key string) int {
	return m[key]
}

func getBetter(m map[string]int, key string) (int, bool) {
	v, ok := m[key]
	return v, ok
}

func ensureMap(m map[string]int) map[string]int {
	if m == nil {
		m = make(map[string]int)
	}
	return m
}

func runTests() {
	m := map[string]int{"a": 1}
	if getNaive(m, "a") != 1 {
		panic("应读到 1")
	}
	if getNaive(m, "missing") != 0 {
		panic("缺失键零值应为 0")
	}
	_, ok := getBetter(m, "missing")
	if ok {
		panic("缺失键 ok 应为 false")
	}
	var nilMap map[string]int
	nilMap = ensureMap(nilMap)
	nilMap["x"] = 7
	if nilMap["x"] != 7 {
		panic("ensureMap 后应可写")
	}
	fmt.Println("✅ Maps 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Maps ************\n")
	m := map[string]int{"go": 1}
	fmt.Println(getBetter(m, "go"))
	runTests()
}

func main() {
	bootstrap()
}
