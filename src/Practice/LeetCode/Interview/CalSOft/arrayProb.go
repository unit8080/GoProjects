// [4:38 PM] Mohsin Khazi (Unverified)
// in_array = [2, 20, 25, 7, 11, 33, 3]

// [4:38 PM] Mohsin Khazi (Unverified)
// out_array = [[2,20,25], [7,11,33], [3]]

// in_arry = [20, 15, 10, 5, 1]
// out_arry = [[20], [15], [10], [5], [1]]
package main

import "fmt"

func main() {
	// in_array := []int{2, 20, 25, 7, 11, 33, 3}
	// in_array := []int{20, 15, 10, 5, 1}
	in_array := []int{}
	fmt.Println(getSortedArray(in_array))
}

func getSortedArray(arr []int) [][]int {

	ans := [][]int{}
	temp := []int{}

	if len(arr) < 1 {
		return ans
	} else if len(arr) == 1 {
		temp = append(temp, arr[0])
		ans = append(ans, temp)
		return ans
	}
	lastNum := arr[0]
	temp = append(temp, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] < lastNum {
			ans = append(ans, temp)
			temp = []int{}
		}

		temp = append(temp, arr[i])
		lastNum = arr[i]
	}
	if len(temp) > 0 {
		ans = append(ans, temp)
	}
	return ans
}
