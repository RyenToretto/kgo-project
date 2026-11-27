package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
)

/*
[主题]: REST CRUD 资源风格（内存商店）
[难度]: Easy
[前置]: 06_web_gin/00_router、04_stdlib_http/01_json

一、主题说明
用名词复数表示资源，HTTP 方法表达动作：GET 查、POST 建、PUT/PATCH 改、DELETE 删。
状态码：201 创建、404 不存在、204 无内容删除成功。

二、写法对比
1. 直观版: 所有操作用 POST + body 里的 action 字段。
2. 进阶版: 方法 + 路径语义化。

三、真实用途
- 用户/订单 CRUD；Gin 同样映射到 Handler。
资料: https://go.dev/doc/tutorial/web-service-gin ；https://gin-gonic.com/en/docs/
*/

type item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type store struct {
	mu    sync.Mutex
	next  int
	items map[int]item
}

func newStore() *store {
	return &store{next: 1, items: make(map[int]item)}
}

func (s *store) create(name string) item {
	s.mu.Lock()
	defer s.mu.Unlock()
	it := item{ID: s.next, Name: name}
	s.items[it.ID] = it
	s.next++
	return it
}

func (s *store) get(id int) (item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	it, ok := s.items[id]
	return it, ok
}

func (s *store) delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return false
	}
	delete(s.items, id)
	return true
}

// 直观版：统一 POST /do
func crudNaive(action, name string) string {
	switch action {
	case "create":
		return "ok:" + name
	default:
		return "unknown"
	}
}

func itemsHandler(s *store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var body struct {
				Name string `json:"name"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "bad json", http.StatusBadRequest)
				return
			}
			it := s.create(body.Name)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(it)
		case http.MethodGet:
			idStr := strings.TrimPrefix(r.URL.Path, "/items/")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			it, ok := s.get(id)
			if !ok {
				http.NotFound(w, r)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(it)
		case http.MethodDelete:
			idStr := strings.TrimPrefix(r.URL.Path, "/items/")
			id, err := strconv.Atoi(idStr)
			if err != nil || !s.delete(id) {
				http.NotFound(w, r)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method", http.StatusMethodNotAllowed)
		}
	}
}

func runTests() {
	if crudNaive("create", "book") != "ok:book" {
		panic("naive")
	}
	s := newStore()
	h := itemsHandler(s)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name":"Go"}`))
	h(rec, req)
	if rec.Code != http.StatusCreated {
		panic("create status")
	}
	var created item
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil || created.ID != 1 {
		panic("create body")
	}
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/items/1", nil)
	h(rec2, req2)
	if rec2.Code != 200 {
		panic("get")
	}
	fmt.Println("✅ RestCrudPattern 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 RestCrudPattern ************\n")
	runTests()
}

func main() {
	bootstrap()
}