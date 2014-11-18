package main

import (
	"fmt"
	"time"
)

// note the string must be "2006-01-02 15:04:05 ..."
const FORMAT = "2006-01-02 15:04:05"

func main() {
	fmt.Printf("Current time is %s\n", time.Now())
	fmt.Printf("Current time is %s\n", time.Now().Format(FORMAT))
}
