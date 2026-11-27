package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
[主题]: 分布式追踪 OpenTelemetry 心智
[难度]: Medium
[前置]: 12_observability/01_metrics、03_concurrency/03_context

一、主题说明
Trace = 多 Span 树；Span 含 trace_id、parent、耗时、属性。
真实 OTel SDK 导出 OTLP；本课内存 Span 列表模拟传播。

二、写法对比
1. 直观版: 日志串无关联 ID。
2. 进阶版: context 传 Span，子调用 child span。

三、真实用途
- 跨服务排障；与 metrics/logs 关联。
资料: https://opentelemetry.io/docs/languages/go/
*/

type span struct {
	Name     string
	TraceID  string
	SpanID   string
	Parent   string
	Start    time.Time
	End      time.Time
	Attrs    map[string]string
}

type tracer struct {
	mu      sync.Mutex
	nextID  int
	spans   []span
	traceID string
}

func newTracer(traceID string) *tracer {
	return &tracer{traceID: traceID}
}

func (t *tracer) id() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.nextID++
	return fmt.Sprintf("s%d", t.nextID)
}

type spanCtx struct {
	tracer   *tracer
	spanID   string
	parentID string
}

type ctxKey struct{}

func startNaive(name string) {
	fmt.Println("log:", name) // 无 trace 关联
}

func startSpan(ctx context.Context, tr *tracer, name string) (context.Context, func()) {
	parent := ""
	if sc, ok := ctx.Value(ctxKey{}).(spanCtx); ok {
		parent = sc.spanID
	}
	sid := tr.id()
	ctx = context.WithValue(ctx, ctxKey{}, spanCtx{tracer: tr, spanID: sid, parentID: parent})
	start := time.Now()
	tr.mu.Lock()
	tr.spans = append(tr.spans, span{
		Name: name, TraceID: tr.traceID, SpanID: sid, Parent: parent, Start: start, Attrs: map[string]string{},
	})
	tr.mu.Unlock()
	return ctx, func() {
		tr.mu.Lock()
		for i := range tr.spans {
			if tr.spans[i].SpanID == sid {
				tr.spans[i].End = time.Now()
				break
			}
		}
		tr.mu.Unlock()
	}
}

func handleRequest(ctx context.Context, tr *tracer) {
	ctx, end := startSpan(ctx, tr, "http.handle")
	defer end()
	db(ctx, tr)
}

func db(ctx context.Context, tr *tracer) {
	ctx, end := startSpan(ctx, tr, "db.query")
	defer end()
	_ = ctx
}

func runTests() {
	startNaive("only log")

	tr := newTracer("trace-1")
	root := context.Background()
	handleRequest(root, tr)
	if len(tr.spans) < 2 {
		panic("need parent child")
	}
	child := tr.spans[len(tr.spans)-1]
	parent := tr.spans[len(tr.spans)-2]
	if child.Parent != parent.SpanID || child.TraceID != "trace-1" {
		panic("hierarchy")
	}
	if child.End.Before(child.Start) {
		panic("timing")
	}

	var buf strings.Builder
	for _, s := range tr.spans {
		buf.WriteString(fmt.Sprintf("%s parent=%s\n", s.Name, s.Parent))
	}
	if !strings.Contains(buf.String(), "db.query") {
		panic("names")
	}

	fmt.Println("✅ TraceOtelMental 所有测试用例通过！")
}

func bootstrap() {
	fmt.Println("\n\n************ L00 TraceOtelMental ************\n")
	tr := newTracer("demo")
	handleRequest(context.Background(), tr)
	fmt.Println("span 数:", len(tr.spans))
	runTests()
}

func main() {
	bootstrap()
}