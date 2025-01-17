package module2

func RangeMinimumQueryWithSegmentTree(arr []int, queries [][]int) []int {

	n := len(arr)
	segtree := make([]int, 4*n)

	var build func(id, l, r int)
	build = func(id, l, r int) {
		if l == r {
			segtree[id] = arr[l]
			return
		}
		mid := (l + r) / 2
		build(id*2, l, mid)
		build(id*2+1, mid+1, r)
		segtree[id] = min(segtree[id*2], segtree[id*2+1])
	}

	var query func(id, l, r, i, j int) int

	query = func(id, l, r, i, j int) int {
		if l > j || r < i {
			return int(Inf)
		}
		if l >= i && r <= j {
			return segtree[id]
		}
		mid := (l + r) / 2
		leftMin := query(id*2, l, mid, i, j)
		rightMin := query(id*2+1, mid+1, r, i, j)

		return min(leftMin, rightMin)
	}

	build(1, 0, n-1)

	answers := make([]int, len(queries))
	for i, q := range queries {
		answers[i] = query(1, 0, n-1, q[0], q[1])
	}

	return answers
}
