package middleware

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
