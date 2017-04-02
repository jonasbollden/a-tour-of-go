// https://tour.golang.org/flowcontrol/8 (2017-04-01)
// Exercise: Loops and Functions

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	t := float64(0) // tmp
	for {
		z = (z - ((z*z)-x)/(2*z))
		if math.Abs(t-z) < 1e-15 {
			break
		}
		t = z
	}
	return t
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
