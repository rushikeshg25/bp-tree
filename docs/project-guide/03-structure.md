# Structure

## What lives where

The baseline has ten tracked files in one root folder: four Go files, the Go manifest/checksums, three project documents, and ignore rules. Runtime logic sits entirely in [lib.go](../../lib.go#L1), with a small [main.go](../../main.go#L1) driver; tests live alongside the implementation and share `package main`. This guide adds the nested documentation directory below.

```text
.
├── main.go                 fixed console demo
├── lib.go                  tree representation and all operations
├── lib_test.go             focused insertion/search checks
├── model_test.go           multiset oracle and duplicate lifecycle
├── go.mod / go.sum         module and dependencies
├── README.md               usage and limitations
├── V1.md                   acceptance contract
├── HISTORY.md              delivery chronology
└── docs/project-guide/     this guide
```

## Root

| File | Responsibility | Key exports or definitions | Called/read by |
| --- | --- | --- | --- |
| [main.go](../../main.go#L5) | Construct and exercise a tree, printing structure and membership. | `main` | Go command runtime |
| [lib.go](../../lib.go#L7) | Own node structure, splits, insertion, lookup, ranges, deletion, and debug output. | `ORDER`, `Node`, `BPlusTree`, `NewBPlusTree`; `Insert`, `Search`, `Range`, `Delete`, `PrintTree` | Demo and both test files; helpers call one another internally |
| [lib_test.go](../../lib_test.go#L9) | Eight focused tests for insertion, splitting, lookup, local key order, duplicates, and input ordering. | `TestBPlusTree*` functions | `go test` |
| [model_test.go](../../model_test.go#L10) | Compare mixed mutations to a sorted multiset; check duplicate range/delete lifecycle. | `TestMixedOperations`, `TestRangeDuplicates` | `go test` |
| [go.mod](../../go.mod#L1) | Name module, declare Go version and assertion dependency. | Module `bp-tree`, Go `1.24.6` | Go toolchain |
| [README.md](../../README.md#L1) | State runnable usage, API semantics, and complexity limits. | Documentation | Maintainers |
| [V1.md](../../V1.md#L1) | Define ordered-multiset scope and acceptance. | Documentation | Implementers/reviewers |
| [HISTORY.md](../../HISTORY.md#L1) | Record baseline-to-v1 chronology and verification evidence. | Documentation | Maintainers |

## Guide folder

| File | Responsibility | Key exports | Read by |
| --- | --- | --- | --- |
| [README.md](README.md) | Entry, run commands, tour, and open questions. | None | New maintainers |
| [01-architecture.md](01-architecture.md) | Components, contracts, representation, and limits. | None | Maintainers |
| [02-flow.md](02-flow.md) | Startup and every significant operation. | None | Maintainers tracing behavior |
| [03-structure.md](03-structure.md) | File ownership map. | None | Maintainers finding code |
| [04-tech-stack.md](04-tech-stack.md) | Versions and tooling. | None | Contributors setting up |
| [05-decisions.md](05-decisions.md) | Rationale, tradeoffs, and traps. | None | Contributors changing algorithms |

## Excluded

[go.sum](../../go.sum) is dependency checksum metadata; [.gitignore](../../.gitignore) is ignore boilerplate. Neither defines runtime behavior. No vendored code, generated source, CI configuration, or deployment files are present in the baseline inventory.
