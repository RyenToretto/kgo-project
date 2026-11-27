package main

import (
	"fmt"
	"sync"
	"time"
)

/*
[主题]: Cache-Aside 旁路缓存
[难度]: Easy
[前置]: 07_persistence/00_sql

一、主题说明
读: 先查缓存，miss 再查 DB 并回填。写: 先更新 DB，再删缓存（常见策略）。
本课内存 cache + store 模拟，不对接真实 Redis。

二、写法对比
1. 直观版: 只读 DB，每次都慢。
2. 进阶版: TTL 缓存 + 失效删除。

三、真实用途
- 热点商品、用户信息；生产用 go-redis GET/SET/DEL。
资料: https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis
*/

type store struct {
	mu   sync.RWMutex
	data map[string]string
}

func newStore() *store {
	return &store{data: make(map[string]string)}
}

func (s *store) get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok
}

func (s *store) set(key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}

type cacheEntry struct {
	val       string
	expiresAt time.Time
}

type memCache struct {
	mu   sync.RWMutex
	data map[string]cacheEntry
}

func newMemCache() *memCache {
	return &memCache{data: make(map[string]cacheEntry)}
}

func (c *memCache) get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.data[key]
	if !ok || time.Now().After(e.expiresAt) {
		return "", false
	}
	return e.val, true
}

func (c *memCache) set(key, val string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{val: val, expiresAt: time.Now().Add(ttl)}
}

func (c *memCache) del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func readNaive(s *store, key string) (string, bool) {
	return s.get(key)
}

func readCacheAside(cache *memCache, s *store, key string, ttl time.Duration) (string, bool) {
	if v, ok := cache.get(key); ok {
		return v, true
	}
	v, ok := s.get(key)
	if !ok {
		return "", false
	}
	cache.set(key, v, ttl)
	return v, true
}

func writeCacheAside(cache *memCache, s *store, key, val string) {
	s.set(key, val)
	cache.del(key)
}

func runTests() {
	st := newStore()
	st.set("user:1", "Ada")
	if v, ok := readNaive(st, "user:1"); !ok || v != "Ada" {
		panic("naive")
	}
	ca := newMemCache()
	v, ok := readCacheAside(ca, st, "user:1", time.Minute)
	if !ok || v != "Ada" {
		panic("miss fill")
	}
	v2, ok := readCacheAside(ca, st, "user:1", time.Minute)
	if !ok || v2 != "Ada" {
		panic("hit")
	}
	writeCacheAside(ca, st, "user:1", "Ada2")
	if v, _ := readCacheAside(ca, st, "user:1", time.Minute); v != "Ada2" {
		panic("after write")
	}
	fmt.Println("✅ CacheAside 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 CacheAside ************\n")
	runTests()
}

func main() {
	bootstrap()
}