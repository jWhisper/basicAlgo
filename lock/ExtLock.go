package lock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // state是否有锁
	mutexWaken
	mutexStarving
	mutexWaiterShift = iota
)

//TryLock, 解析state字段获取
type ExtMutex struct {
	sync.Mutex
}

func (m *ExtMutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	//在处于唤醒, 饥饿, 有锁的情况下不参与获取锁
	if old&(mutexLocked|mutexWaken|mutexStarving) != 0 {
		return false
	}
	//保留old的信息,并尝试获取锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

//查看有多少个waiters
func (m *ExtMutex) WaiterCount() int {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	waitercount := state >> mutexWaiterShift
	runningCount := state & mutexLocked
	waitercount += runningCount
	return int(waitercount)
}

func (m *ExtMutex) IsWaken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWaken == mutexWaken
}
