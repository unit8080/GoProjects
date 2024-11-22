// 797. All Paths From Source to Target
// https://leetcode.com/problems/all-paths-from-source-to-target/

func dfs(u int, Adj *map[int][]int,visited *map[int]bool, ans *[][]int, list *[]int) {
    if (*visited)[u] {
        return
    }
	(*visited)[u] = true

	paths, _ := (*Adj)[u]
	*list = append(*list, u)

	for _, path := range paths {
		if u >= len(*Adj)-1 { // ???
			tmp := make([]int, len(*list))
			copy(tmp, *list)
			*ans = append(*ans, tmp)
			break
		}
		if (*visited)[path] {
			continue
		}
		dfs(path, Adj, visited, ans, list)
		*list = (*list)[:len(*list)-1]
	}
	delete(*visited, u)
}

func allPathsSourceTarget(graph [][]int) [][]int {

	Adj := make(map[int][]int)
	visited := make(map[int]bool)
	ans := make([][]int, 0)

	// build Adj list
	for u, edges := range graph {
		if len(edges) == 0 {
			Adj[u] = append(Adj[u], -1)
			continue
		}
		for _, vertex := range edges {
			Adj[u] = append(Adj[u], vertex)
		}
	}

	list := make([]int, 0)
	dfs(0, &Adj, &visited, &ans, &list)
	return ans
}

// BFS - needs queue
// queue := make([]interface{}, 0)

/*
   for vertex, _ := range Adj {
       queue = append(queue, vertex)
       visited[vertex] = struct{}{}
       list := []interface{vertex}
   }
*/
