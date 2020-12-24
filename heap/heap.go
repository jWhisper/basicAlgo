package heap

import (
	"errors"
)

type maxHeap []int

func NewHeap(li []int) maxHeap {
	llen := len(li)
	mh := make(maxHeap, llen)
	copy(mh, li)
	//自下而上遍历
	for i := llen/2 - 1; i >= 0; i-- {
		heapify(mh, llen-1, i)
	}
	return mh
}

//自上而下去平衡
func heapify(li []int, n int, i int) {
	if n < 0 || i < 0 {
		return
	}
	for {
		maxPos := i
		if 2*i+1 <= n && li[maxPos] < li[2*i+1] {
			maxPos = 2*i + 1
		}
		if 2*i+2 <= n && li[maxPos] < li[2*i+2] {
			maxPos = 2*i + 2
		}
		if maxPos == i {
			break
		}
		li[maxPos], li[i] = li[i], li[maxPos]
		i = maxPos
	}
}

//等于从根节点往下从新平衡
func (h *maxHeap) Pop() (int, error) {
	ret, n := 0, len(*h)-1
	if n < 0 || h == nil {
		return ret, errors.New("Empty maxHeap")
	}
	ret = (*h)[0]
	if n < 1 {
		*h = make(maxHeap, 0)
		return ret, nil
	}
	(*h)[0] = (*h)[n]
	//delete the last one
	(*h) = (*h)[:n]
	heapify(*h, n-1, 0)
	return ret, nil
}

//只要和父节点比较就行,比它大,就交换,直至停止或者到根节点
func (h *maxHeap) Add(e int) {
	*h = append(*h, e)
	n := len(*h) - 1
	for i := n; i >= 1; {
		p := (i - 1) / 2
		if (*h)[p] < (*h)[i] {
			(*h)[p], (*h)[i] = (*h)[i], (*h)[p]
			i = p
			continue
		}
		break
	}
}
