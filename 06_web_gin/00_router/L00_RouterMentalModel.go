package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

/*
[主题]: 迷你路由表 → 对照 Gin 路由心智
[难度]: Easy
[前置]: 04_stdlib_http/00_handler

一、主题说明
HTTP 服务核心是「方法 + 路径 → 处理函数」。Gin 的 gin.Engine 本质上也是一张路由表，
支持参数与分组；本课用标准库 map 模拟，离线可测。

二、写法对比
1. 直观版: 一个 Handler 里用 if/switch 按 path 分支。
2. 进阶版: 注册表 + ServeMux 风格匹配。

三、真实用途
- REST API 多路径分发；Gin 用 router.GET/POST、RouterGroup 前缀。
资料: https://go.dev/doc/tutorial/web-service-gin ；https://gin-gonic.com/en/docs/ ；https://github.com/gin-gonic/gin
*/

type routeKey struct {
	method string
	path   string
}

type miniRouter struct {
	routes map[routeKey]http.HandlerFunc
}

func newMiniRouter() *miniRouter {
	return &miniRouter{routes: make(map[routeKey]http.HandlerFunc)}
}

func (m *miniRouter) Handle(method, path string, h http.HandlerFunc) {
	m.routes[routeKey{method, path}] = h
}

// 直观版：巨型 switch（仅演示反模式）
func routeNaive(method, path string) string {
	if method == http.MethodGet && path == "/ping" {
		return "pong"
	}
	if method == http.MethodPost && path == "/users" {
		return "created"
	}
	return "404"
}

func (m *miniRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, ok := m.routes[routeKey{r.Method, r.URL.Path}]
	if !ok {
		http.NotFound(w, r)
		return
	}
	h(w, r)
}

func runTests() {
	if routeNaive(http.MethodGet, "/ping") != "pong" {
		panic("naive ping")
	}
	if routeNaive(http.MethodGet, "/nope") != "404" {
		panic("naive 404")
	}

	r := newMiniRouter()
	r.Handle(http.MethodGet, "/ping", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprint(w, "pong")
	})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	r.ServeHTTP(rec, req)
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "pong") {
		panic("router ping")
	}
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/missing", nil)
	r.ServeHTTP(rec2, req2)
	if rec2.Code != http.StatusNotFound {
		panic("router 404")
	}
	fmt.Println("✅ RouterMentalModel 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 RouterMentalModel ************\n")
	fmt.Println("Gin 对照: engine.GET(\"/ping\", handler) 等价于注册 (GET,/ping)")
	runTests()
}

func main() {
	bootstrap()
}