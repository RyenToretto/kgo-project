package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

/*
[主题]: HTTP Client 与超时
[难度]: Easy
[前置]: 04_stdlib_http/00_handler 、03_context

一、主题说明
默认 Client 无超时——生产必须设 Timeout 或带 ctx 的 NewRequestWithContext。
用 httptest.Server 做本地假服务端离线测。

二、写法对比
1. 直观版: http.Get（难控超时）。
2. 进阶版: 自定义 Client + ctx。

三、真实用途
- 调用下游 HTTP/REST；微服务出站请求。
资料: https://pkg.go.dev/net/http ；https://go.dev/blog/context
*/
func getNaive(url string) (string, error) {
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}

func getBetter(ctx context.Context, client *http.Client, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), err
}

func runTests() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "pong")
	}))
	defer srv.Close()

	body, err := getNaive(srv.URL)
	if err != nil || body != "pong" {
		panic("naive get")
	}
	client := &http.Client{Timeout: 2 * time.Second}
	ctx := context.Background()
	body, err = getBetter(ctx, client, srv.URL)
	if err != nil || !strings.Contains(body, "pong") {
		panic("better get")
	}
	fmt.Println("✅ HttpClient 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 HttpClient ************\n")
	runTests()
}

func main() {
	bootstrap()
}
