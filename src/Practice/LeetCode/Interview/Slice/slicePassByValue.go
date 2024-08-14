package main

import "fmt"

func main() {

	arr := make([]int, 0, 2)

	doSomething(arr)
	fmt.Println(arr)
	fmt.Println(arr[:1])
	doSomethingByReference(&arr)
	fmt.Println(arr)
}
func doSomething(arr []int) {
	arr = append(arr, 2) // this change is not visible to main
}
func doSomethingByReference(arr *[]int) {
	*arr = append(*arr, 3)
}
