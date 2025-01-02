package module4

func FindJudge(n int, trust [][]int) int {

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

	for personID := range n + 1 {
		if inDeg[personID] == n-1 && outDeg[personID] == 0 {
			return personID
		}
	}

	return -1 // no answer
}
