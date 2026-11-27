package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

/*
[主题]: database/sql 用法模式（内存对照）
[难度]: Medium
[前置]: 07_persistence/00_sql

一、主题说明
标准流程: sql.Open → Ping → ExecContext / QueryRowContext → Scan → Close。
连接池由 *sql.DB 管理；本课用 MemDB 模拟，不接真实 MySQL。

二、写法对比
1. 直观版: 全局 map，无 context、无 ErrNoRows。
2. 进阶版: Exec/QueryRow + context + sql.ErrNoRows。

三、真实用途
- MySQL: github.com/go-sql-driver/mysql；PostgreSQL: jackc/pgx
资料: https://pkg.go.dev/database/sql ；https://github.com/go-sql-driver/mysql ；https://github.com/jackc/pgx
*/

type MemDB struct {
	data map[int64]string
	seq  int64
}

func NewMemDB() *MemDB {
	return &MemDB{data: make(map[int64]string), seq: 0}
}

// ExecInsert 对照: db.ExecContext(ctx, "INSERT INTO users(name) VALUES (?)", name)
func (m *MemDB) ExecInsert(ctx context.Context, name string) (int64, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	m.seq++
	m.data[m.seq] = name
	return m.seq, nil
}

// QueryName 对照: db.QueryRowContext(ctx, "SELECT name FROM users WHERE id=?", id).Scan(&name)
func (m *MemDB) QueryName(ctx context.Context, id int64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	name, ok := m.data[id]
	if !ok {
		return "", sql.ErrNoRows
	}
	return name, nil
}

func storeNaive(m map[int]string, name string) int {
	id := len(m) + 1
	m[id] = name
	return id
}

func runTests() {
	m := make(map[int]string)
	if storeNaive(m, "a") != 1 {
		panic("naive")
	}
	db := NewMemDB()
	ctx := context.Background()
	id, err := db.ExecInsert(ctx, "Ada")
	if err != nil || id != 1 {
		panic("insert")
	}
	got, err := db.QueryName(ctx, id)
	if err != nil || got != "Ada" {
		panic("query")
	}
	_, err = db.QueryName(ctx, 999)
	if !errors.Is(err, sql.ErrNoRows) {
		panic("no rows")
	}
	cancelCtx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := db.ExecInsert(cancelCtx, "x"); err == nil {
		panic("ctx cancel")
	}
	fmt.Println("✅ DatabaseSqlPattern 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 DatabaseSqlPattern ************\n")
	fmt.Println("真实代码: db, err := sql.Open(\"mysql\", dsn)")
	runTests()
}

func main() {
	bootstrap()
}