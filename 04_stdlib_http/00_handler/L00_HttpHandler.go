package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

/*
[主题]: net/http Handler 与 ServeMux
[难度]: Easy
[前置]: 03_concurrency（建议）

一、主题说明
一切 HTTP 处理都是 Handler 接口：ServeHTTP(ResponseWriter, *Request)。
用 httptest 离线测，无需真监听端口。

二、写法对比
1. 直观版: HandleFunc 闭包。
2. 进阶版: 自定义 Handler 类型。

三、真实用途
- 所有 Web 服务地基；Gin 也是在其之上。
资料: https://pkg.go.dev/net/http ；https://gobyexample.com/http-servers
*/
func helloFunc(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello")
}

type helloHandler struct{}

func (helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "hello-typed")
}

func newMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", helloFunc)
	mux.Handle("/typed", helloHandler{})
	return mux
}

func runTests() {
	mux := newMux()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/hi", nil)
	mux.ServeHTTP(rec, req)
	if rec.Body.String() != "hello" {
		panic("/hi 应为 hello")
	}
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/typed", nil)
	mux.ServeHTTP(rec2, req2)
	if rec2.Body.String() != "hello-typed" {
		panic("/typed 错误")
	}
	fmt.Println("✅ HttpHandler 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 HttpHandler ************\n")
	runTests()
}

func main() {
	bootstrap()
}
