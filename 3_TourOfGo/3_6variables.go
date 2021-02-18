// Description: Playing with Go variables
// Resources:
//	https://tour.golang.org/basics/8
//	https://tour.golang.org/basics/9
//	https://tour.golang.org/basics/10

package main

import "fmt"

// part 1
// var declares a list of variables and the type is last
var a, b, c bool

// part 2
// initialize variables in declaration
var j, k int = 1, 2

func main() {
	// part 1
	// the default int value is 0 and for boolean it's false
	var i int
	fmt.Println("part1:", i, a, b, c)

	// part 2
	// you can omit the type if you initialize right away
	var d, e, f = true, false, "no!"
	fmt.Println("part2:", j, k, d, e, f)

	// part 3
	x := 5
	fmt.Println("part3:", x)
}
