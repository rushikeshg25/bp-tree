package main

import "fmt"

func main() {
	tree := NewBPlusTree()

	values := []int{5, 15, 25, 35, 45, 55, 40, 30, 20}
	fmt.Println("Inserting values:", values)
	for _, val := range values {
		tree.Insert(val)
		fmt.Printf("\nAfter inserting %d:\n", val)
		tree.PrintTree()
	}

	fmt.Println("\nFinal B+ Tree structure:")
	tree.PrintTree()

	searchValues := []int{15, 45, 10}
	fmt.Println("\nSearching for values:")
	for _, val := range searchValues {
		if tree.Search(val) {
			fmt.Printf("%d found in the tree\n", val)
		} else {
			fmt.Printf("%d not found in the tree\n", val)
		}
	}
}
