package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
[主题]: Prometheus 指标心智（内存 Registry）
[难度]: Easy
[前置]: 12_observability/00_pprof

一、主题说明
Counter 只增、Gauge 可升降、Histogram 分桶延迟。
真实用 prometheus/client_golang；本课模拟 Counter/Gauge 与 exposition 文本格式。

二、写法对比
1. 直观版: fmt.Println 打日志当指标。
2. 进阶版: 命名规范 service_subsystem_unit + labels。

三、真实用途
- RED/USE 方法；告警与 Grafana。
资料: https://prometheus.io/docs/guides/go-application/ ；https://github.com/prometheus/client_golang
*/

type counter struct {
	mu   sync.Mutex
	name string
	val  float64
}

func (c *counter) Inc() {
	c.mu.Lock()
	c.val++
	c.mu.Unlock()
}

func (c *counter) Value() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

type gauge struct {
	mu  sync.Mutex
	val float64
}

func (g *gauge) Set(v float64) {
	g.mu.Lock()
	g.val = v
	g.mu.Unlock()
}

func (g *gauge) Value() float64 {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.val
}

type registryNaive struct {
	logs []string
}

func (r *registryNaive) Record(s string) {
	r.logs = append(r.logs, s)
}

type registryBetter struct {
	counters map[string]*counter
	gauges   map[string]*gauge
}

func newRegistry() *registryBetter {
	return &registryBetter{
		counters: make(map[string]*counter),
		gauges:   make(map[string]*gauge),
	}
}

func (r *registryBetter) Counter(name string) *counter {
	if c, ok := r.counters[name]; ok {
		return c
	}
	c := &counter{name: name}
	r.counters[name] = c
	return c
}

func (r *registryBetter) Gauge(name string) *gauge {
	if g, ok := r.gauges[name]; ok {
		return g
	}
	g := &gauge{}
	r.gauges[name] = g
	return g
}

// 模拟 Prometheus 文本 exposition 片段
func (r *registryBetter) Render() string {
	var b strings.Builder
	for name, c := range r.counters {
		b.WriteString(fmt.Sprintf("# TYPE %s counter\n%s %g\n", name, name, c.Value()))
	}
	for name, g := range r.gauges {
		b.WriteString(fmt.Sprintf("# TYPE %s gauge\n%s %g\n", name, name, g.Value()))
	}
	return b.String()
}

func runTests() {
	naive := &registryNaive{}
	naive.Record("req")
	if len(naive.logs) != 1 {
		panic("naive")
	}

	reg := newRegistry()
	reg.Counter("http_requests_total").Inc()
	reg.Counter("http_requests_total").Inc()
	reg.Gauge("inflight").Set(3)
	if reg.Counter("http_requests_total").Value() != 2 {
		panic("counter")
	}
	out := reg.Render()
	if !strings.Contains(out, "http_requests_total 2") || !strings.Contains(out, "TYPE") {
		panic("render")
	}

	fmt.Println("✅ MetricsPrometheus 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 MetricsPrometheus ************\n")
	reg := newRegistry()
	reg.Counter("demo_total").Inc()
	fmt.Println(reg.Render())
	runTests()
}

func main() {
	bootstrap()
}