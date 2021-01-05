package link

import (
	"testing"
)

func TestCycle(t *testing.T) {
	cs := NewCycleSlice(5)
	cs.Append(3)
	if cs.Len() != 1 {
		t.Error("cs only has one ele")
	}
	cs.Append(2)
	cs.Append(1)
	cs.Append(0)
	cs.Append(4)
	if ok := cs.Append(6); ok {
		t.Error("cs is full, but not")
	}
	cs.Pop()
	cs.Pop()
	cs.Append(9)
	if cs.Len() != 4 {
		t.Error("cs only has one ele")
	}
}
