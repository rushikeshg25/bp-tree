# Tech stack

## Languages and runtime

| Item | Version | Evidence |
| --- | --- | --- |
| Go | `1.24.6` Go directive; no separate toolchain directive | [go.mod:3](../../go.mod#L3) |
| Module | `bp-tree` | [go.mod:1](../../go.mod#L1) |
| Command packaging | `package main` | [main.go:1](../../main.go#L1), [lib.go:1](../../lib.go#L1) |

The directive establishes the module's required Go version; it does not record which exact local compiler executed a past test run.

## Frameworks and libraries

| Library | Version and declaration | Role and use |
| --- | --- | --- |
| Go standard library | Go directive at [go.mod:3](../../go.mod#L3) | `fmt` prints the demo/debug tree ([main.go:3](../../main.go#L3), [lib.go:3](../../lib.go#L3)); `testing`, `math/rand`, `reflect`, and `sort` support the model test ([model_test.go:3](../../model_test.go#L3)). |
| `github.com/stretchr/testify` | `v1.11.1`, [go.mod:5](../../go.mod#L5) | Assertions in [lib_test.go:6](../../lib_test.go#L6); not used by runtime code. |
| `github.com/davecgh/go-spew` | `v1.1.1`, indirect, [go.mod:8](../../go.mod#L8) | Declared transitive test dependency; no direct project import. |
| `github.com/pmezard/go-difflib` | `v1.0.0`, indirect, [go.mod:9](../../go.mod#L9) | Declared transitive test dependency; no direct project import. |
| `gopkg.in/yaml.v3` | `v3.0.1`, indirect, [go.mod:10](../../go.mod#L10) | Declared transitive test dependency; the application does not load YAML configuration. |

## Data and infrastructure

The datastore is the custom [heap-resident tree](../../lib.go#L20), not an external database. Node key capacity is a compile-time [ORDER constant](../../lib.go#L7), with no runtime configuration. The [demo](../../main.go#L5) has no service connections; persistence and concurrency are explicitly outside [v1](../../README.md#L19).

## Tooling

| Tool | Role | Evidence |
| --- | --- | --- |
| `go run .` | Compile and execute the demo. | [README:15](../../README.md#L15), [main.go:5](../../main.go#L5) |
| `go build .` | Build the sole command package. | [go.mod:1](../../go.mod#L1), [main.go:1](../../main.go#L1) |
| `go test -race ./...` | Execute example/model tests with race instrumentation. | [README:15](../../README.md#L15), [lib_test.go:9](../../lib_test.go#L9), [model_test.go:10](../../model_test.go#L10) |

There is no checked-in CI workflow, lint configuration, Makefile, container, or deployment manifest in the [baseline file map](03-structure.md#what-lives-where). A passing race-enabled sequential test suite does not establish concurrent-use safety; the [state](../../lib.go#L20) has no synchronization.
