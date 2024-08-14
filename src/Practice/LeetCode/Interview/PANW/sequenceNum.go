package main

import "fmt"

func main() {

	num := 3210 // 1234 // 3210
	if num < 1 {
		return
	}
	prev := num % 10
	num = num / 10

	for num > 0 {
		curr := num % 10
		if curr != prev+1 // prev - 1 {
			fmt.Println("not sequence")
			return
		}
		prev = curr
		num = num / 10
	}
	fmt.Println("num has sequence")
}
