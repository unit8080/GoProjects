package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.GOOS)
	fmt.Println(strconv.IntSize)
	fmt.Println(unsafe.Sizeof(int(0)))
	var i int
	fmt.Printf("Size of var (reflect.TypeOf.Size): %d\n", reflect.TypeOf(i).Size())
	fmt.Printf("Size of var (unsafe.Sizeof): %d\n", unsafe.Sizeof(i))
}
