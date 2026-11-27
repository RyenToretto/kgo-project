package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"time"
)

/*
[主题]: 综合小项目 — 鉴权 + CRUD + 内存 Cache-Aside
[难度]: Hard
[前置]: 06_～12_ 全线

一、主题说明
标准库 HTTP 实现 Mini API：Bearer 鉴权、笔记 CRUD、Cache-Aside 读穿。
分层：Handler → Service → Repo + Cache；可观测用请求计数（指标心智）。

二、写法对比
1. 直观版: 全局 map + 无鉴权。
2. 进阶版: 接口分层、统一错误 JSON、缓存失效。

三、真实用途
- 就业作品雏形；对照 Gin/GORM/Redis 替换各层。
资料: https://go.dev/doc/tutorial/web-service-gin ；https://go.dev/doc/security/best-practices ；https://www.gopl.io/
*/

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NoteRepo interface {
	Create(ctx context.Context, title, content string) (Note, error)
	Get(ctx context.Context, id int) (Note, error)
	Update(ctx context.Context, id int, title, content string) (Note, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]Note, error)
}

type memoryRepo struct {
	mu     sync.Mutex
	nextID int
	items  map[int]Note
}

func newMemoryRepo() *memoryRepo {
	return &memoryRepo{items: make(map[int]Note), nextID: 1}
}

func (r *memoryRepo) Create(ctx context.Context, title, content string) (Note, error) {
	if err := ctx.Err(); err != nil {
		return Note{}, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	n := Note{ID: r.nextID, Title: title, Content: content}
	r.nextID++
	r.items[n.ID] = n
	return n, nil
}

func (r *memoryRepo) Get(ctx context.Context, id int) (Note, error) {
	if err := ctx.Err(); err != nil {
		return Note{}, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	n, ok := r.items[id]
	if !ok {
		return Note{}, fmt.Errorf("not found")
	}
	return n, nil
}

func (r *memoryRepo) Update(ctx context.Context, id int, title, content string) (Note, error) {
	if err := ctx.Err(); err != nil {
		return Note{}, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	n, ok := r.items[id]
	if !ok {
		return Note{}, fmt.Errorf("not found")
	}
	n.Title = title
	n.Content = content
	r.items[id] = n
	return n, nil
}

func (r *memoryRepo) Delete(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return fmt.Errorf("not found")
	}
	delete(r.items, id)
	return nil
}

func (r *memoryRepo) List(ctx context.Context) ([]Note, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]Note, 0, len(r.items))
	for _, n := range r.items {
		out = append(out, n)
	}
	return out, nil
}

// 内存缓存（Redis 心智）
type cache struct {
	mu    sync.Mutex
	data  map[int]Note
	ttl   map[int]time.Time
	ttlDur time.Duration
}

func newCache(d time.Duration) *cache {
	return &cache{data: make(map[int]Note), ttl: make(map[int]time.Time), ttlDur: d}
}

func (c *cache) Get(id int) (Note, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if exp, ok := c.ttl[id]; ok && time.Now().After(exp) {
		delete(c.data, id)
		delete(c.ttl, id)
		return Note{}, false
	}
	n, ok := c.data[id]
	return n, ok
}

func (c *cache) Set(n Note) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[n.ID] = n
	c.ttl[n.ID] = time.Now().Add(c.ttlDur)
}

func (c *cache) Invalidate(id int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, id)
	delete(c.ttl, id)
}

type NoteService struct {
	repo  NoteRepo
	cache *cache
}

func (s *NoteService) Get(ctx context.Context, id int) (Note, error) {
	if n, ok := s.cache.Get(id); ok {
		return n, nil
	}
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		return Note{}, err
	}
	s.cache.Set(n)
	return n, nil
}

func (s *NoteService) Create(ctx context.Context, title, content string) (Note, error) {
	n, err := s.repo.Create(ctx, title, content)
	if err != nil {
		return Note{}, err
	}
	s.cache.Set(n)
	return n, nil
}

func (s *NoteService) Update(ctx context.Context, id int, title, content string) (Note, error) {
	n, err := s.repo.Update(ctx, id, title, content)
	if err != nil {
		return Note{}, err
	}
	s.cache.Set(n)
	return n, nil
}

func (s *NoteService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	s.cache.Invalidate(id)
	return nil
}

type apiServer struct {
	svc        *NoteService
	token      string
	reqTotal   int
	reqTotalMu sync.Mutex
}

func (a *apiServer) incReq() {
	a.reqTotalMu.Lock()
	a.reqTotal++
	a.reqTotalMu.Unlock()
}

func (a *apiServer) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") || strings.TrimPrefix(h, "Bearer ") != a.token {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}
		next(w, r)
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func (a *apiServer) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		a.incReq()
		writeJSON(w, 200, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/notes", a.auth(a.notesCollection))
	mux.HandleFunc("/notes/", a.auth(a.noteItem))
	return mux
}

func (a *apiServer) notesCollection(w http.ResponseWriter, r *http.Request) {
	a.incReq()
	ctx := r.Context()
	switch r.Method {
	case http.MethodPost:
		var body struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			writeJSON(w, 400, map[string]string{"error": "bad json"})
			return
		}
		n, err := a.svc.Create(ctx, body.Title, body.Content)
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 201, n)
	case http.MethodGet:
		list, err := a.svc.repo.List(ctx)
		if err != nil {
			writeJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, 200, list)
	default:
		writeJSON(w, 405, map[string]string{"error": "method"})
	}
}

func (a *apiServer) noteItem(w http.ResponseWriter, r *http.Request) {
	a.incReq()
	idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		writeJSON(w, 400, map[string]string{"error": "id"})
		return
	}
	ctx := r.Context()
	switch r.Method {
	case http.MethodGet:
		n, err := a.svc.Get(ctx, id)
		if err != nil {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, n)
	case http.MethodPut:
		var body struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			writeJSON(w, 400, map[string]string{"error": "bad json"})
			return
		}
		n, err := a.svc.Update(ctx, id, body.Title, body.Content)
		if err != nil {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, 200, n)
	case http.MethodDelete:
		if err := a.svc.Delete(ctx, id); err != nil {
			writeJSON(w, 404, map[string]string{"error": "not found"})
			return
		}
		w.WriteHeader(204)
	default:
		writeJSON(w, 405, map[string]string{"error": "method"})
	}
}

func monolithNaive(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "no structure")
}

func runTests() {
	monolithNaive(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))

	repo := newMemoryRepo()
	svc := &NoteService{repo: repo, cache: newCache(time.Minute)}
	api := &apiServer{svc: svc, token: "capstone-secret"}
	srv := httptest.NewServer(api.routes())
	defer srv.Close()

	client := srv.Client()
	authReq := func(method, url string, body string) (*http.Request, error) {
		var r *http.Request
		var err error
		if body == "" {
			r, err = http.NewRequest(method, url, nil)
		} else {
			r, err = http.NewRequest(method, url, strings.NewReader(body))
		}
		if err != nil {
			return nil, err
		}
		r.Header.Set("Authorization", "Bearer capstone-secret")
		r.Header.Set("Content-Type", "application/json")
		return r, nil
	}

	// 无鉴权
	resp, _ := http.Get(srv.URL + "/notes")
	if resp.StatusCode != 401 {
		panic("need auth")
	}
	resp.Body.Close()

	// 创建
	req, _ := authReq(http.MethodPost, srv.URL+"/notes", `{"title":"t1","content":"c1"}`)
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 201 {
		panic("create")
	}
	var created Note
	_ = json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()
	if created.ID != 1 {
		panic("id")
	}

	// 读（走缓存）
	req2, _ := authReq(http.MethodGet, srv.URL+"/notes/1", "")
	resp2, err := client.Do(req2)
	if err != nil || resp2.StatusCode != 200 {
		panic("get")
	}
	resp2.Body.Close()

	// 更新 + 删除
	req3, _ := authReq(http.MethodPut, srv.URL+"/notes/1", `{"title":"t2","content":"c2"}`)
	resp3, _ := client.Do(req3)
	if resp3.StatusCode != 200 {
		panic("update")
	}
	resp3.Body.Close()

	req4, _ := authReq(http.MethodDelete, srv.URL+"/notes/1", "")
	resp4, _ := client.Do(req4)
	if resp4.StatusCode != 204 {
		panic("delete")
	}
	resp4.Body.Close()

	req5, _ := authReq(http.MethodGet, srv.URL+"/notes/1", "")
	resp5, _ := client.Do(req5)
	if resp5.StatusCode != 404 {
		panic("gone")
	}
	resp5.Body.Close()

	if api.reqTotal < 4 {
		panic("metrics count")
	}

	fmt.Println("✅ MiniApiCapstone 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 MiniApiCapstone ************\n")
	fmt.Println("启动心智: 配置 token → NewServer → 优雅退出(context) → 指标/追踪可再叠 12_ 课。")
	runTests()
}

func main() {
	bootstrap()
}