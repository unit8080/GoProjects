package main

import "fmt"

func Pop(nums *[]int) int {
	x := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	return x
}
func Push(nums *[]int, x int) {
	(*nums) = append((*nums), x)
}
func reverse(nums *[]int, st *[]int) {
	// base case
	if len(*nums) <= 0 {
		return
	}

	top := Pop(nums)
	Push(st, top)
	reverse(nums, st)
}
func main() {
	stack := []int{1, 2, 3, 4, 5}
	fmt.Println(stack)
	newSt := []int{}
	reverse(&stack, &newSt)
	fmt.Println(newSt)
}
