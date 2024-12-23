package sorting

func MergeSort(a []int) int {
	n := len(a)
	tmp := make([]int, n)
	return mergeSortAndCountInversions(a, tmp, 0, n-1)
}

func mergeSortAndCountInversions(a, t []int, l, r int) (inversionCount int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	inversionCount += mergeSortAndCountInversions(a, t, l, mid)
	inversionCount += mergeSortAndCountInversions(a, t, mid+1, r)
	inversionCount += mergeAndCountInversions(a, t, l, mid+1, r)
	return
}

func mergeAndCountInversions(a, t []int, l, m, r int) (inversionCount int) {
	i, j, k := l, m, l
	for i < m && j <= r {
		if a[i] <= a[j] {
			t[k] = a[i]
			i += 1
		} else {
			t[k] = a[j]
			j += 1
			inversionCount += m - i
		}
		k += 1
	}

	for i < m {
		t[k] = a[i]
		k += 1
		i += 1
	}

	for j <= r {
		t[k] = a[j]
		k += 1
		j += 1
	}

	for l <= r {
		a[l] = t[l]
		l += 1
	}

	return
}
