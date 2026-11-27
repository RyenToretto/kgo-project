package main

import (
	"fmt"
	"sync"
)

/*
[主题]: 至少一次投递 + 消费幂等
[难度]: Medium
[前置]: 09_message_queue/00_async

一、主题说明
MQ 常保证 at-least-once：可能重复投递。消费端须幂等：
业务唯一键去重、或 INSERT IGNORE / 版本号。

二、写法对比
1. 直观版: 重复消息导致扣款两次。
2. 进阶版: processed 集合记录 msgID。

三、真实用途
- 支付回调、订单状态机；Kafka 手动 commit offset 前先落库去重。
资料: https://github.com/IBM/sarama ；https://nats.io/docs/
*/

type consumer struct {
	mu        sync.Mutex
	processed map[string]bool
	applied   int
}

func newConsumer() *consumer {
	return &consumer{processed: make(map[string]bool)}
}

func (c *consumer) handleNaive(msgID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.applied++
}

func (c *consumer) handleIdempotent(msgID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.processed[msgID] {
		return false
	}
	c.processed[msgID] = true
	c.applied++
	return true
}

func runTests() {
	c1 := newConsumer()
	c1.handleNaive("m1")
	c1.handleNaive("m1")
	if c1.applied != 2 {
		panic("naive duplicate")
	}
	c2 := newConsumer()
	if !c2.handleIdempotent("m1") {
		panic("first")
	}
	if c2.handleIdempotent("m1") {
		panic("dup should skip")
	}
	if c2.applied != 1 {
		panic("applied once")
	}
	c2.handleIdempotent("m2")
	if c2.applied != 2 {
		panic("second msg")
	}
	fmt.Println("✅ AtLeastOnceIdempotent 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 AtLeastOnceIdempotent ************\n")
	runTests()
}

func main() {
	bootstrap()
}