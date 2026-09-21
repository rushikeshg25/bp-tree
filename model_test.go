package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestMixedOperations(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	tree := NewBPlusTree()
	model := []int{}
	for i := 0; i < 1500; i++ {
		k := r.Intn(50) - 25
		if r.Intn(2) == 0 {
			tree.Insert(k)
			model = append(model, k)
			sort.Ints(model)
		} else {
			pos := sort.SearchInts(model, k)
			found := pos < len(model) && model[pos] == k
			if tree.Delete(k) != found {
				t.Fatalf("delete %d", k)
			}
			if found {
				model = append(model[:pos], model[pos+1:]...)
			}
		}
		if got := tree.Range(-100, 100); !reflect.DeepEqual(got, model) {
			t.Fatalf("step %d got %v want %v", i, got, model)
		}
		for key := -25; key < 25; key++ {
			j := sort.SearchInts(model, key)
			if tree.Search(key) != (j < len(model) && model[j] == key) {
				t.Fatalf("search %d", key)
			}
		}
	}
}
func TestRangeDuplicates(t *testing.T) {
	tree := NewBPlusTree()
	for i := 0; i < 50; i++ {
		tree.Insert(7)
	}
	if len(tree.Range(7, 7)) != 50 {
		t.Fatal("lost duplicate")
	}
	for i := 0; i < 50; i++ {
		if !tree.Delete(7) {
			t.Fatal(i)
		}
	}
	if tree.Search(7) || tree.Delete(7) {
		t.Fatal("not empty")
	}
	if len(tree.Range(10, 1)) != 0 {
		t.Fatal("inverted range")
	}
}
