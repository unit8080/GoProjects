package main

// DFS
func validPath(n int, edges [][]int, source int, destination int) bool {
	Adj := make(map[interface{}][]interface{})
	visited := make(map[interface{}]struct{})

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		Adj[u] = append(Adj[u], v)
		Adj[v] = append(Adj[v], u)
	}

    if source == destination { return true }

	validPath := false
	visited[source] = struct{}{}
	for _, vertex := range Adj[source] {
		_, visitedAlready := visited[vertex]
		if !visitedAlready {
			validPath = dfs(Adj, visited, vertex, destination)
		}
		if validPath {
			return true
		}
	}
	return false
}
func dfs(adj map[interface{}][]interface{}, visited map[interface{}]struct{},
	vertex interface{}, destination int) bool {

	visited[vertex] = struct{}{}
	if vertex == destination {
		return true
	}

	for _, v := range adj[vertex] {
		_, visitedAlready := visited[v]
		if !visitedAlready {
			validPath := dfs(adj, visited, v, destination)
			if validPath {
				return true
			}
		}
	}
	return false
}
