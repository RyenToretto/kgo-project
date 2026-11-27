package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

/*
[主题]: JSON API（encoding/json + Handler）
[难度]: Easy
[前置]: 04_stdlib_http/00_handler

一、主题说明
用 struct tag 控制 JSON 字段；Decode 请求体、Encode 响应。
先懂标准库，再上 Gin 绑定会轻松很多。

二、写法对比
1. 直观版: 手写字符串 JSON（易错）。
2. 进阶版: encoding/json。

三、真实用途
- REST API 请求/响应体。
资料: https://pkg.go.dev/encoding/json ；https://gobyexample.com/json
*/
type echoReq struct {
	Name string `json:"name"`
}

type echoResp struct {
	Message string `json:"message"`
}

func echoNaive(body string) string {
	// 极简「假装 JSON」，仅对比
	return `{"message":"hi ` + body + `"}`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var req echoReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(echoResp{Message: "hi " + req.Name})
}

func runTests() {
	if !strings.Contains(echoNaive("x"), "hi x") {
		panic("naive")
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(`{"name":"Ada"}`))
	echoHandler(rec, req)
	if rec.Code != 200 {
		panic("status")
	}
	var resp echoResp
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil || resp.Message != "hi Ada" {
		panic("json resp")
	}
	fmt.Println("✅ JsonApi 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 JsonApi ************\n")
	runTests()
}

func main() {
	bootstrap()
}
