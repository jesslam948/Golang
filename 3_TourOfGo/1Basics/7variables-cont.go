// Description: Basic types
// References:
//	https://tour.golang.org/basics/11
//	https://tour.golang.org/basics/12
//	https://tour.golang.org/basics/13
//	https://tour.golang.org/basics/14

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// part 1
// notice that variable declarations can be "factored" into blocks
// refer back to first link for all of Go's basic types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// part 2
var (
	i int
	f float64
	b bool
	s string
)

func main() {
	// part 1
	// examining types and values
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("----------")

	// part 2
	// examining default values of certain types
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("----------")

	// part 3
	// type conversion (note this has to be explicit)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	fmt.Println("----------")

	// part 4
	// rmb that you can delcare vars without specifying an explicit type
	v := 42
	var j = "hello"
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("j is of type %T\n", j)
}
