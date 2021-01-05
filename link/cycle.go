package link

import "sync"

//https://www.jianshu.com/p/40812439f41
type CycleSlice struct {
	head int
	tail int
	cap  int
	data []int
	mu   sync.Mutex
}

func NewCycleSlice(cap int) *CycleSlice {
	//head pointer need a blank location
	cap++
	return &CycleSlice{
		cap:  cap,
		data: make([]int, cap, cap),
	}
}

func (cs *CycleSlice) IsEmpty() bool {
	return cs.tail == cs.head
}

func (cs *CycleSlice) IsFull() bool {
	return (cs.tail+1)%cs.cap == cs.head
}

func (cs *CycleSlice) Pop() (int, bool) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if cs.IsEmpty() {
		return 0, false
	}
	nextHead := (cs.head + 1) % cs.cap
	v := cs.data[nextHead]
	cs.head = nextHead
	return v, true
}

//append is tail gohead
func (cs *CycleSlice) Append(v int) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if cs.IsFull() {
		return false
	}
	nextTail := (cs.tail + 1) % cs.cap
	cs.data[nextTail] = v
	cs.tail = nextTail
	return true
}

func (cs *CycleSlice) Len() int {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	return (cs.tail + cs.cap - cs.head) % cs.cap
}

func (cs *CycleSlice) Cap() int {
	return cs.cap - 1
}
