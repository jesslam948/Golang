// Definition: More things we can do with functions
// Resources:
//	https://tour.golang.org/basics/6

package main

import "fmt"

// this function swaps two strings
// notice that we can return multiple outputs in this way
func swap(x, y string) (string, string) {
	return y, x
}

// function splits a sum into two values
// notice that we declare the return variables in the func def
// this is called a "naked" return, but it's not very readable for long functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("Hello", "World")
	fmt.Println("Swap:", a, b)
	c, d := split(17)
	fmt.Println("Split:", c, d)
}
