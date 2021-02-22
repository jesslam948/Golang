// Definition: Play with structs in Go
// References:
//	https://tour.golang.org/moretypes/2
//	https://tour.golang.org/moretypes/3
//	https://tour.golang.org/moretypes/4
//	https://tour.golang.org/moretypes/5

package main

import "fmt"

// part 1
// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

/*
// the struct can also be declared like so
type Vertex struct {
	X, Y int
}
*/

// part 4
var (
	// these are struct literals, which denote a subset of the fields (ordering is irrelevant)
	v1 = Vertex{Y: 1}  // X:0 is implicit
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2 = &Vertex{1, 2} // has type *Vertex
)

func main() {
	// part 1
	// notice how we instantiate a struct
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Println("----------")

	// part 2
	// we can use the . operator to access struct fields
	v.X = 4
	fmt.Println(v.X)

	fmt.Println("----------")

	// part 3
	// we can use struct pointers to point to a struct
	p := &v
	// we can access and set struct fields though a struct pointer
	// note we would normally dereferences first (*p).x
	// 	but Go allows us to access the struct fields w/o an explicit dereference
	p.X = 1e9
	fmt.Println(v)

	fmt.Println("----------")

	// part 4
	fmt.Println(v1, v2, v3, p2)
	// note Go prefers local to global scope like most other languages
}
