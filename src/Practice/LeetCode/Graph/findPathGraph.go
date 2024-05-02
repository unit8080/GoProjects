// 1971. Find if Path Exists in Graph
// https://leetcode.com/problems/find-if-path-exists-in-graph/

type Graph map[interface{}][][]interface{}
type Visited map[interface{}]struct{}
type Queue [][]interface{}

func Constructor() Graph {
    g := make(map[interface{}][][]interface{})
    return g
}
func (g Graph) Add(u, v interface{}) {
    val, exists := g[u]
    if exists {
        val = append(val, []interface{}{v, 1})
    }else {
        val = [][]interface{}{{v, 1}}
    }
    g[u] = val
    // fmt.Println("dest:", val)
}
func (v Visited) IsVisited(vertex interface{}) bool {
    _, exists := v[vertex]
    if exists {
        return true
    } else {
        return false
    }
}
func BuildVisit() Visited {
    v := make(map[interface{}]struct{})
    return v
}
func validPath(n int, edges [][]int, source int, destination int) bool {
    g := Constructor()
    q := make([][]interface{}, 0)
    // visit := make(map[interface{}]struct{})
    visit := BuildVisit()
    if len(edges) == 0 { return true }
    for _, edge := range edges {
        // fmt.Println("edge:", edge)
        g.Add(edge[0], edge[1])
        g.Add(edge[1], edge[0])
    }

    // fmt.Println("Map:", g)
    q = append(q, []interface{}{source})
    for len(q) > 0 {
        list := q[0]
        q = q[1:]
        vertex := list[len(list)-1]

        if visit.IsVisited(vertex) {
            continue
        }
        /*
        _, visited := visit[vertex]
        if visited {
            continue
        }
        */
        visit[vertex] = struct{}{}

        vertices := g[vertex]
        for _, dest := range vertices {
            // build path here
            l := make([]interface{}, len(list))
            copy(l, list)
            l = append(l, dest[0])

            if dest[0] == destination {
                // fmt.Println(l)
                return true
            }
            // q = append(q, dest[0])
            q = append(q, l)
        }
    }
    return false
}
