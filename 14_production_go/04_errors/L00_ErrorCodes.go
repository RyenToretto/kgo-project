package main

import (
	"encoding/json"
	"fmt"
)

/*
[主题]: API 错误码与错误体
[难度]: Medium
[前置]: 01_idiomatic_go/00_errors 、04_stdlib_http

一、主题说明
对前端/调用方返回稳定错误码；内部日志保留根因。不要把 SQL 细节直接吐给客户端。

二、写法对比
1. 直观版: 只返回字符串 message。
2. 进阶版: code + message + HTTP 映射。

三、真实用途
- REST/gRPC 统一错误协议。
资料: https://go.dev/doc/security/best-practices ；https://go.dev/wiki/CodeReviewComments
*/
type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func errNaive(msg string) string { return msg }

func toAPIError(kind string) (apiError, int) {
	switch kind {
	case "not_found":
		return apiError{Code: "NOT_FOUND", Message: "resource not found"}, 404
	case "invalid":
		return apiError{Code: "INVALID_ARGUMENT", Message: "bad request"}, 400
	default:
		return apiError{Code: "INTERNAL", Message: "internal error"}, 500
	}
}

func runTests() {
	if errNaive("x") != "x" {
		panic("naive")
	}
	e, code := toAPIError("not_found")
	if code != 404 || e.Code != "NOT_FOUND" {
		panic("not_found mapping")
	}
	b, err := json.Marshal(e)
	if err != nil || string(b) == "" {
		panic("json")
	}
	fmt.Println("✅ ErrorCodes 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ErrorCodes ************\n")
	e, c := toAPIError("invalid")
	fmt.Println(c, e)
	runTests()
}

func main() {
	bootstrap()
}
