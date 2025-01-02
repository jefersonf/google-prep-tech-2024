package module4

func isBipartite(graph [][]int) bool {

	n := len(graph)
	colorSet := make([]int, n)
	UNDEFINED := -1
	for i := range n {
		colorSet[i] = UNDEFINED
	}

	var solve func(node, color int) bool
	solve = func(node, color int) bool {
		colorSet[node] = color
		for _, neighbor := range graph[node] {
			if colorSet[neighbor] == UNDEFINED {
				if !solve(neighbor, color^1) { 
					return false
				}
			}
			if colorSet[neighbor] == colorSet[node] {
				return false
			}
		}
		return true
	}

	for i := range n {
		if colorSet[i] == UNDEFINED && !solve(i, 0) {
			return false
		}
	}

	return true
}
