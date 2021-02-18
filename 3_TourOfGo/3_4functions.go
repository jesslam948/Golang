// Description: My first function
// Resources:
//	https://tour.golang.org/basics/4
//	https://tour.golang.org/basics/5

package main

import "fmt"

// notice that the type is declared after var name
func add(x int, y int) int {
	return x + y
}

// we can also shorten this func def as follows
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("add1:", add(42, 13))
	fmt.Println("add2:", add(42, 13))
}
