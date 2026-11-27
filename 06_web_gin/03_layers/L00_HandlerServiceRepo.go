package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
)

/*
[主题]: Handler → Service → Repository 分层
[难度]: Medium
[前置]: 06_web_gin/02_binding

一、主题说明
Handler 只做 HTTP；Service 放业务规则；Repo 只管存取。
Gin 项目常见目录：api/handler、service、repository/model。

二、写法对比
1. 直观版: Handler 里直接 map 增删改查。
2. 进阶版: 接口依赖倒置，便于单测 mock Repo。

三、真实用途
- 可测试、可替换存储；GORM 实现放在 repo 层。
资料: https://go.dev/doc/tutorial/web-service-gin ；https://gin-gonic.com/en/docs/
*/

var errNotFound = errors.New("not found")
var errDuplicate = errors.New("duplicate email")

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepo interface {
	Create(u User) (User, error)
	ByID(id int) (User, error)
}

type memUserRepo struct {
	mu    sync.Mutex
	next  int
	byID  map[int]User
	emails map[string]int
}

func newMemUserRepo() *memUserRepo {
	return &memUserRepo{next: 1, byID: make(map[int]User), emails: make(map[string]int)}
}

func (r *memUserRepo) Create(u User) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.emails[u.Email]; ok {
		return User{}, errDuplicate
	}
	u.ID = r.next
	r.next++
	r.byID[u.ID] = u
	r.emails[u.Email] = u.ID
	return u, nil
}

func (r *memUserRepo) ByID(id int) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.byID[id]
	if !ok {
		return User{}, errNotFound
	}
	return u, nil
}

type UserService struct {
	repo UserRepo
}

func (s *UserService) Register(name, email string) (User, error) {
	if strings.TrimSpace(name) == "" {
		return User{}, errors.New("name required")
	}
	return s.repo.Create(User{Name: name, Email: email})
}

func (s *UserService) Get(id int) (User, error) {
	return s.repo.ByID(id)
}

// 直观版：全堆在 handler
func handlerNaive(store map[int]User, name string) User {
	u := User{ID: len(store) + 1, Name: name}
	store[u.ID] = u
	return u
}

func makeRegisterHandler(svc *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		u, err := svc.Register(body.Name, body.Email)
		if err != nil {
			if errors.Is(err, errDuplicate) {
				http.Error(w, "duplicate", http.StatusConflict)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(u)
	}
}

func runTests() {
	m := make(map[int]User)
	u := handlerNaive(m, "n")
	if u.ID != 1 {
		panic("naive")
	}
	repo := newMemUserRepo()
	svc := &UserService{repo: repo}
	h := makeRegisterHandler(svc)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name":"K","email":"k@x.com"}`))
	h(rec, req)
	if rec.Code != http.StatusCreated {
		panic("register")
	}
	_, err := svc.Get(1)
	if err != nil {
		panic("get after register")
	}
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name":"K2","email":"k@x.com"}`))
	h(rec2, req2)
	if rec2.Code != http.StatusConflict {
		panic("duplicate")
	}
	fmt.Println("✅ HandlerServiceRepo 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 HandlerServiceRepo ************\n")
	runTests()
}

func main() {
	bootstrap()
}