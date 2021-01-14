package queue

import (
	"sync"
	"testing"
)

//func BenchmarkEnqueue(b *testing.B) {
func TestEnqueue(b *testing.T) {
	q := NewLKQueue()
	var data = []int{1, 2, 3, 4, 5, 6}

	pushdata(q, data)
	ret := popdata(q)

	if sum(data) != sum(ret) {
		b.Error("element error")
	}
}

func pushdata(q *LKQueue, data []int) {
	var wg sync.WaitGroup
	wg.Add(len(data))
	for _, v := range data {
		v := v
		go func() {
			defer wg.Done()
			q.Enqueue(v)
		}()
	}
	wg.Wait()
}

func popdata(q *LKQueue) (ret []int) {
	var wg sync.WaitGroup
	count := 10
	ch := make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch {
			ret = append(ret, v)
		}
	}()

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			for {
				v := q.Dequeue()
				if v == nil {
					return
				}
				ch <- v.(int)
			}
		}()
	}
	wg.Wait()
	return
}

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
