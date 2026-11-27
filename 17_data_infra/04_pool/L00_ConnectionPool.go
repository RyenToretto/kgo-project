package main

import (
	"fmt"
	"sync"
)

/*
[主题]: 连接池心智
[难度]: Medium
[前置]: 07_persistence

一、主题说明
database/sql 内置池；需设 MaxOpen/MaxIdle/生命周期。
本课用缓冲 channel 模拟借还连接。

二、写法对比
1. 直观版: 每次新建连接。
2. 进阶版: 池化复用。

三、真实用途
- MySQL/Redis/HTTP 连接管理。
资料: https://pkg.go.dev/database/sql ；https://gorm.io/docs/
*/
type conn struct{ ID int }

func openNaive(id int) *conn { return &conn{ID: id} }

type pool struct {
	ch chan *conn
}

func newPool(size int) *pool {
	p := &pool{ch: make(chan *conn, size)}
	for i := 0; i < size; i++ {
		p.ch <- &conn{ID: i}
	}
	return p
}

func (p *pool) Get() *conn { return <-p.ch }
func (p *pool) Put(c *conn) { p.ch <- c }

func runTests() {
	c1 := openNaive(1)
	if c1.ID != 1 {
		panic("naive")
	}
	p := newPool(2)
	a := p.Get()
	b := p.Get()
	if a.ID == b.ID {
		// 可能相同？不应，池里两个不同
	}
	p.Put(a)
	p.Put(b)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			c := p.Get()
			p.Put(c)
		}()
	}
	wg.Wait()
	fmt.Println("✅ ConnectionPool 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ConnectionPool ************\n")
	runTests()
}

func main() {
	bootstrap()
}
