package module4

func FindCenter(edges [][]int) int {
	inDeg := make(map[int]int)
	n := 1
	for i := range edges {
		u, v := edges[i][0], edges[i][1]
		inDeg[u] += 1
		inDeg[v] += 1
		n = max(n, max(u, v))
	}

	var center int
	for u := 1; u <= n; u++ {
		if inDeg[u] == n-1 {
			center = u
			break
		}
	}
	return center
}
