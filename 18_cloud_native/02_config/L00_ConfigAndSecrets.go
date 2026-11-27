package main

import (
	"fmt"
	"strings"
)

/*
[主题]: 云原生配置与密钥心智
[难度]: Easy
[前置]: 14_production_go/00_config

一、主题说明
ConfigMap 放非机密配置；Secret 放密钥（仍需权限与加密管控）。
应用侧仍是读环境变量或挂载文件——本课模拟解析。

二、写法对比
1. 直观版: 把密钥写进镜像。
2. 进阶版: 运行时注入。

三、真实用途
- K8s 部署；禁止把密码提交到 git。
资料: https://go.dev/doc/security/best-practices ；https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
*/
func secretInImageNaive() string { return "password=supersecret" }

func loadInjected(env map[string]string) (dsn string, err error) {
	user := env["DB_USER"]
	pass := env["DB_PASSWORD"]
	host := env["DB_HOST"]
	if user == "" || pass == "" || host == "" {
		return "", fmt.Errorf("missing db env")
	}
	if strings.Contains(pass, " ") {
		return "", fmt.Errorf("invalid password")
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/app", user, pass, host), nil
}

func runTests() {
	if !strings.Contains(secretInImageNaive(), "password=") {
		panic("naive")
	}
	dsn, err := loadInjected(map[string]string{
		"DB_USER": "u", "DB_PASSWORD": "p", "DB_HOST": "db:3306",
	})
	if err != nil || !strings.Contains(dsn, "u:p@tcp") {
		panic("inject")
	}
	_, err = loadInjected(map[string]string{})
	if err == nil {
		panic("should fail")
	}
	fmt.Println("✅ ConfigAndSecrets 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 ConfigAndSecrets ************\n")
	runTests()
}

func main() {
	bootstrap()
}
