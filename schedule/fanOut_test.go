package schedule

import (
	"math/rand"
	"testing"
	"time"
)

func TestCloseObserve(t *testing.T) {
	src := make(chan interface{}, 6)
	ob1 := make(chan interface{}, 6)
	observes := make([]chan<- interface{}, 0, 5)
	observes = append(observes, ob1)
	closed := make(chan struct{})
	go func() {
		defer func() {
			close(src)
			close(closed)
		}()

		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			src <- i
		}

	}()
	FanOut(src, observes, true)
	<-closed

	for range ob1 {
		//fmt.Println(v)
	}
}
