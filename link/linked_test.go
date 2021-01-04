package link

import (
	"fmt"
	"testing"
)

func reverseLink(root *linkNode) *linkNode {
	if root.next == nil {
		return root
	}
	last := reverseLink(root.next)
	root.next.next = root
	root.next = nil
	return last
}

func printLink(root *linkNode) {
	for root != nil {
		fmt.Printf("%d", root.val)
		root = root.next
	}
	fmt.Printf("\n")
}

func BenchmarkReverseLink(b *testing.B) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	linkedlist := NewLink(arr)
	reverseLink(linkedlist)
	printLink(linkedlist)
}

func TestReverse(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	linkedlist := NewLink(arr)
	reverseLink(linkedlist)
	printLink(linkedlist)
}
