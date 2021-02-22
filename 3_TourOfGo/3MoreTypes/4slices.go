// Description: Slices in Go
// References:
//	https://tour.golang.org/moretypes/7

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// type []T is a slice with elements of type T
	// slices are dynamically-sized and give a flexible view of the elements of an array
	// slices are more common thann arrays in practice
	// the low bound is inclusive while the high bound in the slice is exclusive
}
