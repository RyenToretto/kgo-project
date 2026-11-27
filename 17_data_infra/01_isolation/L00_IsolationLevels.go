package main

import "fmt"

/*
[主题]: 事务隔离级别直觉
[难度]: Medium
[前置]: 07_persistence 事务课

一、主题说明
读未提交/读已提交/可重复读/串行化——权衡一致性与吞吐。
本课用简易调度模拟脏读是否可见。

二、写法对比
1. 直观版: 无隔离，直接读最新缓冲（含未提交）。
2. 进阶版: 仅读已提交快照。

三、真实用途
- 账户余额、库存；选错隔离级别会导致资损或性能崩。
资料: https://pkg.go.dev/database/sql ；https://gorm.io/docs/
*/
type store struct {
	committed map[string]int
	dirty     map[string]int
}

func (s *store) readUncommitted(k string) int {
	if v, ok := s.dirty[k]; ok {
		return v
	}
	return s.committed[k]
}

func (s *store) readCommitted(k string) int {
	return s.committed[k]
}

func runTests() {
	s := &store{committed: map[string]int{"bal": 100}, dirty: map[string]int{}}
	s.dirty["bal"] = 50 // 未提交改动
	if s.readUncommitted("bal") != 50 {
		panic("脏读应看到 50")
	}
	if s.readCommitted("bal") != 100 {
		panic("读已提交应仍为 100")
	}
	s.committed["bal"] = 50
	delete(s.dirty, "bal")
	if s.readCommitted("bal") != 50 {
		panic("提交后应为 50")
	}
	fmt.Println("✅ IsolationLevels 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 IsolationLevels ************\n")
	runTests()
}

func main() {
	bootstrap()
}
