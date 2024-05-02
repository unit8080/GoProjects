package main

import "fmt"

/*
type void struct{}
var member void

set := make(map[string]void) // New empty set
set["Foo"] = member          // Add
for k := range set {         // Loop
	fmt.Println(k)
}
delete(set, "Foo")      // Delete
size := len(set)        // Size
_, exists := set["Foo"] // Membership
*/

type Set map[interface{}]struct{}

func Constructor() Set {
	set := make(map[interface{}]struct{})
	return set
}
func (set Set) Add(data interface{}) {
	set[data] = struct{}{}
}
func (set Set) Delete(data interface{}) {
	delete(set, data)
}
func (set Set) Exists(data interface{}) bool {
	_, exists := set[data]
	return exists
}
func (set Set) Len() int {
	return len(set)
}
func (set Set) Print() {
	for v := range set {
		fmt.Println(v)
	}
}
func main() {
	set := Constructor()

	set.Add(5)
	set.Add(4)
	set.Delete(5)
	set.Add(8)
	set.Add("test")
	set.Print()
	fmt.Println(set.Exists(4))
	fmt.Println(set.Exists(10))
	fmt.Println(set.Exists("bad"))
	fmt.Println("len:", set.Len())
}
