package module4

func isBipartite(graph [][]int) bool {

	n := len(graph)
	colorSet := make([]int, n)
	for i := range n {
		colorSet[i] = -1
	}

	var solve func(node, color int) bool
	solve = func(node, color int) bool {
		colorSet[node] = color
		for _, neigh := range graph[node] {
			if colorSet[neigh] == -1 {
				if !solve(neigh, color^1) {
					return false
				}
			}
			if colorSet[neigh] == colorSet[node] {
				return false
			}
		}
		return true
	}

	for i := range n {
		if colorSet[i] == -1 && !solve(i, 0) {
			return false
		}
	}

	return true
}
