package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
[主题]: 事务与回滚心智
[难度]: Medium
[前置]: 07_persistence/01_database_sql

一、主题说明
事务 ACID：要么全部提交要么全部回滚。database/sql 用 BeginTx → Commit/Rollback。
本课内存 Tx 用快照 + defer 回滚演示。

二、写法对比
1. 直观版: 转账两步不设回滚，中间失败数据不一致。
2. 进阶版: Tx 包装，panic/error 时 Rollback。

三、真实用途
- 扣库存+写订单；隔离级别见 17_data_infra。
资料: https://pkg.go.dev/database/sql ；https://gorm.io/docs/
*/

type account struct {
	balance int
}

type ledger struct {
	mu       sync.Mutex
	accounts map[string]*account
}

func newLedger() *ledger {
	return &ledger{accounts: make(map[string]*account)}
}

func (l *ledger) balanceOf(name string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	a, ok := l.accounts[name]
	if !ok {
		return 0
	}
	return a.balance
}

type memTx struct {
	ledger   *ledger
	snapshot map[string]int
	done     bool
}

func (l *ledger) beginTx() *memTx {
	l.mu.Lock()
	defer l.mu.Unlock()
	snap := make(map[string]int)
	for k, a := range l.accounts {
		snap[k] = a.balance
	}
	return &memTx{ledger: l, snapshot: snap}
}

func (tx *memTx) credit(name string, delta int) {
	tx.ledger.mu.Lock()
	defer tx.ledger.mu.Unlock()
	a := tx.ledger.accounts[name]
	if a == nil {
		a = &account{}
		tx.ledger.accounts[name] = a
	}
	a.balance += delta
}

func (tx *memTx) commit() {
	tx.done = true
}

func (tx *memTx) rollback() {
	if tx.done {
		return
	}
	tx.ledger.mu.Lock()
	defer tx.ledger.mu.Unlock()
	for name, bal := range tx.snapshot {
		if bal == 0 {
			delete(tx.ledger.accounts, name)
			continue
		}
		if tx.ledger.accounts[name] == nil {
			tx.ledger.accounts[name] = &account{}
		}
		tx.ledger.accounts[name].balance = bal
	}
	tx.done = true
}

func transferNaive(l *ledger, from, to string, amt int) error {
	if l.balanceOf(from) < amt {
		return errors.New("insufficient")
	}
	l.mu.Lock()
	l.accounts[from].balance -= amt
	// 模拟中途失败未回滚
	if amt == 999 {
		l.mu.Unlock()
		return errors.New("boom")
	}
	l.accounts[to].balance += amt
	l.mu.Unlock()
	return nil
}

func transferTx(l *ledger, from, to string, amt int) error {
	tx := l.beginTx()
	defer func() {
		if !tx.done {
			tx.rollback()
		}
	}()
	if l.balanceOf(from) < amt {
		return errors.New("insufficient")
	}
	tx.credit(from, -amt)
	tx.credit(to, amt)
	tx.commit()
	return nil
}

func runTests() {
	l := newLedger()
	l.accounts["a"] = &account{balance: 100}
	l.accounts["b"] = &account{balance: 0}
	if err := transferTx(l, "a", "b", 40); err != nil {
		panic("transfer ok")
	}
	if l.balanceOf("a") != 60 || l.balanceOf("b") != 40 {
		panic("balances after transfer")
	}
	l2 := newLedger()
	l2.accounts["a"] = &account{balance: 50}
	if err := transferTx(l2, "a", "b", 100); err == nil {
		panic("should fail")
	}
	if l2.balanceOf("a") != 50 {
		panic("rollback on fail")
	}
	fmt.Println("✅ Transactions 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 Transactions ************\n")
	fmt.Println("对照: tx, err := db.BeginTx(ctx, nil)")
	runTests()
}

func main() {
	bootstrap()
}