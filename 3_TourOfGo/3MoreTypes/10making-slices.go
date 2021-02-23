// Description: Using the make function to create slices
// References:
//	https://tour.golang.org/moretypes/13

package main

import "fmt"

func main() {
	// make allocs a zeroed array and returns a slice referring to that array

	a := make([]int, 5) // len(a) = 5
	printSlice("a", a)

	b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
	printSlice("b", b)

	// b = b[:cap(b)] 		// len(b) = 5, cap(b) = 5
	// b = b[1:] 			// len(b) = 4, cap(b) = 4

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
