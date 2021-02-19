// Description: if statements in go
// References:
//	https://tour.golang.org/flowcontrol/5
//	https://tour.golang.org/flowcontrol/6
//	https://tour.golang.org/flowcontrol/7

package main

import (
	"fmt"
	"math"
)

// part 1
func sqrt(x float64) string {
	// like for loops, parentheses not required but brackets are
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// part 2
func pow(x, n, lim float64) float64 {
	// can have a short statement to execute before if
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// v is only in the scope of the if statement
	return lim
}

// part 3
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// this is an else
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// part 1
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println("----------")

	// part 2
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("----------")

	// part 3
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
