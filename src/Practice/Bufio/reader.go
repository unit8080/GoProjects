package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createFile() {
	file, err := os.Create("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.Write([]byte("Hello World!\n"))
	file.Write([]byte("Hello Diya#\n"))
	file.Write([]byte("Hello Keshav$$\n"))
}

func main() {

	createFile()
	file, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// fmt.Println(reader) - this will print reader struct

	data := make([]byte, 20)
	count := 0
	for {
		count++
		n, err := reader.Read(data)
		if err == io.EOF {
			fmt.Println("count:", count)
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(n, string(data))
	}
}
