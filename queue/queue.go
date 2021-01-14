package queue

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	val interface{}
	nxt unsafe.Pointer
}

//Free-Lock queue
type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{
		head: n,
		tail: n,
	}
}

func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{val: v}
	for {
		tail := load(&q.tail)
		nxt := load(&tail.nxt)
		if tail == load(&q.tail) {
			if nxt == nil {
				if cas(&tail.nxt, nxt, n) {
					cas(&q.tail, tail, n)
				}
			} else {
				cas(&q.tail, tail, nxt)
			}
		}
	}
}

func (q *LKQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		nxt := load(&head.nxt)
		if head == load(&q.head) {
			if head == tail {
				if nxt == nil {
					return nil
				}
				cas(&q.tail, tail, nxt)
			} else {
				if cas(&q.head, head, nxt) {
					return nxt.val
				}
			}
		}
	}

}

func load(ptr *unsafe.Pointer) *node {
	return (*node)(atomic.LoadPointer(ptr))
}

func cas(ptr *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(ptr, unsafe.Pointer(old), unsafe.Pointer(new))
}
