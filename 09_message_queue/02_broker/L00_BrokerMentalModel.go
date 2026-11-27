package main

import (
	"fmt"
	"sync"
)

/*
[主题]: Kafka / NATS Broker 心智对照
[难度]: Medium
[前置]: 09_message_queue/01_idempotent

一、主题说明
Kafka: topic + partition + consumer group，持久化日志，重放。
NATS: subject 轻量 pub/sub，JetStream 可持久化。本课内存 topic 模拟。

二、写法对比
1. 直观版: 单 channel 广播，无分组、无持久。
2. 进阶版: topic 多订阅者 + 偏移记录。

三、真实用途
- 削峰、事件溯源、微服务集成。
资料: https://github.com/IBM/sarama ；https://nats.io/docs/
*/

type envelope struct {
	Topic   string
	Payload string
	Seq     int64
}

type miniBroker struct {
	mu       sync.Mutex
	seq      int64
	byTopic  map[string][]envelope
	offsets  map[string]int64 // consumerGroup -> last seq per topic 简化
}

func newMiniBroker() *miniBroker {
	return &miniBroker{
		byTopic: make(map[string][]envelope),
		offsets: make(map[string]int64),
	}
}

func (b *miniBroker) publish(topic, payload string) envelope {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.seq++
	env := envelope{Topic: topic, Payload: payload, Seq: b.seq}
	b.byTopic[topic] = append(b.byTopic[topic], env)
	return env
}

func (b *miniBroker) consume(group, topic string) (envelope, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := group + "|" + topic
	off := b.offsets[key]
	log := b.byTopic[topic]
	if int(off) >= len(log) {
		return envelope{}, false
	}
	env := log[off]
	b.offsets[key] = off + 1
	return env, true
}

func fanoutNaive(ch chan string, msg string) {
	ch <- msg
}

func runTests() {
	ch := make(chan string, 1)
	fanoutNaive(ch, "hi")
	if <-ch != "hi" {
		panic("naive")
	}
	br := newMiniBroker()
	br.publish("orders", "created:1")
	br.publish("orders", "created:2")
	e1, ok := br.consume("g1", "orders")
	if !ok || e1.Payload != "created:1" {
		panic("c1")
	}
	e2, ok := br.consume("g1", "orders")
	if !ok || e2.Payload != "created:2" {
		panic("c2")
	}
	if _, ok := br.consume("g1", "orders"); ok {
		panic("eof")
	}
	// 另一消费组从头读
	e3, ok := br.consume("g2", "orders")
	if !ok || e3.Seq != e1.Seq {
		panic("group2")
	}
	fmt.Println("✅ BrokerMentalModel 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 BrokerMentalModel ************\n")
	fmt.Println("Kafka: topic/partition/offset；NATS: subject 轻量")
	runTests()
}

func main() {
	bootstrap()
}