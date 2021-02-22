// Description: Length and capacity of a slice
// References:
//	https://tour.golang.org/moretypes/11

package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	printSlice(primes)

	// Slice to give 0 length
	primes = primes[:0]
	printSlice(primes)

	// Extend length
	primes = primes[:4]
	printSlice(primes)

	// Drop first two values
	primes = primes[2:]
	printSlice(primes)

	// // Try to add back first two values (CAN'T)
	// primes = primes[-2:]
	// printSlice(primes)

	// // Try to extend past cap (CAN'T)
	// primes = primes[:8]
	// printSlice(primes)
}

func printSlice(s []int) {
	// len is the number of elements in the slice
	// cap is the number of elems in the array starting from the beginning of the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
