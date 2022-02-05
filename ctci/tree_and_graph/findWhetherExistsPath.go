package ctci

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	adjancencyList := make([][]int, n)
	for _, v := range graph {
		source, target := v[0], v[1]
		adjancencyList[source] = append(adjancencyList[source], target)
	}

	// BFS
	searchQueue := adjancencyList[start]
	visited := make(map[int]bool, n)
	for len(searchQueue) > 0 {
		node := searchQueue[0]
		if target == node {
			return true
		}
		visited[node] = true
		searchQueue = searchQueue[1:]

		// find child node
		for _, v := range adjancencyList[node] {
			if !visited[v] {
				searchQueue = append(searchQueue, v)
			}
		}
	}
	return false
}
