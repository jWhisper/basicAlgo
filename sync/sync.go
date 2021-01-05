package sync

import "fmt"

func SafeGO(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("goroutine error:%v.\n", err)
			}
		}()
		f()
	}()
}
