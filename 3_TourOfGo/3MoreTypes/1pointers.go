// Description: Pointers in Go
// References:
//	https://tour.golang.org/moretypes/1

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i                  // points to i
	fmt.Println("i:", *p)    // read i through the pointer
	*p = 21                  // set i through the pointer
	fmt.Println("i_new:", i) // print the new value of i

	p = &j                   // points to j
	fmt.Println("j:", *p)    // read j through the pointer
	*p = *p / 37             // divide j through the pointer
	fmt.Println("j_new:", j) // print the new value of j

	// type *T is a pointer to a T value
	// 	zero value is `nil`
	// 	i.e. var p *int is a pointer to an int
	// the & operator generates a pointer to its operand
	// the * pointer 'dereferences' or 'indirects' the pointer and gives us the underlying value
}
