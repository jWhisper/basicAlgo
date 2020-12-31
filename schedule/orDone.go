package schedule

import "reflect"

//orDone只有一个完成了,就是全部完成
func OrDone(workChs ...chan interface{}) <-chan interface{} {
	switch len(workChs) {
	case 0:
		return nil
	case 1:
		return workChs[0]
	}

	orDone := make(chan interface{})

	go func() {
		defer close(orDone)

		var cases []reflect.SelectCase
		for i := 0; i < len(workChs); i++ {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(workChs[i]),
			})
		}

		reflect.Select(cases)
	}()

	return orDone
}
