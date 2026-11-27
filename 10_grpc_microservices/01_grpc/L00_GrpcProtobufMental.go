package main

import (
	"encoding/json"
	"fmt"
)

/*
[主题]: gRPC 与 Protobuf 契约心智（内存编解码）
[难度]: Easy
[前置]: 10_grpc_microservices/00_rpc

一、主题说明
Protobuf 定义 .proto → 生成 Go struct；线上 gRPC 用二进制 Protobuf 传参。
学习阶段用 JSON 模拟「序列化契约」，理解字段编号、版本兼容直觉即可。

二、写法对比
1. 直观版: map[string]interface{} 传参，字段名靠约定。
2. 进阶版: 显式 Message 结构 + 编解码函数。

三、真实用途
- 高频内部 RPC、流式；需 protoc + grpc 插件生成代码。
资料: https://protobuf.dev/getting-started/gotutorial/ ；https://grpc.io/docs/languages/go/quickstart/
*/

// 心智契约：等价于 message HelloRequest { string name = 1; }
type HelloRequest struct {
	Name string `json:"name"` // proto field 1
}

type HelloReply struct {
	Message string `json:"message"` // proto field 1
}

func encodeNaive(name string) []byte {
	// 无结构：改字段易漏
	b, _ := json.Marshal(map[string]string{"n": name})
	return b
}

func encodeBetter(req HelloRequest) ([]byte, error) {
	return json.Marshal(req)
}

func decodeBetter(b []byte) (HelloReply, error) {
	var req HelloRequest
	if err := json.Unmarshal(b, &req); err != nil {
		return HelloReply{}, err
	}
	if req.Name == "" {
		return HelloReply{}, fmt.Errorf("name required")
	}
	return HelloReply{Message: "hello, " + req.Name}, nil
}

// 模拟 gRPC unary：请求字节 → 响应字节
func unaryCall(reqBytes []byte) ([]byte, error) {
	var req HelloRequest
	if err := json.Unmarshal(reqBytes, &req); err != nil {
		return nil, err
	}
	reply, err := decodeBetter(reqBytes)
	if err != nil {
		return nil, err
	}
	return json.Marshal(reply)
}

func runTests() {
	if len(encodeNaive("a")) == 0 {
		panic("naive")
	}
	body, err := encodeBetter(HelloRequest{Name: "proto"})
	if err != nil {
		panic(err)
	}
	out, err := unaryCall(body)
	if err != nil {
		panic(err)
	}
	var rep HelloReply
	if err := json.Unmarshal(out, &rep); err != nil || rep.Message != "hello, proto" {
		panic("unary")
	}
	_, err = unaryCall(encodeNaive(""))
	if err == nil {
		panic("invalid should fail")
	}
	fmt.Println("✅ GrpcProtobufMental 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GrpcProtobufMental ************\n")
	b, _ := encodeBetter(HelloRequest{Name: "kgo"})
	rep, _ := decodeBetter(b)
	fmt.Println("解码演示:", rep.Message)
	runTests()
}

func main() {
	bootstrap()
}