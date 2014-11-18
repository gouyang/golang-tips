package main

import (
	"fmt"
	"math"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
