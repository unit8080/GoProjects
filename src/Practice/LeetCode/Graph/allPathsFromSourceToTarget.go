// 797. All Paths From Source to Target
// https://leetcode.com/problems/all-paths-from-source-to-target/

func dfs(vertex interface{}, Adj *map[interface{}][]interface{},

	visited *map[interface{}]struct{}, ans *[][]int,
	list *[]int) {

	_, alreadyVisited := (*visited)[vertex.(int)]
	if alreadyVisited {
		return
	}
	(*visited)[vertex.(int)] = struct{}{}

	paths, _ := (*Adj)[vertex]
	// fmt.Println("Vertex:", vertex, "paths:", paths)
	*list = append(*list, vertex.(int))
	// fmt.Println("List:", *list)

	for _, path := range paths {

		if vertex.(int) >= len(*Adj)-1 {
			//fmt.Println("List:", *list)

			tmp := make([]int, len(*list))
			copy(tmp, *list)
			*ans = append(*ans, tmp)
			break
		}
		_, alreadyVisited := (*visited)[path]
		if alreadyVisited {
			continue
		}

		dfs(path, Adj, visited, ans, list)
		*list = (*list)[:len(*list)-1]
	}
	// fmt.Println("vertex removed:", vertex)
	delete(*visited, vertex.(int))
}

func allPathsSourceTarget(graph [][]int) [][]int {

	Adj := make(map[interface{}][]interface{})
	visited := make(map[interface{}]struct{})
	ans := make([][]int, 0)

	// build Adj list
	for i, edges := range graph {
		u := i
		if len(edges) == 0 {
			Adj[u] = append(Adj[u], -1)
			continue
		}
		for _, vertex := range edges {
			Adj[u] = append(Adj[u], vertex)
		}
	}
	// fmt.Println("Adj:", Adj)
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
