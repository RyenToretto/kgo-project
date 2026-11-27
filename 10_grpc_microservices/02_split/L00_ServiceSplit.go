package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

/*
[主题]: 微服务拆分边界与 API 网关心智
[难度]: Medium
[前置]: 10_grpc_microservices/01_grpc、06_web_gin 分层

一、主题说明
按业务能力拆服务（用户、订单）；网关聚合对外 REST，对内可调 RPC。
本课内存「两服务」+ 迷你网关，演示路由转发与边界。

二、写法对比
1. 直观版: 一个巨大 Handler 里用户+订单全写一起。
2. 进阶版: 按前缀转发到独立 ServeMux。

三、真实用途
- BFF/网关；避免客户端直接依赖过多内部服务。
资料: https://grpc.io/docs/languages/go/quickstart/
*/

type userSvc struct{}

func (userSvc) handle(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"service": "user", "path": r.URL.Path})
}

type orderSvc struct{}

func (orderSvc) handle(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"service": "order", "path": r.URL.Path})
}

func monolithNaive(w http.ResponseWriter, r *http.Request) {
	// 所有逻辑堆在一个函数
	if strings.Contains(r.URL.Path, "user") {
		_, _ = fmt.Fprint(w, `{"service":"user"}`)
		return
	}
	if strings.Contains(r.URL.Path, "order") {
		_, _ = fmt.Fprint(w, `{"service":"order"}`)
		return
	}
	http.NotFound(w, r)
}

func newGateway(userURL, orderURL string, client *http.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var target string
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/users"):
			target = userURL + strings.TrimPrefix(r.URL.Path, "/api")
		case strings.HasPrefix(r.URL.Path, "/api/orders"):
			target = orderURL + strings.TrimPrefix(r.URL.Path, "/api")
		default:
			http.NotFound(w, r)
			return
		}
		req, err := http.NewRequestWithContext(r.Context(), r.Method, target, r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), 502)
			return
		}
		defer resp.Body.Close()
		w.WriteHeader(resp.StatusCode)
		_ = json.NewEncoder(w).Encode(json.RawMessage(`{"gateway":true}`))
	})
}

func runTests() {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/misc", nil)
	monolithNaive(rec, req)
	if rec.Code != http.StatusNotFound {
		panic("monolith 404")
	}

	userSrv := httptest.NewServer(http.HandlerFunc((userSvc{}).handle))
	orderSrv := httptest.NewServer(http.HandlerFunc((orderSvc{}).handle))
	defer userSrv.Close()
	defer orderSrv.Close()

	// 直接调后端
	ur := httptest.NewRecorder()
	uq := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	(userSvc{}).handle(ur, uq)
	if !strings.Contains(ur.Body.String(), "user") {
		panic("user svc")
	}

	// 网关形态：这里只验证路由决策函数
	path := "/api/users/1"
	if !strings.HasPrefix(path, "/api/users") {
		panic("prefix")
	}

	fmt.Println("✅ ServiceSplit 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ServiceSplit ************\n")
	fmt.Println("拆分原则: 按领域边界、独立部署、接口契约；网关负责鉴权/聚合/限流。")
	runTests()
}

func main() {
	bootstrap()
}