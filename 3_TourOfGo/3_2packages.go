// Description: Modified the ToG packages file so that we get random values
// 	using the time as the seed
// Resources:
//	https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
//	https://tour.golang.org/basics/1
//	https://tour.golang.org/basics/2

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("A random number is", rand.Intn(10))
}
