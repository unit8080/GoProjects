package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6} // { 1, 4, 3, 2, 5, 6}

	start := 2
	end := 4
	reverse(arr, start, end)
	fmt.Println("arr:", arr)
}

func reverse(arr []int, l int, r int) {
	size := len(arr)
	if size <= 1 {
		return
	}
	if l >= r || l >= size || r >= size || l < 0 || r < 0 {
		return
	}

	i := l - 1
	j := r - 1
	for i < j {

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println("arr:", arr)
}
