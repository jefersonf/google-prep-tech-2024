package recursion

func AllSubsets(aa []int) [][]int {
	subsets := make([][]int, 0)
	var gen func(i int, ss []int)
	gen = func(i int, ss []int) {
		subsets = append(subsets, ss)
		for j := i; j < len(aa); j++ {
			gen(j+1, append(ss, aa[j]))
		}
	}
	gen(0, []int{})
	return subsets
}
