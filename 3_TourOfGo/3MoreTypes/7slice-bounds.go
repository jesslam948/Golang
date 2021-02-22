// Description: Explore slice bound defaults
// References:
//	https://tour.golang.org/moretypes/10

package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	primes = primes[1:4]
	fmt.Println(primes)

	primes = primes[:2]
	fmt.Println(primes)

	primes = primes[1:]
	fmt.Println(primes)

	// slices in Go act similar to slicing in python
	// default low bound is 0 and high bound is the length of the slice/array
}
