package module4

func CanFinishAllCourses(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	for _, p := range prerequisites {
		a, b := p[0], p[1]
		if adj[b] == nil {
			adj[b] = make([]int, 0)
		}
		adj[b] = append(adj[b], a)
		indeg[a] += 1
	}

	queue := make([]int, 0)
	for c := range numCourses {
		if indeg[c] == 0 {
			queue = append(queue, c)
		}
	}

	courses := make([]int, 0)

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		courses = append(courses, c)
		for _, p := range adj[c] {
			indeg[p] -= 1
			if indeg[p] == 0 {
				queue = append(queue, p)
			}
		}
	}

	return len(courses) == numCourses
}
