// https://tour.golang.org/flowcontrol/8 (2017-04-01)
// Exercise: Loops and Functions

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for n := 0; n < 10; n++ {
		z = (z - ((z*z)-x)/(2*z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
