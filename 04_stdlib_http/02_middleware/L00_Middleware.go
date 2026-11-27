package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

/*
[主题]: HTTP 中间件链式心智
[难度]: Easy
[前置]: 04_stdlib_http/00_handler

一、主题说明
中间件 = 包装 Handler，在前后插入日志/鉴权/恢复等横切逻辑。
Gin 中间件同构：func(c *gin.Context) + Next()。

二、写法对比
1. 直观版: 在每个 Handler 里复制鉴权代码。
2. 进阶版: withAuth 包装。

三、真实用途
- 鉴权、限流、日志、CORS、recover。
资料: https://pkg.go.dev/net/http ；https://gobyexample.com/http-servers
*/
func core(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "core")
}

func withAuth(next http.HandlerFunc, token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Token") != token {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func runTests() {
	h := withAuth(core, "secret")
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h(rec, req)
	if rec.Code != http.StatusUnauthorized {
		panic("无 token 应 401")
	}
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.Header.Set("X-Token", "secret")
	h(rec2, req2)
	if rec2.Code != 200 || !strings.Contains(rec2.Body.String(), "core") {
		panic("有 token 应到 core")
	}
	fmt.Println("✅ Middleware 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Middleware ************\n")
	runTests()
}

func main() {
	bootstrap()
}
