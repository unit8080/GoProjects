package main

import "fmt"

type Graph map[interface{}][][]interface{} // Adjacency list
type Visited map[interface{}]struct{}      // hashSet
type Queue []interface{}                   // BFS queue

func Constructor() Graph {
	g := make(map[interface{}][][]interface{})
	return g
}
func (g Graph) Add(data1, data2, data3 interface{}) {
	v, exists := g[data1]
	if exists {
		v = append(v, []interface{}{data2, data3})
	} else {
		v = [][]interface{}{{data2, data3}}
	}
	g[data1] = v
}

func (v Visited) IsVisited(vertex interface{}) bool {
	_, exists := v[vertex]
	if exists {
		return true
	} else {
		return false
	}
}
func (g Graph) Print(visit Visited) {
	fmt.Println("size:", len(g))

	for k, v := range g {
		if visit.IsVisited(k) {
			fmt.Println("")
		}
		for _, val := range v {
			fmt.Println("U:", k, "-->", "V:", val[0], val[1])
		}
	}
}
func main() {

	g := Constructor()
	v := make(map[interface{}]struct{})
	g.Add(0, 1, 5)
	g.Add(0, 2, 10)
	g.Add("A", "B", 20)
	g.Print(v)
	fmt.Println("Graph:", g)

}
