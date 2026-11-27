package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
[主题]: go-redis 客户端 API 心智（内存假客户端）
[难度]: Medium
[前置]: 08_redis/00_cache

一、主题说明
go-redis 常用: rdb.Set(ctx, key, val, ttl)、Get、Del、Incr。
所有命令带 context；Pipeline 批量减少 RTT。

二、写法对比
1. 直观版: 每条命令 new 连接（无池）。
2. 进阶版: 单 Client 复用 + context 超时。

三、真实用途
- github.com/redis/go-redis/v9
资料: https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis
*/

type fakeRedis struct {
	mu      sync.Mutex
	strings map[string]string
	ttl     map[string]time.Time
}

func newFakeRedis() *fakeRedis {
	return &fakeRedis{
		strings: make(map[string]string),
		ttl:     make(map[string]time.Time),
	}
}

func (r *fakeRedis) Set(ctx context.Context, key, val string, exp time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.strings[key] = val
	if exp > 0 {
		r.ttl[key] = time.Now().Add(exp)
	} else {
		delete(r.ttl, key)
	}
	return nil
}

func (r *fakeRedis) Get(ctx context.Context, key string) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if exp, ok := r.ttl[key]; ok && time.Now().After(exp) {
		delete(r.strings, key)
		delete(r.ttl, key)
		return "", errKeyMissing
	}
	v, ok := r.strings[key]
	if !ok {
		return "", errKeyMissing
	}
	return v, nil
}

func (r *fakeRedis) Del(ctx context.Context, keys ...string) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	var n int64
	for _, k := range keys {
		if _, ok := r.strings[k]; ok {
			delete(r.strings, k)
			delete(r.ttl, k)
			n++
		}
	}
	return n, nil
}

var errKeyMissing = fmt.Errorf("redis: nil")

func kvNaive(m map[string]string, k, v string) {
	m[k] = v
}

func runTests() {
	m := make(map[string]string)
	kvNaive(m, "k", "v")
	if m["k"] != "v" {
		panic("naive")
	}
	r := newFakeRedis()
	ctx := context.Background()
	if err := r.Set(ctx, "session:1", "tok", time.Minute); err != nil {
		panic("set")
	}
	got, err := r.Get(ctx, "session:1")
	if err != nil || got != "tok" {
		panic("get")
	}
	if _, err := r.Get(ctx, "nope"); err == nil {
		panic("nil key")
	}
	if n, _ := r.Del(ctx, "session:1"); n != 1 {
		panic("del")
	}
	fmt.Println("✅ GoRedisMental 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 GoRedisMental ************\n")
	fmt.Println("真实: rdb := redis.NewClient(&redis.Options{Addr: \"localhost:6379\"})")
	runTests()
}

func main() {
	bootstrap()
}