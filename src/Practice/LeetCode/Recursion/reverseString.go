package main

import "fmt"

func main() {

	str := "leetcode"
	reverse := ""
	printReverse([]byte(str), 0, &reverse)
	fmt.Println(reverse)
}

func printReverse(str []byte, idx int, s *string) {
	// base case
	if idx >= len(str) {
		return
	}
	printReverse(str, idx+1, s)
	(*s) += string(str[idx])
}
