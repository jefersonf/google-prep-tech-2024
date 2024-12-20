package module4

func findJudge(n int, trust [][]int) int {

	if n == 1 {
		return 1
	}

	inDeg := make([]int, n+1)
	outDeg := make([]int, n+1)

	for _, pair := range trust {
		a, b := pair[0], pair[1]
		inDeg[b] += 1
		outDeg[a] += 1
	}

	for person := range n + 1 {
		if inDeg[person] == n-1 && outDeg[person] == 0 {
			return person
		}
	}

	return -1
}
