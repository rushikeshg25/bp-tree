# B+ tree project guide

> Generated: 2026-09-21 from commit `d3c3323`.

## What this is

This project implements an in-memory integer multiset with insertion, membership lookup, inclusive ranges, and removal of one occurrence ([implementation](../../lib.go#L34)). It is useful for studying a small B+ tree whose records live in linked leaves and whose internal nodes route searches ([splitting](../../lib.go#L81)). It currently ships as a Go `main` package with a hardcoded console demonstration, rather than an importable library ([entry point](../../main.go#L1)).

## Run it

With the Go version declared in [go.mod:3](../../go.mod#L3), run from the repository root:

```bash
go run .
go test -race ./...
go build .
```

The [demo](../../main.go#L5) needs no arguments, environment variables, or services. Dependency resolution is needed for the [test assertion library](../../go.mod#L5); the runtime itself uses only the standard library.

## The five-file tour

There are four Go source files; the fifth stop is the module manifest.

| # | File | Why this one | Then look at |
| --- | --- | --- | --- |
| 1 | [main.go](../../main.go#L5) | Observe construction, insertions, printing, and searches. | [Startup](02-flow.md#startup-and-demo) |
| 2 | [lib.go](../../lib.go#L9) | Follow nodes, splitting, linked scans, and rebuild deletion. | [Architecture](01-architecture.md#components) |
| 3 | [lib_test.go](../../lib_test.go#L9) | See the focused insertion and lookup expectations. | [Source inventory](03-structure.md#root) |
| 4 | [model_test.go](../../model_test.go#L10) | Read the mixed-operation multiset oracle and duplicate checks. | [Verification flow](02-flow.md#build-and-verification) |
| 5 | [go.mod](../../go.mod#L1) | Establish the module, Go directive, and test dependency. | [Tech stack](04-tech-stack.md#languages-and-runtime) |

## Reading order for this guide

1. [Architecture](01-architecture.md) — components, state, and contracts.
2. [Flow](02-flow.md) — follow each operation from call to result.
3. [Structure](03-structure.md) — find the implementation and tests.
4. [Tech stack](04-tech-stack.md) — runtime and tooling dependencies.
5. [Decisions](05-decisions.md) — evidence, tradeoffs, and gotchas.

## Open questions

- Is a reusable importable package intended? The [contract](../../V1.md#L5) describes operations, while the implementation is [package main](../../lib.go#L1).
- Should a future test verify recursive node occupancy, child counts, and equal leaf depth directly? The [model test](../../model_test.go#L30) compares observable range/search results, and the older [order test](../../lib_test.go#L83) only inspects the root and immediate children.
- Should returned-slice independence and insertion after complete deletion get explicit regressions? [Range](../../lib.go#L173) allocates its result and [Delete](../../lib.go#L216) resets the root, but the [lifecycle test](../../model_test.go#L41) does not exercise those two claims directly.
