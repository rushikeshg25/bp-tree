package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBPlusTreeInsert(t *testing.T) {
	tree := NewBPlusTree()

	// Test single insert
	tree.Insert(5)
	assert.True(t, tree.Search(5), "Should find key 5 after insertion")
	assert.False(t, tree.Search(10), "Should not find non-existent key")

	// Test multiple inserts
	values := []int{15, 25, 35, 45}
	for _, val := range values {
		tree.Insert(val)
		assert.True(t, tree.Search(val), "Should find key %d after insertion", val)
	}
}

func TestBPlusTreeSplitting(t *testing.T) {
	tree := NewBPlusTree()

	// Insert enough values to cause splitting
	values := []int{10, 20, 30, 40, 50}
	for _, val := range values {
		tree.Insert(val)
	}

	// Verify all values are still searchable after splits
	for _, val := range values {
		assert.True(t, tree.Search(val), "Should find key %d after tree splits", val)
	}
}

func TestBPlusTreeSearch(t *testing.T) {
	tree := NewBPlusTree()

	// Test searching in empty tree
	assert.False(t, tree.Search(1), "Should not find key in empty tree")

	// Insert some values
	insertValues := []int{5, 10, 15, 20, 25}
	for _, val := range insertValues {
		tree.Insert(val)
	}

	// Test successful searches
	for _, val := range insertValues {
		assert.True(t, tree.Search(val), "Should find inserted key %d", val)
	}

	// Test unsuccessful searches
	notPresentValues := []int{1, 7, 13, 22, 30}
	for _, val := range notPresentValues {
		assert.False(t, tree.Search(val), "Should not find key %d", val)
	}
}

func TestBPlusTreeOrder(t *testing.T) {
	tree := NewBPlusTree()

	// Insert values in random order
	values := []int{30, 10, 50, 20, 40}
	for _, val := range values {
		tree.Insert(val)
	}

	// Helper function to check if keys in a node are sorted
	assertNodeOrder := func(node *Node) bool {
		for i := 1; i < len(node.keys); i++ {
			if node.keys[i] <= node.keys[i-1] {
				return false
			}
		}
		return true
	}

	// Check root node order
	assert.True(t, assertNodeOrder(tree.root), "Keys in root node should be in order")

	// If root is not leaf, check children
	if !tree.root.isLeaf {
		for _, child := range tree.root.children {
			assert.True(t, assertNodeOrder(child), "Keys in child nodes should be in order")
		}
	}
}

func TestBPlusTreeDuplicateInsert(t *testing.T) {
	tree := NewBPlusTree()

	// Insert same value multiple times
	for i := 0; i < 3; i++ {
		tree.Insert(10)
	}

	assert.True(t, tree.Search(10), "Should find key 10 after multiple insertions")
}

func TestBPlusTreeEdgeCases(t *testing.T) {
	tree := NewBPlusTree()

	// Test with extreme values
	extremeValues := []int{-1000, 1000, 0, -1, 1}
	for _, val := range extremeValues {
		tree.Insert(val)
		assert.True(t, tree.Search(val), "Should handle extreme value %d", val)
	}
}

func TestBPlusTreeSequentialInsert(t *testing.T) {
	tree := NewBPlusTree()

	// Insert sequential values
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
		assert.True(t, tree.Search(i), "Should find sequential key %d", i)
	}
}

func TestBPlusTreeReverseInsert(t *testing.T) {
	tree := NewBPlusTree()

	// Insert values in reverse order
	for i := 10; i >= 1; i-- {
		tree.Insert(i)
		assert.True(t, tree.Search(i), "Should find reverse sequential key %d", i)
	}
}
