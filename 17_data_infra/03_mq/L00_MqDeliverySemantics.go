package main

import "fmt"

/*
[主题]: MQ 投递语义：至少一次 / 至多一次 / 正好一次心智
[难度]: Medium
[前置]: 09_message_queue

一、主题说明
至少一次最常见：可能重复 → 消费者必须幂等。
正好一次是端到端目标，常靠幂等键 + 去重表逼近。

二、写法对比
1. 直观版: 处理消息不记状态（重复会双写）。
2. 进阶版: 幂等键去重。

三、真实用途
- 订单、支付回调、库存扣减。
资料: https://nats.io/docs/ ；https://github.com/IBM/sarama
*/
func handleNaive(processed *[]string, id string) {
	*processed = append(*processed, id)
}

func handleIdempotent(seen map[string]bool, processed *[]string, id string) {
	if seen[id] {
		return
	}
	seen[id] = true
	*processed = append(*processed, id)
}

func runTests() {
	var a []string
	handleNaive(&a, "m1")
	handleNaive(&a, "m1")
	if len(a) != 2 {
		panic("naive 允许重复")
	}
	seen := map[string]bool{}
	var b []string
	handleIdempotent(seen, &b, "m1")
	handleIdempotent(seen, &b, "m1")
	if len(b) != 1 {
		panic("幂等应去重")
	}
	fmt.Println("✅ MqDeliverySemantics 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 MqDeliverySemantics ************\n")
	runTests()
}

func main() {
	bootstrap()
}
