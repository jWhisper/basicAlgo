package schedule

import (
	"math/rand"
	"testing"
	"time"
)

func TestCloseObserve(t *testing.T) {
	src := make(chan interface{}, 5)
	ob1 := make(chan interface{}, 10)
	observes := make([]chan<- interface{}, 5)
	observes = append(observes, ob1)
	closed := make(chan struct{})
	go func() {
		defer close(closed)
		for i := 0; i < 5; i++ {
			i := i
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			src <- i
		}
		close(src)
	}()
	FanOut(src, observes, true)
	<-closed
	// for v := range ob1 {
	// 	fmt.Println(v)
	// }
}
