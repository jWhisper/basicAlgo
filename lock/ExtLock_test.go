package lock

import (
	"testing"
	"time"
)

func TestExtLock(t *testing.T) {
	extLock := new(ExtMutex)
	//hold a lock for a period of time
	extLock.Lock()
	if extLock.TryLock() {
		t.Errorf("can not get a lock")
	}
	extLock.Unlock()
}

func TestExtLockCount(t *testing.T) {
	elock := new(ExtMutex)
	waiter := 10
	for i := 0; i < waiter; i++ {
		go func() {
			defer elock.Unlock()
			elock.Lock()
			time.Sleep(2 * time.Second)
		}()
	}
	time.Sleep(time.Millisecond * 100)
	count := elock.WaiterCount()
	if count != waiter {
		t.Errorf("count waiter [%d] not match %d", count, waiter)
	}
}
