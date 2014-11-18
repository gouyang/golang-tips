package main

import (
	"fmt"
)

func for1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func for2() int {
	sum := 1
	// leave the pre and post statements empty.
	for sum < 100 {
		sum += sum
	}
	return sum
}

func for3() int {
	sum := 1
	// while like in C.
	for sum < 200 {
		sum += sum
	}
	return sum
}

func main() {
	fmt.Println(for1())
	fmt.Println(for2())
	fmt.Println(for3())
}
