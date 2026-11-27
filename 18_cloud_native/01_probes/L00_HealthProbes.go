package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

/*
[主题]: 存活 / 就绪探针心智
[难度]: Easy
[前置]: 04_stdlib_http

一、主题说明
liveness：进程是否活着；readiness：能否接流量（依赖是否就绪）。
K8s 探针打这两个 HTTP 路径。

二、写法对比
1. 直观版: 只有 /healthz 恒 200。
2. 进阶版: /livez 与 /readyz 分离。

三、真实用途
- 滚动发布、摘除故障实例。
资料: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
*/
func healthzAlwaysOK(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("ok"))
}

type appState struct{ ready bool }

func (a *appState) livez(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("alive"))
}

func (a *appState) readyz(w http.ResponseWriter, r *http.Request) {
	if !a.ready {
		http.Error(w, "not ready", http.StatusServiceUnavailable)
		return
	}
	_, _ = w.Write([]byte("ready"))
}

func runTests() {
	rec := httptest.NewRecorder()
	healthzAlwaysOK(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if rec.Code != 200 {
		panic("healthz")
	}
	a := &appState{ready: false}
	rec2 := httptest.NewRecorder()
	a.readyz(rec2, httptest.NewRequest(http.MethodGet, "/readyz", nil))
	if rec2.Code != 503 {
		panic("not ready")
	}
	a.ready = true
	rec3 := httptest.NewRecorder()
	a.readyz(rec3, httptest.NewRequest(http.MethodGet, "/readyz", nil))
	if rec3.Code != 200 {
		panic("ready")
	}
	fmt.Println("✅ HealthProbes 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 HealthProbes ************\n")
	runTests()
}

func main() {
	bootstrap()
}
