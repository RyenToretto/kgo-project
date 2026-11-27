package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

/*
[主题]: httptest 测 HTTP Handler
[难度]: Easy
[前置]: 04_stdlib_http、11_testing/00_table

一、主题说明
httptest.NewRecorder 录响应；NewRequest 造请求；NewServer 起临时真实端口（本课多用 Recorder）。

二、写法对比
1. 直观版: main 里 listen 真端口，手工 curl。
2. 进阶版: Recorder + 表驱动状态码/JSON。

三、真实用途
- Handler 单测；与 Gin 的 PerformRequest 同构。
资料: https://pkg.go.dev/net/http/httptest ；https://go.dev/doc/tutorial/add-a-test
*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	_ = json.NewEncoder(w).Encode(map[string]string{"greeting": "hello, " + name})
}

func runTests() {
	cases := []struct {
		name       string
		method     string
		path       string
		wantCode   int
		wantSubstr string
	}{
		{"默认", http.MethodGet, "/hello", 200, "world"},
		{"带参", http.MethodGet, "/hello?name=kgo", 200, "kgo"},
		{"POST拒", http.MethodPost, "/hello", 405, ""},
	}
	for _, tc := range cases {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, tc.path, nil)
		helloHandler(rec, req)
		if rec.Code != tc.wantCode {
			panic(tc.name + " code")
		}
		if tc.wantSubstr != "" && !strings.Contains(rec.Body.String(), tc.wantSubstr) {
			panic(tc.name + " body")
		}
	}

	srv := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer srv.Close()
	resp, err := http.Get(srv.URL + "/hello?name=server")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("server")
	}

	fmt.Println("✅ HttpTest 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 HttpTest ************\n")
	runTests()
}

func main() {
	bootstrap()
}