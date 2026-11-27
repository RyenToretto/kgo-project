package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

/*
[主题]: 请求绑定与校验心智（对照 Gin binding）
[难度]: Medium
[前置]: 06_web_gin/01_rest

一、主题说明
Gin 的 ShouldBindJSON + validator 标签在标准库中对应：Decode → 业务校验 → 返回 400。
先手动校验再抽象，避免「框架魔法」看不懂。

二、写法对比
1. 直观版: 不校验，空 name 也入库。
2. 进阶版: validateUser + 结构化错误信息。

三、真实用途
- 注册/下单参数校验；生产常用 go-playground/validator，心智与本课相同。
资料: https://gin-gonic.com/en/docs/ ；https://github.com/gin-gonic/gin
*/

type createUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type apiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func bindNaive(body []byte) (string, error) {
	var m map[string]string
	if err := json.Unmarshal(body, &m); err != nil {
		return "", err
	}
	return m["name"], nil
}

func validateUser(u createUserReq) []apiError {
	var errs []apiError
	if strings.TrimSpace(u.Name) == "" {
		errs = append(errs, apiError{Field: "name", Message: "required"})
	}
	if u.Email != "" && !strings.Contains(u.Email, "@") {
		errs = append(errs, apiError{Field: "email", Message: "invalid"})
	}
	return errs
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req createUserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if errs := validateUser(req); len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]any{"errors": errs})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "name": req.Name})
}

func runTests() {
	name, err := bindNaive([]byte(`{"name":"Ada"}`))
	if err != nil || name != "Ada" {
		panic("bindNaive")
	}
	if len(validateUser(createUserReq{})) == 0 {
		panic("need name error")
	}
	if len(validateUser(createUserReq{Name: "x", Email: "bad"})) == 0 {
		panic("need email error")
	}
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"","email":"x"}`))
	createUserHandler(rec, req)
	if rec.Code != http.StatusBadRequest {
		panic("handler 400")
	}
	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"name":"Lin","email":"a@b.c"}`))
	createUserHandler(rec2, req2)
	if rec2.Code != 200 {
		panic("handler ok")
	}
	fmt.Println("✅ BindingValidation 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 BindingValidation ************\n")
	fmt.Println("Gin: c.ShouldBindJSON(&req) + binding tags")
	runTests()
}

func main() {
	bootstrap()
}