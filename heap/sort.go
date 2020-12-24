package heap

//HeapSort
func HeapSort(li []int, p int, n int) []int {
	mh := NewHeap(li)
	for i := n; i > p; i-- {
		mh[i], mh[0] = mh[0], mh[i]
		heapify(mh, i-1, 0)
	}
	return mh
}
