package main

import (
	"fmt"
	"strings"
)

/*
[主题]: DDD 战术：实体 / 值对象 / 聚合根心智
[难度]: Medium
[前置]: 13_project_capstone

一、主题说明
聚合根保证不变量；值对象不可变；应用服务编排用例。
本课用订单聚合：总额=明细之和。

二、写法对比
1. 直观版: 贫血 map 乱改字段。
2. 进阶版: 聚合根方法维护不变量。

三、真实用途
- 复杂业务域建模。
资料: https://github.com/joelparkerhenderson/architecture-decision-record ；https://sre.google/books/
*/
type line struct {
	SKU   string
	Price int
	Qty   int
}

type orderAnemic struct {
	Lines []line
	Total int
}

type Order struct {
	lines []line
}

func (o *Order) Add(sku string, price, qty int) error {
	if price < 0 || qty <= 0 || strings.TrimSpace(sku) == "" {
		return fmt.Errorf("invalid line")
	}
	o.lines = append(o.lines, line{SKU: sku, Price: price, Qty: qty})
	return nil
}

func (o Order) Total() int {
	s := 0
	for _, l := range o.lines {
		s += l.Price * l.Qty
	}
	return s
}

func runTests() {
	a := orderAnemic{Lines: []line{{"x", 10, 2}}, Total: 999} // 不一致
	if a.Total != 999 {
		panic("anemic allows inconsistency")
	}
	var o Order
	if err := o.Add("sku", 10, 2); err != nil || o.Total() != 20 {
		panic("aggregate")
	}
	if err := o.Add("", 1, 1); err == nil {
		panic("should reject")
	}
	fmt.Println("✅ DddTactical 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 DddTactical ************\n")
	runTests()
}

func main() {
	bootstrap()
}
