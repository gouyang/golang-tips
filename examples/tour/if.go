package main

import (
	"fmt"
	"math"
)

func if1() {
	if 1 < 2 {
		fmt.Println("1 < 2 is true")
	}
}

func if2() {
	if v := math.Pow(2, 3); v < 10 {
		fmt.Printf("number %f less than 10\n", v)
	} else {
		fmt.Printf("number %f large than 10\n", v)
	}
}

func main() {
	if1()
	if2()
}
