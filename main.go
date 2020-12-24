package main

import (
	"fmt"

	"github.com/jungleWhisper/basicAlgo/heap"
)

var toSort = []int{3, 5, 1, 2, 7, 4}

func main() {
	fmt.Printf("list is %v before sort\n", toSort)
	// t := tree.NewBTree(toSort)
	// fmt.Println(t.Inorder(), t.Preorder(), t.Postorder())
	// mh := heap.NewHeap(toSort)
	// p, _ := mh.Pop()
	// fmt.Println(p, mh)
	// BubbleSort(toSort)
	// InsertionSort(toSort)
	// SelectionSort(toSort)
	// sort.QuickSort(toSort, 0, len(toSort)-1)
	a1 := heap.HeapSort(toSort, 0, len(toSort)-1)
	fmt.Println(a1)
	fmt.Printf("list is %v after sort\n", toSort)
}
