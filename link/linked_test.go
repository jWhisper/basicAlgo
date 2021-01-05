package link

import (
	"testing"
)

func BenchmarkReverseLink(b *testing.B) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	linkedlist := NewLink(arr)
	ReverseLink(linkedlist)
	PrintLink(linkedlist)
}

func TestReverse(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7}
	linkedlist := NewLink(arr)
	ReverseLink(linkedlist)
	PrintLink(linkedlist)
}
