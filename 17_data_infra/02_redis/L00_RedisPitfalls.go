package main

import (
	"fmt"
	"sync"
	"time"
)

/*
[主题]: Redis 缓存三大坑：穿透 / 击穿 / 雪崩（心智）
[难度]: Medium
[前置]: 08_redis

一、主题说明
穿透：查不存在的 key 打穿到 DB；击穿：热点 key 过期瞬间打穿；雪崩：大量 key 同时过期。
本课用内存 map + 互斥重建模拟。

二、写法对比
1. 直观版: 缓存未命中直接查「DB」。
2. 进阶版: 空值缓存 / singleflight 式互斥。

三、真实用途
- 秒杀、详情页缓存。
资料: https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis
*/
type fakeDB struct {
	hits int
	data map[string]string
}

func (d *fakeDB) Get(k string) (string, bool) {
	d.hits++
	v, ok := d.data[k]
	return v, ok
}

type cache struct {
	mu   sync.Mutex
	data map[string]string
	neg  map[string]time.Time // 空值
}

func getNaive(c *cache, db *fakeDB, k string) string {
	c.mu.Lock()
	if v, ok := c.data[k]; ok {
		c.mu.Unlock()
		return v
	}
	c.mu.Unlock()
	v, ok := db.Get(k)
	if !ok {
		return ""
	}
	c.mu.Lock()
	c.data[k] = v
	c.mu.Unlock()
	return v
}

func getBetter(c *cache, db *fakeDB, k string) string {
	c.mu.Lock()
	if v, ok := c.data[k]; ok {
		c.mu.Unlock()
		return v
	}
	if exp, ok := c.neg[k]; ok && time.Now().Before(exp) {
		c.mu.Unlock()
		return ""
	}
	c.mu.Unlock()
	v, ok := db.Get(k)
	c.mu.Lock()
	defer c.mu.Unlock()
	if !ok {
		c.neg[k] = time.Now().Add(time.Minute)
		return ""
	}
	c.data[k] = v
	return v
}

func runTests() {
	db := &fakeDB{data: map[string]string{"a": "1"}}
	c := &cache{data: map[string]string{}, neg: map[string]time.Time{}}
	_ = getNaive(c, db, "missing")
	_ = getNaive(c, db, "missing")
	if db.hits < 2 {
		panic("naive 穿透应多次打 DB")
	}
	db.hits = 0
	c2 := &cache{data: map[string]string{}, neg: map[string]time.Time{}}
	_ = getBetter(c2, db, "nope")
	_ = getBetter(c2, db, "nope")
	if db.hits != 1 {
		panic("空值缓存后应只打 1 次 DB")
	}
	fmt.Println("✅ RedisPitfalls 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 RedisPitfalls ************\n")
	runTests()
}

func main() {
	bootstrap()
}
