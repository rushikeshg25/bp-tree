# B+ tree v1

An in-memory integer multiset with B+ tree insertion, splitting and lookup.

```go
tree := NewBPlusTree()
tree.Insert(3)
tree.Insert(3)
tree.Insert(7)
fmt.Println(tree.Search(3)) // true
fmt.Println(tree.Range(3, 7)) // [3 3 7]
tree.Delete(3) // removes one occurrence, returns true
```

Run the demo with `go run .`; verify with `go test -race ./...`.

Duplicate keys are retained. Range endpoints are inclusive; inverted ranges return an empty slice. Returned slices are independent of the tree. Missing deletes return false. Empty trees remain reusable.

V1 is single-threaded and in memory. Search/insertion descend the tree; range scans walk leaf links from the first leaf (O(n)). Deletion rebuilds the remaining multiset using insertion (O(n log n)); incremental merge/redistribution and persistence are deferred. The deterministic model test checks mixed operations against a sorted multiset, including duplicates and negative values.
