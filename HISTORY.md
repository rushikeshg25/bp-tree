# In-memory ordered integer index history

> Scope: repository baseline through the v1 delivery branch.
> Last updated: 2026-09-21. Dates below retain Git author timezone offsets.

## At a glance

The v1 scope and acceptance contract is recorded in [V1.md](V1.md). Current behavior and limitations are documented in [README.md](README.md).

## Timeline

### 2025-10-12T23:02:28+05:30: Initial commit

- **What happened:** The repository records `Initial commit`.
- **Evidence:** [commit 05868e34ea](https://github.com/rushikeshg25/bp-tree/commit/05868e34ea92fe3fff96a9459948cd4998f8132a).
- **Confidence:** Confirmed by Git history; the title alone does not establish runtime correctness.

### 2026-09-21T13:21:36+05:30: docs: define bp-tree v1 contract

- **What happened:** The repository records `docs: define bp-tree v1 contract`.
- **Evidence:** [commit d5ce4d5195](https://github.com/rushikeshg25/bp-tree/commit/d5ce4d5195078e656ee1eb6d2dae386465b594f5).
- **Confidence:** Confirmed by Git history; the title alone does not establish runtime correctness.

### 2026-09-21T13:23:38+05:30: feat: add duplicate-safe range scans and deletion

- **What happened:** The repository records `feat: add duplicate-safe range scans and deletion`.
- **Evidence:** [commit 69bed24fc9](https://github.com/rushikeshg25/bp-tree/commit/69bed24fc9999f4272b2ca6e21f03b00c93b43e5).
- **Confidence:** Confirmed by Git history; the title alone does not establish runtime correctness.

### 2026-09-21T13:23:38+05:30: test: compare mixed tree operations with a sorted multiset

- **What happened:** The repository records `test: compare mixed tree operations with a sorted multiset`.
- **Evidence:** [commit 19cb8c9fc6](https://github.com/rushikeshg25/bp-tree/commit/19cb8c9fc646b3f90870b08f65172839c6972f07).
- **Confidence:** Confirmed by Git history; the title alone does not establish runtime correctness.

### 2026-09-21T13:27:16+05:30: docs: document bp-tree v1 usage and limitations

- **What happened:** The repository records `docs: document bp-tree v1 usage and limitations`.
- **Evidence:** [commit feb57e49e9](https://github.com/rushikeshg25/bp-tree/commit/feb57e49e9f898f912731f912f511a21e7916e0e).
- **Confidence:** Confirmed by Git history; the title alone does not establish runtime correctness.

## Delivery verification

Passed `go test -race ./...`, including the sorted-multiset model in [model_test.go](model_test.go).

## Turning points

The v1 contract made failure behavior, lifecycle semantics and executable verification part of the delivery. The new tests and README describe the resulting boundaries.

## Open questions

No production deployment or long-running operational validation was performed as part of this delivery.

## 2026-09-21: Maintainer project guide

Added the six-file [project guide](docs/project-guide/README.md), tracing architecture, runtime flows, source structure, dependencies and decisions against the v1 code. Relative paths, source/heading anchors and Mermaid syntax were checked. The guide distinguishes observed behavior from inferred rationale and records remaining limitations.
