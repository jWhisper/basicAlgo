package middleware

import (
	"fmt"
	"testing"
)

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

func TestOneMiddleware(t *testing.T) {
	log1 := logger1()
	h := log1(handler1)
	h()
}

func TestMoreMiddleware(t *testing.T) {
	log1 := logger1()
	log2 := logger2()
	logMiddleware := Chain(log1, log2)
	h := logMiddleware(handler1)
	h()
}
