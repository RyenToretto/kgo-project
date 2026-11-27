package main

import (
	"fmt"
	"sync"
)

/*
[主题]: 异步解耦（channel 对照 MQ）
[难度]: Easy
[前置]: 03_concurrency/01_channel

一、主题说明
同步调用链路过长会拖慢接口；把「发邮件/对账」扔到队列由 worker 消费。
本课用 buffered channel 模拟 broker，无 Kafka/NATS 进程。

二、写法对比
1. 直观版: HTTP 里直接 sleep 发邮件，超时。
2. 进阶版: Publish 后立即返回，worker 异步处理。

三、真实用途
- NATS Publish/Subscribe；Kafka Producer/Consumer。
资料: https://nats.io/docs/ ；https://github.com/IBM/sarama
*/

type job struct {
	ID   int
	Kind string
}

type asyncBus struct {
	ch chan job
	wg sync.WaitGroup
}

func newAsyncBus(buf int) *asyncBus {
	return &asyncBus{ch: make(chan job, buf)}
}

func (b *asyncBus) startWorker() {
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()
		for j := range b.ch {
			_ = j // 模拟处理
		}
	}()
}

func (b *asyncBus) publish(j job) {
	b.ch <- j
}

func (b *asyncBus) close() {
	close(b.ch)
	b.wg.Wait()
}

func sendMailSync() int {
	// 假装耗时
	return 1
}

func runTests() {
	if sendMailSync() != 1 {
		panic("sync")
	}
	bus := newAsyncBus(4)
	bus.startWorker()
	bus.publish(job{ID: 1, Kind: "email"})
	bus.publish(job{ID: 2, Kind: "email"})
	bus.close()
	fmt.Println("✅ AsyncDecouple 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 AsyncDecouple ************\n")
	fmt.Println("MQ 对照: producer 写 topic，consumer group 拉取")
	runTests()
}

func main() {
	bootstrap()
}