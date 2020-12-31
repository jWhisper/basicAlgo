package schedule

import "time"

//一个src对应多个chan,就是1:N, 可用于观察者模式, 数据变动后,所有观察者都要进行改变, 类似发布订阅
func FanOut(src <-chan interface{}, observes []chan<- interface{}, sync bool) {
	go func() {

		defer func() {
			for _, ob := range observes {
				close(ob)
			}
		}()

		for v := range src {
			//v := v
			for _, ob := range observes {
				if sync {
					//异步的情况下,不代表所有v都会发送到ob, 如果ob来不及消费, defer就
					//已经执行了, 那么chan的等待者列表就会抛弃,panic
					go func() {
						ob <- v
					}()
				} else {
					ob <- v
				}
			}
		}
		// give time to observes to consume message
		time.Sleep(100 * time.Millisecond)
	}()
}
