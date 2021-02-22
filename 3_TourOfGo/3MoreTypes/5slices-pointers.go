// Description: Working with slice pointers
// References:
//	https://tour.golang.org/moretypes/8

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:4]
	fmt.Println(a, b)

	// slices don't store any data
	// changing the elements of a slice changes the underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
