// Description: Loops in Go
// References:
//	https://tour.golang.org/flowcontrol/1
//	https://tour.golang.org/flowcontrol/2
//	https://tour.golang.org/flowcontrol/3

package main

import "fmt"

func main() {
	// part 1
	sum := 0
	// Go only has for loops
	// they have to fit this structure:
	// 	init statement
	// 	condition expr
	//	post statement
	// no parentheses and braces required
	for i := 0; i < 10; i++ {
		sum += i
	}
	// like other languages, i is only visible in scope of the for block
	fmt.Println(sum)

	fmt.Println("----------")

	// part 2
	sum2 := 1
	// the init and post statements are optional
	// could be written for ; sum < 1000 ' {
	// but at this point acts as a 'while' loop
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// part 3
	// this is an infinite loop
	// for {
	//
	// }
}
