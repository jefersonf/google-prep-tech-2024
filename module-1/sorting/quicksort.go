package sorting

func QuickSort(a []int) {
	qsort(a, 0, len(a)-1)
}

func qsort(a []int, l, r int) {
	if l >= r {
		return
	}
	m := partition(a, l, r)
	qsort(a, l, m-1)
	qsort(a, m+1, r)
}

func partition(a []int, l, r int) int {
	pivot := a[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if a[i] < pivot {
			j++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[j], a[l] = a[l], a[j]
	return j
}
