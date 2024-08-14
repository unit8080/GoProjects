package main

import "fmt"

type gopher struct {
	name string
}

func (r *gopher) print() {
	fmt.Printf("receiver: %v \n", r)
	fmt.Println("gopher-printer works!")
	// Following nil receiver dereference will cause:
	// panic: runtime error: invalid memory address
	// or nil pointer dereference
	fmt.Println("gopher-printer works!", r.name)

	// Receivers in Go act like a basic argument, and
	// in this case, when we call a method on an
	// uninitialised struct, inside the methods nil has
	// passed as a receiver, therefore no effect at compile
	// and runtime unless access the receiver struct
	// fileds.

	// Summary: The nil receiver is just a parameter.

}

func main() {
	var gpr *gopher
	gpr.print()
}
