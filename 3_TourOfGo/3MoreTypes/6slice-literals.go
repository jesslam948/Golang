// Description: Play with slice literals
// References:
//	https://tour.golang.org/moretypes/9

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)

	isPrime := []bool{false, true, true, false, true, false}
	fmt.Println(isPrime)

	combined := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(combined)

	// slice literals are like array literals without the length
	// it creates the array and then builds a slice that references it
}
