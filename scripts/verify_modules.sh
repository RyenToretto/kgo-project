#!/usr/bin/env bash
# 递归查找 L*.go 学习课，go run 执行（要求 package main 且含 main）。
# 用法:
#   ./scripts/verify_modules.sh 00_go_basics
#   ./scripts/verify_modules.sh 00_go_basics/00_hello
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

if ! command -v go >/dev/null 2>&1; then
  echo "❌ 需要安装 Go（建议最新稳定版 ≥ 1.22）。安装说明: https://go.dev/doc/install"
  exit 1
fi

GO_VER="$(go env GOVERSION 2>/dev/null || go version)"
echo "go toolchain: $GO_VER"

if [[ $# -lt 1 ]]; then
  echo "用法: $0 <目录...>"
  exit 1
fi

total=0
passed=0
failed=0

run_one() {
  local file="$1"
  total=$((total + 1))
  echo "—— 运行: $file"
  if ! go run "$file"; then
    echo "❌ 运行失败: $file"
    failed=$((failed + 1))
    return
  fi
  passed=$((passed + 1))
}

for target in "$@"; do
  abs="$ROOT/$target"
  if [[ ! -e "$abs" ]]; then
    echo "❌ 不存在: $target"
    failed=$((failed + 1))
    continue
  fi
  echo ""
  echo "===== 验证目录: $target ====="
  while IFS= read -r -d '' file; do
    run_one "$file"
  done < <(find "$abs" -type f -name 'L*.go' -print0 | sort -z)
done

echo ""
echo "全部汇总: 通过 $passed / $total ，失败 $failed"
if [[ "$failed" -gt 0 ]] || [[ "$total" -eq 0 ]]; then
  exit 1
fi
