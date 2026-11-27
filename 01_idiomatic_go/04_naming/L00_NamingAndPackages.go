package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 命名与包设计直觉
[难度]: Easy
[前置]: 01_idiomatic_go/03_composition

一、主题说明
短小本地变量；包名简短有意义；避免 stutter（如 http.HTTPServer → http.Server）。
Getters 通常不加 Get 前缀。

二、写法对比
1. 直观版: Java 风格长名字 / GetXxx。
2. 进阶版: 地道短名与包前缀表达语义。

三、真实用途
- 代码评审第一印象；Uber 风格指南是团队常用参考。
资料: https://go.dev/wiki/CodeReviewComments ；https://github.com/uber-go/guide/blob/master/style.md
*/
type UserProfileNaive struct {
	UserName string
}

func (u UserProfileNaive) GetUserName() string {
	return u.UserName
}

type User struct {
	Name string
}

func (u User) NameOr(fallback string) string {
	if strings.TrimSpace(u.Name) == "" {
		return fallback
	}
	return u.Name
}

func runTests() {
	n := UserProfileNaive{UserName: "a"}
	if n.GetUserName() != "a" {
		panic("naive getter")
	}
	u := User{Name: "Ada"}
	if u.NameOr("anon") != "Ada" {
		panic("NameOr 有名时应返回 Name")
	}
	if (User{}).NameOr("anon") != "anon" {
		panic("空名应 fallback")
	}
	fmt.Println("✅ NamingAndPackages 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 NamingAndPackages ************\n")
	fmt.Println(User{Name: "Lin"}.NameOr("anon"))
	runTests()
}

func main() {
	bootstrap()
}
