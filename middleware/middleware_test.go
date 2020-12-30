package middleware

import (
	"testing"
)

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
