package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	}
}
