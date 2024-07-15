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
		data, err := reader.ReadBytes('\n')
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

// The difference between bufio.Reader.ReadSlice and bufio.Reader.ReadBytes
// is that bufio.Reader.ReadSlice will return an error if the delimiter
// is not found.
