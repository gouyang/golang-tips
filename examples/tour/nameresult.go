package main

import (
	"fmt"
)

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	fmt.Printf("3 + 4 = %d\n", add(3, 4))

}
