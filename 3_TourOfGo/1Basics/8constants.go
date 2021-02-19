// Description: Constants in Go
// References:
//	https://tour.golang.org/basics/15
//	https://tour.golang.org/basics/16

package main

import "fmt"

// part 1
const Pi = 3.14

// part 2
const (
	// Create a huge number by shifting a 1 b it left 100 places
	Big = 1 << 100
	// Shift it right again 99 places so it becomes 2
	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	// part 1
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println("----------")

	// part 2
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
