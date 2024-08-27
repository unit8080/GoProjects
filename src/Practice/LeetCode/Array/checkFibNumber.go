// To execute Go code, please declare a func main() in a package "main"
/*
Write a function that when given an array of integers, returns an array of those integers that are in the Fibonacci sequence

For x=[]
x[n] = x[n-1] + x[n-2]

0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811

[35, 233, 44, 6, 23] —> [233]

[89, 32, 55, 142, 13, 1, 365] —> [89, 55, 13, 1]

No decimal numbers in input
No negative numbers in input
No duplicated numbers in input
Input will not be sorted and output is not required to be sorted
*/

package main

import "fmt"

func nextFib(num1, num2 int) int {
	return num1 + num2
}
func getNextFib(num int) int {

	// check if it is fib numbers
	if num == 0 || num == 1 {
		return num
	}
	num1 := 0
	num2 := 1

	nextNum := nextFib(num1, num2)
	for nextNum < num {
		num1 = num2
		num2 = nextNum
		nextNum = nextFib(num1, num2)
	}
	if nextNum == num {
		return num
	} else {
		return -1
	}
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, Sanjeev!")
	}
	// list := []int {35, 233, 44, 6, 23}
	list := []int{89, 32, 55, 142, 13, 1, 365}
	// list := []int{}
	ans := []int{}
	for i := 0; i < len(list); i++ {
		num := getNextFib(list[i])
		if num == list[i] {
			ans = append(ans, num)
		}
	}
	fmt.Println("List", ans)
}
