package main

//冒泡排序
func BubbleSort(l []int) {
	llen := len(l)
	if llen <= 1 {
		return
	}

	for j := llen - 1; j > 0; j-- {
		//每次冒泡是否有交换数据
		flag := false
		for i := 1; i <= j; i++ {
			// 每次都调换，是为了保证稳定
			if l[i-1] > l[i] {
				l[i-1], l[i] = l[i], l[i-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

//插入排序
func InsertionSort(l []int) {
	llen := len(l)
	if llen <= 1 {
		return
	}

	for j := 1; j < llen; j++ {
		val := l[j]
		i := j - 1
		for ; i >= 0; i-- {
			if l[i] > val {
				l[i+1] = l[i]
			} else {
				break
			}
		}
		l[i+1] = val
	}
}

//选择排序
func SelectionSort(l []int) {
	llen := len(l)
	if llen <= 1 {
		return
	}

	for i := 0; i < llen; i++ {
		min := i
		for j := i; j < llen; j++ {
			if l[j] < l[min] {
				min = j
			}
		}
		l[i], l[min] = l[min], l[i]
	}
}

func partition(l []int, p int, r int) int {
	return 1
}

// QuickSort
func QuickSort(l []int, p int, r int) {
	if p >= r {
		return
	}
	pviot := partition(l, p, r)
	QuickSort(l, p, pviot-1)
	QuickSort(l, pviot+1, r)
}
