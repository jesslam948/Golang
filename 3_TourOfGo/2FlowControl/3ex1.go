// Description: Exercise on Loops and Functions
// References:
//	https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	smallNum := 0.000000000000001
	z := float64(1)
	for {
		fix := (z*z - x) / (2 * z)
		z -= fix
		if math.Abs(fix) <= smallNum {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
