package main

import "fmt"

/*
[主题]: 组合与嵌入（embedding），而非类继承
[难度]: Easy
[前置]: 01_idiomatic_go/02_interfaces

一、主题说明
Go 没有经典继承。通过嵌入结构体提升字段/方法；通过接口组合行为。

二、写法对比
1. 直观版: 手动转发方法。
2. 进阶版: 嵌入自动提升方法。

三、真实用途
- HTTP middleware 包装、领域实体组合、sync.Mutex 嵌入。
资料: https://go.dev/doc/effective_go ；https://go.dev/wiki/CodeReviewComments
*/
type Engine struct {
	Power int
}

func (e Engine) Start() string {
	return fmt.Sprintf("engine:%d", e.Power)
}

type CarNaive struct {
	Eng Engine
}

func (c CarNaive) Start() string {
	return c.Eng.Start()
}

type CarBetter struct {
	Engine // 嵌入
	Brand  string
}

func runTests() {
	n := CarNaive{Eng: Engine{Power: 100}}
	if n.Start() != "engine:100" {
		panic("naive Start 错误")
	}
	b := CarBetter{Engine: Engine{Power: 200}, Brand: "X"}
	if b.Start() != "engine:200" {
		panic("嵌入提升 Start 错误")
	}
	if b.Power != 200 {
		panic("嵌入应提升字段 Power")
	}
	fmt.Println("✅ CompositionEmbedding 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 CompositionEmbedding ************\n")
	c := CarBetter{Engine: Engine{Power: 150}, Brand: "GoCar"}
	fmt.Println(c.Brand, c.Start())
	runTests()
}

func main() {
	bootstrap()
}
