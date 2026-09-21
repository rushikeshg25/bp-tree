# Decisions

The choices below distinguish written rationale from inference based on the implementation.

## Rebuild after deletion

- **What:** Copy survivors, replace the root, and reinsert them; remove one occurrence per call.
- **Evidence:** [Delete:195](../../lib.go#L195) explicitly says rebuilding keeps deletion simple and preserves split/leaf invariants; [lib.go:216](../../lib.go#L216) performs the rebuild.
- **Why:** Reuse insertion correctness rather than implement incremental merge and redistribution in v1. **Confidence: confirmed by code comment and [README:19](../../README.md#L19).**
- **Tradeoff:** Simple mutation path and reusable empty state, at O(n log n) cost for successful deletion plus O(n) temporary values. A missing delete still scans and copies all existing values before returning false.

## Begin ranges at the first leaf

- **What:** Scan from the global leftmost leaf, stopping at the first value above the high bound.
- **Evidence:** The [Range comment](../../lib.go#L178) states that duplicates split across leaves must be retained; [Search](../../lib.go#L136) routes equality right, which is appropriate for membership but would skip earlier equal occurrences in a range.
- **Why:** Preserve duplicate completeness with a straightforward scan. **Confidence: confirmed by comment.**
- **Tradeoff:** Correct inclusive multiset results and independent output storage, but even a narrow range near the end can visit almost all values. A specialized lower-bound descent remains unimplemented.

## Split full nodes before descending

- **What:** Split a full root or child before inserting into it. Leaves retain all values and copy a separator upward; internal nodes move the middle separator upward.
- **Evidence:** [Insert:34](../../lib.go#L34), [insertNonFull:69](../../lib.go#L69), and [splitChild:91](../../lib.go#L91).
- **Why, inferred:** Keep the recursive insertion target non-full and avoid propagating an overflow result back up the call stack. **Confidence: inferred from control flow.**
- **Tradeoff:** Compact insertion logic and maintained leaf links, but leaf and internal splits require distinct rules; changing one without the other can break routing.

## Use a tiny fixed order

- **What:** The maximum key count is a package constant set to three; there is no configurable order.
- **Evidence:** [ORDER:7](../../lib.go#L7) and the [demo's print-after-each-insert loop](../../main.go#L10).
- **Why, inferred:** Frequent splits make the behavior easy to observe with a small input. **Confidence: inferred; no stated rationale accompanies the constant.**
- **Tradeoff:** Readable demonstrations at the cost of many small allocations and a deeper tree than a larger fanout would require. Here `ORDER` means maximum keys, not maximum children.

## Validate against a simple model

- **What:** A deterministic sorted slice is the oracle for mixed insertions/deletions and subsequent ranges/searches.
- **Evidence:** [TestMixedOperations](../../model_test.go#L10) uses seed 42, 1,500 iterations, and keys in [-25, 24]; [TestRangeDuplicates](../../model_test.go#L41) exercises 50 identical values.
- **Why, inferred:** A repeatable behavioral comparison catches interactions that isolated insertion tests miss. **Confidence: inferred from the test design.**
- **Tradeoff:** Good repeatability and duplicate coverage, but one seed and bounded key space do not prove structural invariants or all integer-boundary behavior. The [index](README.md#open-questions) lists specific coverage gaps.

## Gotchas

- Construct with [NewBPlusTree](../../lib.go#L24). A zero-value `BPlusTree` has no root, while operations such as [Insert](../../lib.go#L36) dereference it immediately.
- The exported-looking API lives in [package main](../../lib.go#L1); the README snippet is usable in this package, not as a ready-to-import library.
- Duplicate keys are valid. The older [order-test helper](../../lib_test.go#L74) checks strictly increasing keys only for its unique-key fixture; do not promote that predicate to a universal invariant.
- [PrintTree](../../lib.go#L156) writes directly to stdout and labels recursion depth; it provides neither a formatter result nor a configurable writer.
- [Range](../../lib.go#L173) returns a non-nil empty slice for empty/inverted results; [Delete](../../lib.go#L206) removes only one duplicate. Neither operation returns an error object.
- The [race command](../../README.md#L15) runs sequential tests; it does not grant concurrent safety to [unlocked tree mutations](../../lib.go#L34).

## Conventions

Keep implementation and same-package tests together ([lib_test.go:1](../../lib_test.go#L1), [model_test.go:1](../../model_test.go#L1)). Follow the existing distinction between public tree operations and unexported representation/helper methods ([lib.go:9](../../lib.go#L9), [lib.go:48](../../lib.go#L48)). When changing duplicate behavior, update the [model and lifecycle cases](../../model_test.go#L10) as well as the [documented contract](../../README.md#L17).
