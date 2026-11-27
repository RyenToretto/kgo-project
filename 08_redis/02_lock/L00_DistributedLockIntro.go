package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
[主题]: 分布式锁入门与常见坑
[难度]: Hard
[前置]: 08_redis/01_client、03_concurrency/04_mutex

一、主题说明
多实例争抢同一资源需分布式锁。Redis 常用 SET key token NX EX ttl；
释放须校验 token（Lua），避免删别人的锁。本课进程内模拟 NX+TTL。

二、写法对比
1. 直观版: 本地 Mutex（多机无效）。
2. 进阶版: 带 token 的租约锁 + 过期。

三、真实用途
- Redlock 争议多，简单场景 NX+TTL+续期即可；注意锁粒度与超时。
资料: https://redis.io/docs/latest/develop/clients/go/ ；https://github.com/redis/go-redis
*/

type lockStore struct {
	mu    sync.Mutex
	locks map[string]lockRecord
}

type lockRecord struct {
	token     string
	expiresAt time.Time
}

func newLockStore() *lockStore {
	return &lockStore{locks: make(map[string]lockRecord)}
}

func (s *lockStore) tryLock(key, token string, ttl time.Duration) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	if rec, ok := s.locks[key]; ok && now.Before(rec.expiresAt) {
		return false
	}
	s.locks[key] = lockRecord{token: token, expiresAt: now.Add(ttl)}
	return true
}

func (s *lockStore) unlock(key, token string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	rec, ok := s.locks[key]
	if !ok || rec.token != token {
		return false
	}
	delete(s.locks, key)
	return true
}

var localMu sync.Mutex

func lockNaive() {
	localMu.Lock()
	defer localMu.Unlock()
}

func runTests() {
	lockNaive()
	store := newLockStore()
	if !store.tryLock("order:1", "t1", time.Second) {
		panic("first lock")
	}
	if store.tryLock("order:1", "t2", time.Second) {
		panic("should block second")
	}
	if !store.unlock("order:1", "t1") {
		panic("unlock")
	}
	if !store.tryLock("order:1", "t2", time.Second) {
		panic("relock")
	}
	var counter int32
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			tok := fmt.Sprintf("g%d", id)
			for !store.tryLock("shared", tok, 50*time.Millisecond) {
				time.Sleep(time.Millisecond)
			}
			atomic.AddInt32(&counter, 1)
			_ = store.unlock("shared", tok)
		}(i)
	}
	wg.Wait()
	if counter != 5 {
		panic("counter")
	}
	fmt.Println("✅ DistributedLockIntro 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 DistributedLockIntro ************\n")
	fmt.Println("Redis: SET lock_key token NX EX 30")
	runTests()
}

func main() {
	bootstrap()
}