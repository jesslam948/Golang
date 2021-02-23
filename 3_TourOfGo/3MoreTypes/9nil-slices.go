// Description: Nil slices
// References:
//	https://tour.golang.org/moretypes/12

package main

import "fmt"

func main() {
	// a slice's default/zero value is nil and has 0 length and capacity
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
