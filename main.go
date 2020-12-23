package main

import (
	"fmt"
)

var toSort = []int{3, 5, 1, 2, 7, 4}

type node struct {
	val         int
	left, right *node
}

func main() {
	t := NewBTree(toSort)
	fmt.Println(t.inorder(), t.preorder(), t.postorder())
	fmt.Printf("list is %v before sort\n", toSort)
	// BubbleSort(toSort)
	// InsertionSort(toSort)
	// SelectionSort(toSort)
	QuickSort(toSort, 0, len(toSort)-1)
	fmt.Printf("list is %v after sort\n", toSort)
}
