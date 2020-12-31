package schedule

import "fmt"

//一个src对应多个chan,就是1:N, 可用于观察者模式, 数据变动后,所有观察者都要进行改变, 类似发布订阅
func FanOut(src <-chan interface{}, observes []chan<- interface{}, sync bool) {
	go func() {

		defer func() {
			fmt.Println("close chan")
			for _, ob := range observes {
				close(ob)
			}
		}()

		for v := range src {
			v := v
			for _, ob := range observes {
				if sync {
					go func() {
						ob <- v
					}()
				} else {
					ob <- v
				}
			}
		}
	}()
}
