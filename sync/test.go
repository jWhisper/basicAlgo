package sync

import (
	"fmt"
	"runtime"
	"sync"
)

func OneThread() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	count := 3
	wg.Add(count)
	for i := 0; i < count; i++ {
		fmt.Printf("%p\n", &i)
		i := i
		fmt.Printf("%p\n", &i)
		go func() {
			fmt.Println("A:", i)
			wg.Done()
		}()
		//time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}
