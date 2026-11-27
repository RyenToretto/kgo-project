package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
[主题]: 生产配置加载（环境变量优先）
[难度]: Medium
[前置]: 05_modules_tooling 、13 建议

一、主题说明
12-factor：配置来自环境，不写死密钥。提供默认值；关键项缺失应失败。

二、写法对比
1. 直观版: 全局常量。
2. 进阶版: 从 env 解析结构体配置。

三、真实用途
- 所有线上服务；K8s ConfigMap/Secret 最终也常注入 env。
资料: https://go.dev/doc/security/best-practices ；https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
*/
type Config struct {
	Addr    string
	Timeout int // seconds
}

func loadNaive() Config {
	return Config{Addr: ":8080", Timeout: 5}
}

func loadFromEnv(getenv func(string) string) (Config, error) {
	cfg := Config{
		Addr:    getenv("APP_ADDR"),
		Timeout: 5,
	}
	if cfg.Addr == "" {
		cfg.Addr = ":8080"
	}
	if t := strings.TrimSpace(getenv("APP_TIMEOUT_SEC")); t != "" {
		n, err := strconv.Atoi(t)
		if err != nil || n <= 0 {
			return Config{}, fmt.Errorf("invalid APP_TIMEOUT_SEC")
		}
		cfg.Timeout = n
	}
	return cfg, nil
}

func runTests() {
	if loadNaive().Addr != ":8080" {
		panic("naive")
	}
	cfg, err := loadFromEnv(func(k string) string {
		if k == "APP_ADDR" {
			return ":9090"
		}
		if k == "APP_TIMEOUT_SEC" {
			return "3"
		}
		return ""
	})
	if err != nil || cfg.Addr != ":9090" || cfg.Timeout != 3 {
		panic("env load")
	}
	_, err = loadFromEnv(func(k string) string {
		if k == "APP_TIMEOUT_SEC" {
			return "x"
		}
		return ""
	})
	if err == nil {
		panic("invalid timeout should error")
	}
	_ = os.Getenv // 提醒真实场景用 os.Getenv
	fmt.Println("✅ ConfigLoading 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ConfigLoading ************\n")
	cfg, _ := loadFromEnv(os.Getenv)
	fmt.Println(cfg)
	runTests()
}

func main() {
	bootstrap()
}
