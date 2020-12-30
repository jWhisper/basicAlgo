package middleware

import "fmt"

type handler func() error

type Middleware func(handler) handler

func Chain(first Middleware, others ...Middleware) Middleware {
	return func(last handler) handler {
		for i := len(others) - 1; i >= 0; i-- {
			last = others[i](last)
		}
		return first(last)
	}
}

//testcase
func handler1() error {
	fmt.Println("handler1 print1:xxx")
	return nil
}

func logger1() Middleware {
	return func(h handler) handler {
		return func() error {
			fmt.Println("logger1 print:in logger1")
			return h()
		}
	}
}

func logger2() Middleware {
	return func(h handler) handler {
		return func() error {
			fmt.Println("logger2 print:in logger2")
			return h()
		}
	}
}
