package main

import "fmt"

func main() {

	hNum := "2AX" // 42

	total := 0
	size := len(hNum)
	if size < 1 {
		fmt.Println("total:", total)
		return
	}

	for i := 0; i < size; i++ {
		ch := hNum[i]
		if ch >= '0' && ch <= '9' {
			total = 16*total + int(ch-'0')
		} else if ch >= 'a' && ch <= 'f' {
			total = total*16 + int(ch-'a') + 10
		} else if ch >= 'A' && ch <= 'F' {
			total = 16*total + int(ch-'A') + 10
		} else {
			fmt.Println("total:", total)
			return
		}
	}
	fmt.Println("total:", total)
	return
}
