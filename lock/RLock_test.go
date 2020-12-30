package lock

import (
	"sync"
	"testing"
)

func TestRLock(t *testing.T) {
	//var token1, token2 int64 = 3, 5
	var token1 int64 = 3
	var wg sync.WaitGroup

	loop := 10000
	wg.Add(loop)
	rl := new(TokenRecursiveMutex)

	sum := 0
	for i := 0; i < loop; i++ {
		go func() {
			defer wg.Done()
			rl.Lock(token1)
			sum++
		}()
	}
	wg.Wait()
	// fmt.Printf("recursion in lock is %d\n", rl.recursion)

	if sum != loop {
		t.Errorf("try%d doesnot match lock%d", loop, sum)
	}
	// fmt.Printf("lock:%+v", rl)
	for i := 0; i < loop; i++ {
		rl.Unlock(token1)
	}
	// wg.Wait()
}
