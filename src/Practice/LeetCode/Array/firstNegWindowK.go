package main

import "fmt"

func main() {

	input := []int{8, 1, -2, 2, -3, 6, 8, -1} // {-2, -2, -2, -3, -3, -1}, k = 3
	// input := []int{-8, 2, 3, -6, 10} // {-8, 0, -6, -6}, k = 2

	k := 3 // window size
	i := 0
	j := 0

	n := len(input)

	ans := []int{}
	queue := []int{}
	for j < n {
		if input[j] < 0 {
			queue = append(queue, input[j])
		}
		if j-i+1 == k { // meet the window size
			if len(queue) > 0 {
				ans = append(ans, queue[0])
			} else {
				ans = append(ans, 0)
			}
			if input[i] < 0 {
				queue = queue[1:] // dequeue front
			}
			i++
		}
		j++
	}
	fmt.Println("ans:", ans)
}
