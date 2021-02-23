// Description: Appending to slices
// References:
//	https://tour.golang.org/moretypes/15

package main

import "fmt"

func main() {
	// when appending, the first parameter is a []T (a slice of type T)
	// the rest of the parameters are the values to append
	// the resulting slice will be returned
	// if the backing array is too small, a bigger one will be allocated

	var s []int
	printSlice(s)

	// append on nil slices
	s = append(s, 0)
	printSlice(s)

	// the slice can grow as needed
	s = append(s, 1)
	printSlice(s)

	// we can also add more than one elem at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
