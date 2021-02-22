// Definition: Play around with switch statements
// References:
//	https://tour.golang.org/flowcontrol/9
//	https://tour.golang.org/flowcontrol/10
//	https://tour.golang.org/flowcontrol/11

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// part 1
	fmt.Print("Go runs on ")
	// runs the first case whose value matches the condition
	// unline other languages, Go only runs the selected cases
	// also, switch cases don't have to be constants and values don't need to be integers
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("----------")

	// part 2
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday() // today gets a number 0-6 representing Sunday, Monday, etc.
	// cases are evaluated top to bottom
	switch time.Saturday { // time.Saturday = 6
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Println("----------")

	// part 3
	t := time.Now()
	// can have a switch without a condition, it's the same as `switch true`
	// and helps write long if-then-else chains
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
