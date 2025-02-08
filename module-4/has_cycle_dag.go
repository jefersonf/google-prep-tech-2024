package module4

func HasCycle(edges [][]int) bool {
	g := make(map[int][]int)
	for _, e := range edges {
		u, v := e[0], e[1]
		if g[u] == nil {
			g[u] = make([]int, 0)
		}
		g[u] = append(g[u], v)
	}

	seen := make(map[int]bool)
	exited := make(map[int]bool)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		seen[u] = true
		for _, v := range g[u] {
			if (seen[v] && !exited[v]) || (!seen[v] && dfs(v)) {
				return true
			}
		}
		exited[u] = true
		return false
	}

	for u := range g {
		if dfs(u) {
			return true
		}
	}
	return false
}
