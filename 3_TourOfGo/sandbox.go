// Description: Simple program that prints out the time
// Resources:
//	https://tour.golang.org/welcome/4

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the sandbox! This is from the Tour of Go tutorial!")

	fmt.Println("The time is", time.Now())
}
