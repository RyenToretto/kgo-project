package main

import (
	"fmt"
	"strings"
)

/*
[主题]: SQL 索引与 EXPLAIN 直觉
[难度]: Medium
[前置]: 07_persistence（可并行理解）

一、主题说明
索引加速查找但拖慢写入。EXPLAIN 看是否走索引。本课用内存「表+索引」模拟。

二、写法对比
1. 直观版: 全表扫描。
2. 进阶版: 哈希索引 O(1) 查找。

三、真实用途
- 慢查询优化；线上事故高发区。
资料: https://pkg.go.dev/database/sql ；https://gorm.io/docs/
*/
type row struct {
	ID   int
	Name string
}

func fullScan(rows []row, id int) *row {
	for i := range rows {
		if rows[i].ID == id {
			return &rows[i]
		}
	}
	return nil
}

func indexLookup(idx map[int]int, rows []row, id int) *row {
	i, ok := idx[id]
	if !ok {
		return nil
	}
	return &rows[i]
}

func explainPlan(usedIndex bool) string {
	if usedIndex {
		return "type=ref key=PRIMARY"
	}
	return "type=ALL key=NULL"
}

func runTests() {
	rows := []row{{1, "a"}, {2, "b"}, {3, "c"}}
	idx := map[int]int{1: 0, 2: 1, 3: 2}
	if fullScan(rows, 2).Name != "b" {
		panic("scan")
	}
	if indexLookup(idx, rows, 3).Name != "c" {
		panic("index")
	}
	if !strings.Contains(explainPlan(true), "PRIMARY") {
		panic("explain")
	}
	fmt.Println("✅ SqlIndexExplain 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 SqlIndexExplain ************\n")
	runTests()
}

func main() {
	bootstrap()
}
