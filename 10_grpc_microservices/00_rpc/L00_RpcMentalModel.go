package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
)

/*
[主题]: RPC 与 HTTP API 心智对比
[难度]: Easy
[前置]: 04_stdlib_http、06_web_gin 心智

一、主题说明
RPC 强调「像调本地函数一样调远程服务」，常带强类型契约（Protobuf）。
REST 用 URL + HTTP 动词，常 JSON；更利于浏览器与缓存。
本课用标准库模拟 RPC 客户端/服务端，不写真实 gRPC 服务器。

二、写法对比
1. 直观版: 处处用 http.Get + 手解析 JSON，路径散落各处。
2. 进阶版: 客户端封装方法，服务端单一 Handler 按方法名路由。

三、真实用途
- 内部微服务常 gRPC；对外网关仍常 REST。
资料: https://grpc.io/docs/languages/go/quickstart/
*/
const rpcMethodHeader = "X-RPC-Method"

type Greeter interface {
	SayHello(ctx context.Context, name string) (string, error)
}

type greeterClient struct {
	base   string
	client *http.Client
}

func (c *greeterClient) SayHello(ctx context.Context, name string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+"/rpc", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set(rpcMethodHeader, "Greeter.SayHello")
	req.Header.Set("X-RPC-Arg", name)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("rpc http %d", resp.StatusCode)
	}
	var out struct {
		Reply string `json:"reply"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.Reply, nil
}

func greeterServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(rpcMethodHeader) != "Greeter.SayHello" {
		http.Error(w, "unknown method", http.StatusBadRequest)
		return
	}
	name := r.Header.Get("X-RPC-Arg")
	if name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]string{"reply": "hello, " + name})
}

func callRESTNaive(base, name string) (string, error) {
	resp, err := http.Get(base + "/hello?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var m map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return "", err
	}
	return m["msg"], nil
}

func runTests() {
	mux := http.NewServeMux()
	mux.HandleFunc("/rpc", greeterServer)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"msg": "hi, " + r.URL.Query().Get("name")})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	ctx := context.Background()
	var g Greeter = &greeterClient{base: srv.URL, client: srv.Client()}
	reply, err := g.SayHello(ctx, "go")
	if err != nil || reply != "hello, go" {
		panic(fmt.Sprintf("rpc want hello, go got %q err %v", reply, err))
	}

	rest, err := callRESTNaive(srv.URL, "go")
	if err != nil || rest != "hi, go" {
		panic("rest naive")
	}

	// 超时：ctx 已取消应失败
	canceled, cancel := context.WithCancel(ctx)
	cancel()
	_, err = g.SayHello(canceled, "x")
	if err == nil {
		panic("canceled ctx should fail")
	}
	if !errors.Is(err, context.Canceled) {
		// 客户端可能包装错误，至少要有 error
		if err.Error() == "" {
			panic("need error on cancel")
		}
	}

	fmt.Println("✅ RpcMentalModel 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 RpcMentalModel ************\n")
	fmt.Println("要点: RPC=强契约+方法名；REST=资源+动词；内部微服务多用 gRPC，边缘常用 HTTP 网关转 REST。")
	runTests()
}

func main() {
	bootstrap()
}