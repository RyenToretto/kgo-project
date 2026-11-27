package main

import (
	"fmt"
	"strings"
)

/*
[主题]: ORM / GORM 心智对照
[难度]: Medium
[前置]: 07_persistence/02_tx

一、主题说明
GORM 用 struct tag 映射表、链式 API：db.Where(...).First(&u)。
本质是生成 SQL + 扫描结果；复杂查询仍需懂 SQL。

二、写法对比
1. 直观版: 手写 SQL 字符串拼接（注入风险）。
2. 进阶版: 参数占位 + 结构体映射（GORM 自动完成）。

三、真实用途
- 快速 CRUD；迁移 AutoMigrate；注意 N+1 与事务。
资料: https://gorm.io/docs/ ；https://pkg.go.dev/database/sql
*/

type UserModel struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"size:64"`
}

type miniORM struct {
	table string
	rows  map[int64]UserModel
	seq   int64
}

func newMiniORM() *miniORM {
	return &miniORM{table: "users", rows: make(map[int64]UserModel)}
}

// Create 对照: db.Create(&user)
func (o *miniORM) Create(name string) UserModel {
	o.seq++
	u := UserModel{ID: o.seq, Name: name}
	o.rows[u.ID] = u
	return u
}

// FirstByID 对照: db.First(&u, id)
func (o *miniORM) FirstByID(id int64) (UserModel, bool) {
	u, ok := o.rows[id]
	return u, ok
}

// WhereName 对照: db.Where("name = ?", name).First(&u)
func (o *miniORM) WhereName(name string) (UserModel, bool) {
	for _, u := range o.rows {
		if u.Name == name {
			return u, true
		}
	}
	return UserModel{}, false
}

func insertNaive(name string) string {
	// 反模式：拼接 SQL
	return "INSERT INTO users(name) VALUES('" + name + "')"
}

func insertBetter(name string) string {
	// 参数化
	return fmt.Sprintf("INSERT INTO users(name) VALUES(%q)", name)
}

func runTests() {
	if !strings.Contains(insertNaive("O'Reilly"), "O'Reilly") {
		panic("naive demo")
	}
	if insertBetter(`a"b`) == "" {
		panic("better")
	}
	o := newMiniORM()
	u := o.Create("Gopher")
	if u.ID != 1 {
		panic("create")
	}
	got, ok := o.FirstByID(1)
	if !ok || got.Name != "Gopher" {
		panic("first")
	}
	got2, ok := o.WhereName("Gopher")
	if !ok || got2.ID != 1 {
		panic("where")
	}
	fmt.Println("✅ GormMentalModel 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GormMentalModel ************\n")
	fmt.Println("GORM: gorm.Open(mysql.Open(dsn), &gorm.Config{})")
	runTests()
}

func main() {
	bootstrap()
}