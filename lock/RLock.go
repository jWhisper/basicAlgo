package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type TokenRecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

/*
 */
func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.owner) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.owner) != token {
		panic(fmt.Sprintf("token not match: you [%d] are not the owner [%d]", token, m.owner))
	}

	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, 0)
	m.recursion = 0
	m.Mutex.Unlock()
}
