package main

import "fmt"

/*
[主题]: CQRS 读写分离心智
[难度]: Medium
[前置]: 19_architecture/01_events

一、主题说明
Command 改状态；Query 读模型可另存投影。本课：写库 + 读侧缓存投影。

二、写法对比
1. 直观版: 同一模型读写。
2. 进阶版: 命令更新 + 查询走投影。

三、真实用途
- 高读低写报表、复杂查询优化。
资料: https://github.com/joelparkerhenderson/architecture-decision-record
*/
type user struct {
	ID   string
	Name string
}

func getSameModel(users map[string]user, id string) (user, bool) {
	u, ok := users[id]
	return u, ok
}

type writeSide struct{ users map[string]user }
type readSide struct{ names map[string]string }

func (w *writeSide) Rename(id, name string, r *readSide) {
	u := w.users[id]
	u.Name = name
	w.users[id] = u
	r.names[id] = name
}

func runTests() {
	users := map[string]user{"1": {ID: "1", Name: "a"}}
	u, ok := getSameModel(users, "1")
	if !ok || u.Name != "a" {
		panic("same")
	}
	w := &writeSide{users: map[string]user{"1": {ID: "1", Name: "a"}}}
	r := &readSide{names: map[string]string{"1": "a"}}
	w.Rename("1", "b", r)
	if r.names["1"] != "b" || w.users["1"].Name != "b" {
		panic("cqrs")
	}
	fmt.Println("✅ CqrsSplit 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 CqrsSplit ************\n")
	runTests()
}

func main() {
	bootstrap()
}
