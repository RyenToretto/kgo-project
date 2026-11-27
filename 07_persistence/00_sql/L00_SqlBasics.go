package main

import "fmt"

/*
[主题]: SQL CRUD 心智（内存表模拟）
[难度]: Easy
[前置]: 02_data_structures/01_maps

一、主题说明
表 = 行集合；主键唯一；SELECT/WHERE/INSERT/UPDATE/DELETE 语义。
本课用 map[int]Row 模拟，无真实数据库。

二、写法对比
1. 直观版: 全表扫描字符串匹配。
2. 进阶版: 主键索引 O(1) 查。

三、真实用途
- 写 ORM/SQL 前先懂关系模型。
资料: https://pkg.go.dev/database/sql ；https://github.com/go-sql-driver/mysql
*/

type row struct {
	ID   int
	Name string
}

type memTable struct {
	rows map[int]row
	next int
}

func newMemTable() *memTable {
	return &memTable{rows: make(map[int]row), next: 1}
}

func (t *memTable) insert(name string) row {
	r := row{ID: t.next, Name: name}
	t.rows[r.ID] = r
	t.next++
	return r
}

func (t *memTable) selectByID(id int) (row, bool) {
	r, ok := t.rows[id]
	return r, ok
}

func (t *memTable) updateName(id int, name string) bool {
	r, ok := t.rows[id]
	if !ok {
		return false
	}
	r.Name = name
	t.rows[id] = r
	return true
}

func (t *memTable) deleteByID(id int) bool {
	if _, ok := t.rows[id]; !ok {
		return false
	}
	delete(t.rows, id)
	return true
}

// 直观版：用切片线性找
func findNaive(all []row, id int) (row, bool) {
	for _, r := range all {
		if r.ID == id {
			return r, true
		}
	}
	return row{}, false
}

func runTests() {
	all := []row{{ID: 1, Name: "a"}, {ID: 2, Name: "b"}}
	if _, ok := findNaive(all, 2); !ok {
		panic("findNaive")
	}
	t := newMemTable()
	r := t.insert("Go")
	if r.ID != 1 {
		panic("insert id")
	}
	got, ok := t.selectByID(1)
	if !ok || got.Name != "Go" {
		panic("select")
	}
	if !t.updateName(1, "Golang") {
		panic("update")
	}
	got, _ = t.selectByID(1)
	if got.Name != "Golang" {
		panic("after update")
	}
	if !t.deleteByID(1) || len(t.rows) != 0 {
		panic("delete")
	}
	fmt.Println("✅ SqlBasics 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 SqlBasics ************\n")
	fmt.Println("SQL 语句对照: INSERT→insert, SELECT BY id→selectByID")
	runTests()
}

func main() {
	bootstrap()
}