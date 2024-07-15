package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// writer := bufio.NewWriter(file)
	writer := bufio.NewWriterSize(file, 100)
	_, err = writer.Write([]byte("Hello Humans!\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}
