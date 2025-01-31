package recursion

func Permutations[T any](a []T) [][]T {
	n := len(a)
	perms := make([][]T, 0)
	var gen func(i int)
	gen = func(i int) {
		if i == n {
			p := make([]T, n)
			copy(p, a)
			perms = append(perms, p)
			return
		}
		for j := i; j < n; j++ {
			a[i], a[j] = a[j], a[i]
			gen(i + 1)
			a[i], a[j] = a[j], a[i]
		}
	}
	gen(0)
	return perms
}
