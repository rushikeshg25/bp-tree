package main

import (
	"fmt"
)

const ORDER = 3

type Node struct {
	// Keys stores the key values in sorted order
	keys []int
	// Children stores pointers to child nodes (internal nodes only)
	children []*Node
	// Next stores pointer to the next leaf node (leaf nodes only)
	next *Node
	// IsLeaf indicates whether this node is a leaf node
	isLeaf bool
}

type BPlusTree struct {
	root *Node
}

func NewBPlusTree() *BPlusTree {
	// Create a new leaf node as root
	return &BPlusTree{
		root: &Node{
			keys:   make([]int, 0),
			isLeaf: true,
		},
	}
}

func (t *BPlusTree) Insert(key int) {
	// If root is full, create new root
	if len(t.root.keys) == ORDER {
		newRoot := &Node{
			keys:     make([]int, 0),
			children: []*Node{t.root},
			isLeaf:   false,
		}
		t.root = newRoot
		t.splitChild(newRoot, 0)
	}
	t.insertNonFull(t.root, key)
}

func (t *BPlusTree) insertNonFull(node *Node, key int) {
	// For leaf nodes, insert the key directly
	if node.isLeaf {
		insertPos := 0
		// Find position to insert the key
		for insertPos < len(node.keys) && node.keys[insertPos] < key {
			insertPos++
		}
		// Insert key at the correct position
		node.keys = append(node.keys, 0)
		copy(node.keys[insertPos+1:], node.keys[insertPos:])
		node.keys[insertPos] = key
		return
	}

	// For internal nodes, find the appropriate child
	childIndex := 0
	for childIndex < len(node.keys) && node.keys[childIndex] < key {
		childIndex++
	}

	// If child is full, split it
	if len(node.children[childIndex].keys) == ORDER {
		t.splitChild(node, childIndex)
		if key >= node.keys[childIndex] {
			childIndex++
		}
	}

	t.insertNonFull(node.children[childIndex], key)
}

// splitChild splits a full child node
func (t *BPlusTree) splitChild(parent *Node, childIndex int) {
	child := parent.children[childIndex]
	newNode := &Node{
		keys:   make([]int, 0),
		isLeaf: child.isLeaf,
	}

	// Split the keys
	mid := len(child.keys) / 2

	if child.isLeaf {
		// For leaf nodes: split keys, keeping all in leaves
		newNode.keys = append(newNode.keys, child.keys[mid:]...)
		child.keys = child.keys[:mid]

		// Maintain the linked list
		newNode.next = child.next
		child.next = newNode

		// Promote a copy of the first key in newNode
		promotedKey := newNode.keys[0]

		// Insert the promoted key into parent
		parent.keys = append(parent.keys, 0)
		copy(parent.keys[childIndex+1:], parent.keys[childIndex:])
		parent.keys[childIndex] = promotedKey
	} else {
		// For internal nodes: promote middle key, split children
		promotedKey := child.keys[mid]
		newNode.keys = append(newNode.keys, child.keys[mid+1:]...)
		child.keys = child.keys[:mid]

		// Split children (n keys -> n+1 children)
		newNode.children = append(newNode.children, child.children[mid+1:]...)
		child.children = child.children[:mid+1]

		// Insert the promoted key into parent
		parent.keys = append(parent.keys, 0)
		copy(parent.keys[childIndex+1:], parent.keys[childIndex:])
		parent.keys[childIndex] = promotedKey
	}

	// Insert new child into parent
	parent.children = append(parent.children, nil)
	copy(parent.children[childIndex+2:], parent.children[childIndex+1:])
	parent.children[childIndex+1] = newNode
}

// Search looks for a key in the B+ tree
func (t *BPlusTree) Search(key int) bool {
	node := t.root

	// Traverse down to leaf node
	for !node.isLeaf {
		pos := 0
		for pos < len(node.keys) && key >= node.keys[pos] {
			pos++
		}
		node = node.children[pos]
	}

	// Search in leaf node
	for _, k := range node.keys {
		if k == key {
			return true
		}
	}
	return false
}

// PrintTree prints the tree structure for debugging
func (t *BPlusTree) PrintTree() {
	t.printNode(t.root, 0)
}

func (t *BPlusTree) printNode(node *Node, level int) {
	fmt.Printf("Level %d: ", level)
	fmt.Printf("Keys: %v ", node.keys)
	if node.isLeaf {
		fmt.Printf("(Leaf)")
	}
	fmt.Println()

	if !node.isLeaf {
		for _, child := range node.children {
			t.printNode(child, level+1)
		}
	}
}
