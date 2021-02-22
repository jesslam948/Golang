// Description: Playing with defer statements
// References:
//	https://tour.golang.org/flowcontrol/12
//	https://tour.golang.org/flowcontrol/13

package main

import "fmt"

func main() {
	// part 1
	// defer call args are evaluated immediately but isn't executed until the surrounding function returns
	defer fmt.Println("world")
	fmt.Println("hello")
	defer fmt.Println("----------")

	fmt.Println("----------")

	// part 2
	fmt.Println("counting")
	// deferred function calls are pushed onto a stack so they are executed in LIFO order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
