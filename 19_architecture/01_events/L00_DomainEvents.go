package main

import "fmt"

/*
[主题]: 领域事件
[难度]: Medium
[前置]: 19_architecture/00_ddd

一、主题说明
状态变更发布事件，下游订阅解耦。本课内存 event bus。

二、写法对比
1. 直观版: 直接调用下游函数。
2. 进阶版: 发布事件再分发。

三、真实用途
- 下单后发邮件/积分/库存；对接 MQ。
资料: https://nats.io/docs/ ；https://sre.google/books/
*/
type Event struct {
	Name string
	Payload string
}

func notifyNaive(email *[]string, user string) {
	*email = append(*email, "mail:"+user)
}

type bus struct {
	handlers map[string][]func(Event)
}

func newBus() *bus { return &bus{handlers: map[string][]func(Event){}} }

func (b *bus) On(name string, h func(Event)) {
	b.handlers[name] = append(b.handlers[name], h)
}

func (b *bus) Publish(e Event) {
	for _, h := range b.handlers[e.Name] {
		h(e)
	}
}

func runTests() {
	var mails []string
	notifyNaive(&mails, "a")
	if len(mails) != 1 {
		panic("naive")
	}
	b := newBus()
	var got []string
	b.On("OrderPlaced", func(e Event) { got = append(got, e.Payload) })
	b.Publish(Event{Name: "OrderPlaced", Payload: "o1"})
	if len(got) != 1 || got[0] != "o1" {
		panic("bus")
	}
	fmt.Println("✅ DomainEvents 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 DomainEvents ************\n")
	runTests()
}

func main() {
	bootstrap()
}
