package main

import (
	"context"
	"errors"
	"fmt"
)

/*
[主题]: 用接口 + 假实现做 mock
[难度]: Medium
[前置]: 01_idiomatic_go/02_interfaces

一、主题说明
依赖接口而非具体类型；测试注入 fake/stub，不连真实 DB/Redis。
生产可用 go.uber.org/mock/gomock 生成；本课手写 Fake。

二、写法对比
1. 直观版: Service 里直接 new 真实 Repo，测试只能集成测。
2. 进阶版: 构造函数注入 Repo 接口，测试用 map 假仓储。

三、真实用途
- 单元测业务逻辑；控制错误/慢调用。
资料: https://github.com/uber-go/mock ；https://pkg.go.dev/testing
*/

type Item struct {
	ID   int
	Name string
}

type ItemRepo interface {
	Get(ctx context.Context, id int) (Item, error)
}

type realRepoMental struct{} // 真实会连 SQL，本课不实现

type fakeRepo struct {
	data map[int]Item
}

func (f *fakeRepo) Get(ctx context.Context, id int) (Item, error) {
	if err := ctx.Err(); err != nil {
		return Item{}, err
	}
	it, ok := f.data[id]
	if !ok {
		return Item{}, errors.New("not found")
	}
	return it, nil
}

type ItemService struct {
	repo ItemRepo
}

func (s *ItemService) Label(ctx context.Context, id int) (string, error) {
	it, err := s.repo.Get(ctx, id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("item:%s", it.Name), nil
}

func newServiceNaive() *ItemService {
	// 无法替换依赖
	_ = realRepoMental{}
	return &ItemService{repo: &fakeRepo{data: map[int]Item{1: {ID: 1, Name: "book"}}}}
}

func newServiceBetter(repo ItemRepo) *ItemService {
	return &ItemService{repo: repo}
}

func runTests() {
	s1 := newServiceNaive()
	l, err := s1.Label(context.Background(), 1)
	if err != nil || l != "item:book" {
		panic("naive svc")
	}

	fake := &fakeRepo{data: map[int]Item{2: {ID: 2, Name: "pen"}}}
	s2 := newServiceBetter(fake)
	l2, err := s2.Label(context.Background(), 2)
	if err != nil || l2 != "item:pen" {
		panic("better svc")
	}
	_, err = s2.Label(context.Background(), 99)
	if err == nil {
		panic("missing id")
	}

	fmt.Println("✅ InterfaceMock 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 InterfaceMock ************\n")
	runTests()
}

func main() {
	bootstrap()
}