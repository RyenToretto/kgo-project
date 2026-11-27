package main

import "fmt"

/*
[主题]: 结构体与方法（值接收者 / 指针接收者）
[难度]: Easy
[前置]: 00_go_basics/02_functions

一、主题说明
struct 组合字段；方法挂在类型上。要修改接收者或避免大对象拷贝时用指针接收者。

二、写法对比
1. 直观版: 函数式更新字段（返回新结构体）。
2. 进阶版: 指针方法原地修改 + 值方法只读。

三、真实用途
- 领域模型、DTO、配置结构体；GORM/JSON 也映射到 struct。
资料: https://go.dev/tour/methods ；https://gobyexample.com/structs
*/
type User struct {
	Name string
	Age  int
}

func bumpAgeNaive(u User) User {
	u.Age++
	return u
}

func (u User) Display() string {
	return fmt.Sprintf("%s(%d)", u.Name, u.Age)
}

func (u *User) Birthday() {
	if u == nil {
		return
	}
	u.Age++
}

func runTests() {
	u := User{Name: "Ada", Age: 30}
	u2 := bumpAgeNaive(u)
	if u.Age != 30 || u2.Age != 31 {
		panic("值拷贝更新不应改原对象")
	}
	if u.Display() != "Ada(30)" {
		panic("Display 错误")
	}
	u.Birthday()
	if u.Age != 31 {
		panic("指针方法应修改 Age")
	}
	var p *User
	p.Birthday() // 安全：nil 检查
	fmt.Println("✅ StructsAndMethods 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 StructsAndMethods ************\n")
	u := User{Name: "Lin", Age: 20}
	fmt.Println(u.Display())
	u.Birthday()
	fmt.Println("after birthday:", u.Display())
	runTests()
}

func main() {
	bootstrap()
}
